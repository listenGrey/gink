package transfer

import (
	"fmt"
	"gink/pkg/logger"
	"gink/pkg/utils"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
	"time"
)

type WebSocketTransfer struct{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有CORS请求
	},
}

// Send 发送文件到指定的地址和端口
func (wst *WebSocketTransfer) Send(filepath string, destinationIndex string) error {
	// 获取当前时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 获取文件信息，打开文件
	fileInfo, file, err := GetFile(filepath)
	if err != nil {
		NewRecord(filepath, "", now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error getting file", zap.Error(err))
		return err
	}
	defer file.Close()

	// 计算文件哈希值
	filehash, err := utils.CalculateFileHash(file)
	if err != nil {
		NewRecord(filepath, "", now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error calculating file hash", zap.Error(err))
		return err
	}

	// 获取目的地IP
	destination, err := GetDestination(destinationIndex)
	if err != nil {
		NewRecord(filepath, "", now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error transform index", zap.Error(err))
		return err
	}

	// 建立连接
	c, _, err := websocket.DefaultDialer.Dial("ws://"+destination+"/receive", nil)
	if err != nil {
		NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error connecting to host", zap.Error(err))
		return err
	}
	defer c.Close()

	// 发送文件名和哈希值
	payload := fileInfo.Name() + "/" + filehash
	err = c.WriteMessage(websocket.TextMessage, []byte(payload))
	if err != nil {
		NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error sending file info", zap.Error(err))
		return err
	}

	// 重置文件指针
	file.Seek(0, 0)

	// 读取文件并发送
	buf := make([]byte, 1024) // 创建一个缓冲区来存储文件数据
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
			logger.Log.Error("Error sending file", zap.Error(err))
		}
		err = c.WriteMessage(websocket.BinaryMessage, buf[:n])
		if err != nil {
			NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
			logger.Log.Error("Error sending file", zap.Error(err))
			break
		}
	}

	// 文件传输完毕，发送关闭消息
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error sending confirmation", zap.Error(err))
		return err
	}

	NewRecord(filepath, destination, now, "", true, false)
	fmt.Println("File sent successfully")
	return nil
}

// Receive 启动监听服务，接收文件
func (wst *WebSocketTransfer) Receive() error {
	// 监听端口
	http.HandleFunc("/receive", handleConnection)
	fmt.Println("WebSocket listening started on :8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		return err
	}

	return nil
}

// Stop 关闭监听
func (wst *WebSocketTransfer) Stop() error {
	return nil
}

// handleConnection 处理接收的连接
func handleConnection(w http.ResponseWriter, r *http.Request) {
	// 获取当前时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 建立连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		NewRecord("", conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error creating connection", zap.Error(err))
		return
	}
	defer conn.Close()

	// 读取第一个消息作为文件名和哈希值
	_, fileinfo, err := conn.ReadMessage()
	if err != nil {
		NewRecord("", conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error receiving file name", zap.Error(err))
		return
	}
	fileinfos := strings.Split(string(fileinfo), "/")
	filename, filehash := fileinfos[0], fileinfos[1]

	// 保存文件的位置
	filePath := utils.NewFilePath(filename) // 生成新路径
	file, err := os.Create(filePath)
	if err != nil {
		NewRecord(filename, conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error creating file", zap.String("file", filePath), zap.Error(err))
		return
	}
	defer file.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				break
			}
			NewRecord(filename, conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
			logger.Log.Error("Error reading to file", zap.String("file", filePath), zap.Error(err))
			break
		}
		if mt == websocket.BinaryMessage {
			if _, err := file.Write(message); err != nil {
				NewRecord(filename, conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
				logger.Log.Error("Error writing to file", zap.String("file", filePath), zap.Error(err))
				break
			}
		}
	}

	// 重置文件指针
	file.Seek(0, 0)

	// 计算文件哈希值
	filehashed, err := utils.CalculateFileHash(file)
	if err != nil {
		NewRecord(filename, conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error calculating file hash", zap.Error(err))
		return
	}

	// 比较哈希值，如果不一致，则建议删除文件
	if filehash != filehashed {
		fmt.Printf("%s has risk, should delete it", filePath)
		return
	}

	NewRecord(filename, conn.RemoteAddr().String(), now, "", true, true)
	fmt.Printf("File received and saved: %s\n", filePath)
}
