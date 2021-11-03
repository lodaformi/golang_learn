package main

import (
	"errors"

	"code.loda.com/GoStudy/day05/mylogger"
)

func main() {
	//console log
	logger := mylogger.NewLogger("info")
	logger.Debug("这是一条Debug级别的日志信息")
	logger.Trace("这是一条Trace级别的日志信息")
	logger.Warning("这是一条Warning级别的日志信息")
	err := errors.New("something wrong")
	logger.Info("这是一条Info级别的日志信息, %v", err)
	logger.Error("这是一条error级别的日志信息")
	logger.Fatal("这是一条Fatal级别的日志信息")

	//file log
	fileLog := mylogger.NewFileLog("info", "./", "normal", 10*1024*1024)
	for {
		fileLog.Debug("这是一条Debug级别的日志信息")
		fileLog.Trace("这是一条Trace级别的日志信息")
		fileLog.Warning("这是一条Warning级别的日志信息")
		err = errors.New("something wrong")
		fileLog.Info("这是一条Info级别的日志信息, %v", err)
		fileLog.Error("这是一条error级别的日志信息")
		fileLog.Fatal("这是一条Fatal级别的日志信息")
		// time.Sleep(time.Second)
	}

}
