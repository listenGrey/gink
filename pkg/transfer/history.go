package transfer

import (
	"encoding/json"
	"fmt"
	"gink/config"
	"gink/pkg/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"sync"
)

type Record struct {
	FileName     string `json:"file_name"`
	Destination  string `json:"destination"`
	Time         string `json:"time"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	Receive      bool   `json:"receive"` // 是接收到的文件就是true
}

var (
	transferHistory []Record
	mutex           sync.Mutex
)

// NewRecord 插入新的历史记录
func NewRecord(filename, destination, time, errorMessage string, success, receive bool) {
	record := &Record{
		FileName:     filename,
		Destination:  destination,
		Time:         time,
		Success:      success,
		ErrorMessage: errorMessage,
		Receive:      receive,
	}
	addRecord(record)
}

func GetHistory() []Record {
	mutex.Lock()
	defer mutex.Unlock()
	return transferHistory
}

func LoadHistory(history string) error {
	data, err := ioutil.ReadFile(history)
	if os.IsNotExist(err) {
		logger.Log.Error("History file is not exist", zap.Error(err))
		return fmt.Errorf("History file is not exist, %s\n", err)
	} else if err != nil {
		logger.Log.Error("Error reading history", zap.Error(err))
		return fmt.Errorf("Error reading history, %s\n", err)
	}
	err = json.Unmarshal(data, &transferHistory)
	if err != nil {
		logger.Log.Error("Error marshaling history", zap.Error(err))
		return fmt.Errorf("Error marshaling history, %s\n", err)
	}
	return nil
}

func addRecord(record *Record) {
	mutex.Lock()
	defer mutex.Unlock()
	transferHistory = append(transferHistory, *record)
	saveHistory()
}

func saveHistory() {
	data, err := json.Marshal(transferHistory)
	if err != nil {
		logger.Log.Error("Error marshaling history", zap.Error(err))
		return
	}
	err = ioutil.WriteFile(config.AppConfig.HistoryFilePath, data, 0644)
	if err != nil {
		logger.Log.Error("Error writing history", zap.Error(err))
	}
}
