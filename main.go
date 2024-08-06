package main

import (
	"fmt"
	"gink/cmd"
	"gink/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("配置文件读取错误")
		return
	}
	cmd.Execute()
}

/*
 - 增加日志
 - 增加进度条
 - 加密
*/
