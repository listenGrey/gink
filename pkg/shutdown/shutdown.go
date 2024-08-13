package shutdown

import (
	"fmt"
	"gink/config"
	"gink/pkg/transfer"
	"os"
)

var StopChan = make(chan struct{})

func Close() error {
	// 清理资源
	var Trans transfer.Transfer
	switch config.AppConfig.Protocols[0] {
	case "websocket":
		Trans = &transfer.WebSocketTransfer{} // 使用websocket协议
	case "tcp":
		Trans = &transfer.TCPTransfer{} // 使用TCP协议
	default:
		fmt.Println("Protocol error")
	}
	err := Trans.Stop()
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
