package mylogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type ConsoleLogger struct {
	Level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("日志级别错误")
	}
}

func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{Level: level}
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	// 如果当前文件大小 大于等于 日志文件的最大值 就应该放回 true
	return fileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) writeLogBackground() {
	for {
		select {
		case logTmp := <-f.logChan:
			if f.checkSize(f.fileObj) {
				neFile, err := f.splitFile(f.fileObj) // 切割日志文件
				if err != nil {
					return
				}
				f.fileObj = neFile
			}
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s \n", logTmp.timestamp, getLogString(logTmp.Level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.Level > ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于等于ERROR级别，还要再errlog日志文件中再记录一编
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志休息
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	// 把日志发送到通道中
	logTmp := &logMsg{
		lv,
		msg,
		funcName,
		fileName,
		lineNo,
		now.Format("2006-01-02 15:04:05"),
	}
	select {
	case f.logChan <- logTmp:
	default:
		// 日志丢了
	}
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {

	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) // 拿到当前的日志文件完整路径

	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, f.fileName, nowStr) // 拼接一个日志文件备份的名字

	// 需要切割日志为你教案
	// 1.关闭当前的日志i文件
	file.Close()
	// 2.备份一下 rename

	os.Rename(logName, newLogName)
	// 3.打开一个新文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err :%v\n", err)
		return nil, err
	}
	// 4.将打开的新日志文件对象复制给 f.fileObj
	return fileObj, err
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
	}
	return runtime.FuncForPC(pc).Name(), path.Base(file), lineNo
}

func (f *FileLogger) Debug(msg string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.log(DEBUG, msg, a...)
	}
}

func (f *FileLogger) Trace(msg string, a ...interface{}) {
	if f.enable(TRACE) {
		f.log(TRACE, msg, a...)
	}
}

func (f *FileLogger) Info(msg string, a ...interface{}) {
	if f.enable(INFO) {
		f.log(INFO, msg, a...)
	}
}

func (f *FileLogger) Warning(msg string, a ...interface{}) {
	if f.enable(WARNING) {
		f.log(WARNING, msg, a...)
	}
}

func (f *FileLogger) Error(msg string, a ...interface{}) {
	if f.enable(ERROR) {
		f.log(ERROR, msg, a...)
	}

}

func (f *FileLogger) Fatal(msg string, a ...interface{}) {
	if f.enable(FATAL) {
		f.log(FATAL, msg, a...)
	}
}
