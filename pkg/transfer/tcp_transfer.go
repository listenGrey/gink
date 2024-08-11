package transfer

import (
	"encoding/binary"
	"fmt"
	"gink/pkg/logger"
	"gink/pkg/utils"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"go.uber.org/zap"
	"io"
	"net"
	"os"
	"time"
)

type TCPTransfer struct{}

// Send 发送文件到指定的地址和端口
func (t *TCPTransfer) Send(filepath string, destinationIndex string) error {
	// 获取当前时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 获取文件信息
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

	// 设置网络连接
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error connecting to host", zap.Error(err))
		return err
	}
	defer conn.Close()

	// 发送文件名
	fileNameLength := uint32(len(fileInfo.Name()))
	binary.Write(conn, binary.LittleEndian, fileNameLength)
	conn.Write([]byte(fileInfo.Name()))

	// 发送文件的哈希值
	filehashLength := uint32(len(filehash))
	binary.Write(conn, binary.LittleEndian, filehashLength)
	conn.Write([]byte(filehash))

	// 重置文件指针
	file.Seek(0, 0)

	// 创建进度条
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(fileInfo.Size(),
		mpb.PrependDecorators(
			decor.Name("Sending: "),
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_MMSS, 60),
			decor.Name("  "),
			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60),
		),
	)

	// 创建一个多重写入器，一个写入网络连接，另一个更新进度条
	proxyReader := bar.ProxyReader(file)
	_, err = io.Copy(conn, proxyReader)
	if err != nil {
		NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
		logger.Log.Error("Error sending file", zap.Error(err))
		return err
	}

	p.Wait() // 等待所有进度条完成

	/*
		// 直接发送
		_, err = io.Copy(conn, file)
		if err != nil {
			t.NewRecord(filepath, destination, now, fmt.Sprintf("%s", err), false, false)
			logger.Log.Error("Error sending file", zap.Error(err))
			return fmt.Errorf("error sending file: %v", err)
		}
	*/

	NewRecord(filepath, destination, now, "", true, false)
	fmt.Println("File sent successfully")
	return nil
}

// Receive 启动监听服务，接收文件
func (t *TCPTransfer) Receive() error {
	// 确保端口与配置匹配
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		logger.Log.Error("Error listening on port 8000", zap.Error(err))
		return fmt.Errorf("Error listening: %v\n", err)
	}
	defer listener.Close()

	fmt.Println("TCP listening started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Log.Error("Error accepting connection", zap.Error(err))
			continue
		}
		go t.handleConnection(conn)
	}
}

// handleConnection 处理接收的连接
func (t *TCPTransfer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 获取当前时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 获取文件名
	var fileNameLength uint32
	binary.Read(conn, binary.LittleEndian, &fileNameLength)
	filename := make([]byte, fileNameLength)
	conn.Read(filename)

	// 获取文件的哈希值
	var filehashLength uint32
	binary.Read(conn, binary.LittleEndian, &filehashLength)
	filehash := make([]byte, filehashLength)
	conn.Read(filehash)

	// 保存文件的位置
	filePath := utils.NewFilePath(string(filename)) // 生成新路径
	file, err := os.Create(filePath)
	if err != nil {
		NewRecord(string(filename), conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error creating file", zap.String("file", filePath), zap.Error(err))
		return
	}
	defer file.Close()

	// 将接收的文件内容写入本地文件
	_, err = io.Copy(file, conn)
	if err != nil {
		NewRecord(string(filename), conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error writing to file", zap.String("file", filePath), zap.Error(err))
		return
	}

	// 重置文件指针
	file.Seek(0, 0)

	// 计算文件哈希值
	filehashed, err := utils.CalculateFileHash(file)
	if err != nil {
		NewRecord(string(filename), conn.RemoteAddr().String(), now, fmt.Sprintf("%s", err), false, true)
		logger.Log.Error("Error calculating file hash", zap.Error(err))
		return
	}

	// 比较哈希值，如果不一致，则删除文件
	if string(filehash) != filehashed {
		fmt.Printf("%s has risk, should delete it", filePath)
		return
	}

	NewRecord(string(filename), conn.RemoteAddr().String(), now, "", true, true)
	fmt.Printf("File received and saved: %s\n", filePath)
}
