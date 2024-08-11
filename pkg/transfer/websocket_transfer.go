package transfer

import (
	"encoding/json"
	"fmt"
	"gink/pkg/logger"
	"gink/pkg/utils"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"time"
)

type WebSocketTransfer struct{}

type FileData struct {
	FileName string `json:"filename"`
	FileHash string `json:"filehash"`
}

func (t *WebSocketTransfer) Send(filepath string, destinationIndex string) error {
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

	// 连接到 WebSocket 服务器
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(destination, nil)
	if err != nil {
		return fmt.Errorf("dial error: %v", err)
	}
	defer conn.Close()

	// 发送文件名和元数据
	meta := FileData{
		FileName: fileInfo.Name(),
		FileHash: filehash,
	}
	metaData, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %v", err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, metaData); err != nil {
		return fmt.Errorf("failed to send metadata: %v", err)
	}

	// 发送文件数据
	buffer := make([]byte, 1024) // 调整缓冲区大小以适应需要
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %v", err)
		}
		if n == 0 {
			break
		}

		err = conn.WriteMessage(websocket.BinaryMessage, buffer[:n])
		if err != nil {
			return fmt.Errorf("write message error: %v", err)
		}
	}

	return nil
}

func (t *WebSocketTransfer) Receive() error {
	// 读取文件数据
	/*_, message, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read metadata error: %v", err)
	}

	var meta FileMetadata
	if err := json.Unmarshal(message, &meta); err != nil {
		return fmt.Errorf("unmarshal metadata error: %v", err)
	}

	fmt.Println("Receiving file:", meta.Filename)

	// 接收文件数据
	file, err := os.Create(meta.Filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// 继续接收数据直到文件传输完成
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				break
			}
			return fmt.Errorf("read file data error: %v", err)
		}
		if _, err := file.Write(data); err != nil {
			return fmt.Errorf("failed to write file: %v", err)
		}
	}
	*/
	return nil
}
