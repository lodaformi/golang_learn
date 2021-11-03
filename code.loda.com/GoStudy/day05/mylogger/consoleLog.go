package mylogger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type consoleLog struct {
	level logLevel
}

func NewLogger(lv string) *consoleLog {
	lev, err := parseLogLevel(lv)
	if err != nil {
		fmt.Printf("日志类型错误, %v\n", err)
		panic(err)
	}
	return &consoleLog{level: lev}
}

func getFileInfo() string {
	pc, fileName, lineNo, ok := runtime.Caller(3)
	if !ok {
		// fmt.Println("get file info error.")
		panic("get file info error.")
	}
	//函数名
	funcName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	// fmt.Println("pcName ", pcName)
	//文件名
	fileName = path.Base(fileName)
	// fmt.Println("fileName ", fileName)
	// fmt.Println("lineNo ", lineNo)
	fileInfo := fmt.Sprintf("%s %s %d ", funcName, fileName, lineNo)
	// fmt.Println(fileInfo)
	return fileInfo
}

func (c *consoleLog) dealMsg(funcLv logLevel, msg string, arg ...interface{}) {
	if funcLv >= c.level {
		msg := fmt.Sprintf(msg, arg...)
		fileInfo := getFileInfo()
		s := time.Now().Format("2006-01-02 15:04:05") + " [" + logLevel2string(funcLv) + "] " + fileInfo + msg
		fmt.Fprintln(os.Stdout, s)
	}
}

func (c *consoleLog) Debug(msg string, arg ...interface{}) {
	c.dealMsg(DEBUG, msg, arg...)
}

func (c *consoleLog) Trace(msg string, arg ...interface{}) {
	c.dealMsg(TRACE, msg, arg...)
}

func (c *consoleLog) Warning(msg string, arg ...interface{}) {
	c.dealMsg(WARNING, msg, arg...)
}

func (c *consoleLog) Info(msg string, arg ...interface{}) {
	c.dealMsg(INFO, msg, arg...)
}

func (c *consoleLog) Error(msg string, arg ...interface{}) {
	c.dealMsg(ERROR, msg, arg...)
}

func (c *consoleLog) Fatal(msg string, arg ...interface{}) {
	c.dealMsg(FATAL, msg, arg...)
}
