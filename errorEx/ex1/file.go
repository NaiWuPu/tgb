package mylogger

import (
	"fmt"
	"os"
	"path"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	line      int
	timestamp string
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 5000),
	}
	err = fl.initFile() // 按照文件路径和文件名打开文件
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err :%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err :%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启一个后台的goroutine 取写日志
	go f.writeLogBackground()
	return nil
}

func (f *FileLogger) Close() {

}
