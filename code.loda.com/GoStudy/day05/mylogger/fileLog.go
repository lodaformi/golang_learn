package mylogger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type fileLog struct {
	level         logLevel
	logPath       string
	logName       string
	errLogName    string
	maxSize       uint64
	logFileObj    *os.File
	errLogFileObj *os.File
}

func NewFileLog(lv, path, name string, size uint64) *fileLog {
	lev, err := parseLogLevel(lv)
	if err != nil {
		fmt.Printf("日志类型错误, %v\n", err)
		panic(err)
	}

	tmpF := &fileLog{
		level:         lev,
		logPath:       path,
		logName:       name,
		errLogName:    name + ".err",
		maxSize:       size,
		logFileObj:    nil,
		errLogFileObj: nil,
	}
	err = tmpF.fileInit()
	if err != nil {
		panic(err)
	}
	return tmpF
}

func (f *fileLog) fileInit() error {
	//打开两个日志文件
	fullName := path.Join(f.logPath, f.logName)
	logfile, err := os.OpenFile(fullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stdout, "open log file error, %v", err)
		return err
	}
	errFullName := path.Join(f.logPath, f.errLogName)
	errLogfile, err := os.OpenFile(errFullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stdout, "open err log file error, %v", err)
		return err
	}

	f.logFileObj = logfile
	f.errLogFileObj = errLogfile
	return nil
}

//获取调用该函数的文件信息，函数名、文件名、所在行号
func (f *fileLog) getFileInfo() string {
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

func (f *fileLog) SplitFile(file *os.File, fileName string) (*os.File, error) {
	//根据文件大小进行切割
	oldFullPath := path.Join(f.logPath, fileName)
	timeStamp := time.Now().Format("20060102150405000")

	bakFullPath := fmt.Sprintf("%s.bak-%s", oldFullPath, timeStamp)

	//步骤2：关闭源日志文件
	file.Close()
	//步骤3：备份源日志文件，重命名加上时间戳
	os.Rename(oldFullPath, bakFullPath)
	//步骤4：打开新文件
	newFileObj, err := os.OpenFile(oldFullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return newFileObj, nil
}

func (f *fileLog) checkFileSize(file *os.File) (bool, error) {
	fileStat, err := file.Stat()
	if err != nil {
		return false, err
	}

	//如果大于maxSize
	return fileStat.Size() > int64(f.maxSize), nil
}

func (f *fileLog) dealMsg(funcLv logLevel, msg string, arg ...interface{}) {
	if funcLv >= f.level {
		msg := fmt.Sprintf(msg, arg...)
		fileInfo := f.getFileInfo()
		// fileInfo := fmt.Sprintf("%s %s %d ", funcName, fileName, lineNo)
		s := time.Now().Format("2006-01-02 15:04:05") + " [" + logLevel2string(funcLv) + "] " + fileInfo + msg
		//获取日志文件大小

		//如果大于maxSize
		flag, err := f.checkFileSize(f.logFileObj)
		if err != nil {
			fmt.Fprintf(os.Stdout, "log file stat error, %v\n", err)
			return
		}
		if flag {
			fileObj, err := f.SplitFile(f.logFileObj, f.logName)
			if err != nil {
				fmt.Fprintf(os.Stdout, "open new log file error, %v\n", err)
				return
			}
			f.logFileObj = fileObj
		}

		fmt.Fprintln(f.logFileObj, s)
		if funcLv >= ERROR {
			flag, err := f.checkFileSize(f.errLogFileObj)
			if err != nil {
				fmt.Fprintf(os.Stdout, "log file stat error, %v\n", err)
				return
			}
			if flag {
				errFileObj, err := f.SplitFile(f.errLogFileObj, f.errLogName)
				if err != nil {
					fmt.Fprintf(os.Stdout, "open new log file error, %v\n", err)
					return
				}
				f.errLogFileObj = errFileObj
			}
			fmt.Fprintln(f.errLogFileObj, s)
		}
	}
}

func (f *fileLog) CloseLogFile() {
	f.logFileObj.Close()
	f.errLogFileObj.Close()
}

func (f *fileLog) Debug(msg string, arg ...interface{}) {
	f.dealMsg(DEBUG, msg, arg...)
}

func (f *fileLog) Trace(msg string, arg ...interface{}) {
	f.dealMsg(TRACE, msg, arg...)
}

func (f *fileLog) Warning(msg string, arg ...interface{}) {
	f.dealMsg(WARNING, msg, arg...)
}

func (f *fileLog) Info(msg string, arg ...interface{}) {
	f.dealMsg(INFO, msg, arg...)
}

func (f *fileLog) Error(msg string, arg ...interface{}) {
	f.dealMsg(ERROR, msg, arg...)
}

func (f *fileLog) Fatal(msg string, arg ...interface{}) {
	f.dealMsg(FATAL, msg, arg...)
}
