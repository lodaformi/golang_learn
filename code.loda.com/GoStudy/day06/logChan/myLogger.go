package logChan

import (
	"errors"
	"strings"
)

type logLevel uint8

type Buffer struct {
	con string
	lv  logLevel
}

const (
	UNKOWN logLevel = iota
	DEBUG
	TRACE
	WARNING
	INFO
	ERROR
	FATAL
)

func parseLogLevel(lv string) (logLevel, error) {
	s := strings.ToLower(lv)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "warning":
		return WARNING, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("未知日志类型")
		return UNKOWN, err
	}
}

func logLevel2string(lv logLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case WARNING:
		return "WARNING"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}
