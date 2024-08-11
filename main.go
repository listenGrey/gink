package main

import (
	"fmt"
	"gink/cmd"
	"gink/config"
	"gink/pkg/transfer"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("配置文件读取错误")
		return
	}
	err = transfer.LoadHistory(config.AppConfig.HistoryFilePath)
	if err != nil {
		fmt.Println("历史文件读取错误")
		return
	}
	cmd.Execute()
}

/*
 - 增加传输日志 yes
 - 增加进度条 yes
 - 加密 yes
 - 传文件夹 NO
 - 同名文件 yes
 - WebSocket、UDP、gRpc
 - README、
 - 环境变量（直接输入gink）、window、Linux的下载release
*/
