package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gink/config"
	"io"
	"os"
	"path/filepath"
)

// NewFilePath 生成文件的新路径
func NewFilePath(filename string) string {
	// 生成文件的完整路径
	path := config.AppConfig.LocalDirection + "/" + filename

	// 如果文件不存在，返回原始路径
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path
	}

	// 如果文件已存在，创建新文件名，并不断检测
	ext := filepath.Ext(filename)                // 原文件后缀名
	name := filename[0 : len(filename)-len(ext)] // 原文件名
	newPath := ""
	for i := 1; i < 99999; i++ {
		newFilename := fmt.Sprintf("%s(%d)%s", name, i, ext)
		newPath = config.AppConfig.LocalDirection + "/" + newFilename
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			break
		}
	}

	return newPath
}

// CalculateFileHash 读取文件并计算其 SHA-256 哈希值
func CalculateFileHash(file *os.File) (string, error) {
	harsher := sha256.New()
	if _, err := io.Copy(harsher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(harsher.Sum(nil)), nil
}
