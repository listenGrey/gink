package main

import (
	"fmt"
	"gink/cmd"
	"gink/config"
	"gink/pkg/transfer"
)

func main() {
	// 加载配置文件
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("配置文件读取错误")
		return
	}

	// 加载历史文件
	err = transfer.LoadHistory(config.AppConfig.HistoryFilePath)
	if err != nil {
		fmt.Println("历史文件读取错误")
		return
	}

	// 主任务
	cmd.Execute()
}
