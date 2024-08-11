package transfer

import (
	"fmt"
	"gink/config"
	"os"
	"strconv"
)

// GetFile 获取文件信息
func GetFile(filepath string) (fileInfo os.FileInfo, file *os.File, err error) {
	// 文件信息
	fileInfo, err = os.Stat(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting file info: %v", err)
	}

	// 判断是文件还是文件夹，如果是文件夹需要压缩
	if fileInfo.IsDir() {
		return nil, nil, fmt.Errorf("error file type: %v", err)
	}

	// 使用文件路径打开文件
	file, err = os.Open(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}

	return
}

// GetDestination 获取目的地IP
func GetDestination(destinationIndex string) (string, error) {
	index, err := strconv.Atoi(destinationIndex)
	if err != nil {
		return "", fmt.Errorf("error transform index")
	}
	return config.AppConfig.Destinations[index-1], nil
}
