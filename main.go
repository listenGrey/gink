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

/*
 - 增加传输日志 yes
 - 增加进度条 yes
 - 加密 yes
 - 传文件夹 NO
 - 同名文件 yes
 - WebSocket 本机传yes、远程NO
 - 路径覆盖而不是追加 yes
 - 关闭应用释放资源 稍后
 - README
 - 环境变量（直接输入gink）、windows、Linux的下载release、下载后创建历史和配置文件，默认路径
 - UDP、gRpc
*/
