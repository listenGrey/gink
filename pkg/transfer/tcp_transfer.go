package transfer

import (
	"encoding/binary"
	"fmt"
	"gink/config"
	"io"
	"net"
	"os"
	"strconv"
)

type TCPTransfer struct{}

// StartListener 启动监听服务
func (t *TCPTransfer) StartListener() error {
	listener, err := net.Listen("tcp", ":8000") // 确保端口与配置匹配
	if err != nil {
		return fmt.Errorf("Error listening: %v\n", err)
	}
	defer listener.Close()
	fmt.Println("Listening for incoming files...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting: %v\n", err)
			continue
		}
		go t.handleConnection(conn)
	}
}

// handleConnection 处理接收的连接
func (t *TCPTransfer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 获取文件名
	var fileNameLength uint32
	binary.Read(conn, binary.LittleEndian, &fileNameLength)
	fileName := make([]byte, fileNameLength)
	conn.Read(fileName)

	// 保存文件的位置
	filePath := config.AppConfig.LocalDirection + "/" + string(fileName) // 确保路径存在
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// 将接收的文件内容写入本地文件
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Printf("File received and saved to %s\n", filePath)
}

// SendFile 发送文件到指定的地址和端口
func (t *TCPTransfer) SendFile(filename string, destinationIndex string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	index, err := strconv.Atoi(destinationIndex)
	if err != nil {
		return fmt.Errorf("error transform index")
	}

	destination := config.AppConfig.Destinations[index-1]
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return fmt.Errorf("error connecting to host: %v", err)
	}
	defer conn.Close()

	// 发送文件名
	fileInfo, _ := file.Stat()
	fileName := fileInfo.Name()
	fileNameLength := uint32(len(fileName))
	binary.Write(conn, binary.LittleEndian, fileNameLength)
	conn.Write([]byte(fileName))

	// 发送文件
	_, err = io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("error sending file: %v", err)
	}

	fmt.Println("File sent successfully")
	return nil
}
