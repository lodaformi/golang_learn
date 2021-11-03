package logChan

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type FileLog struct {
	level         logLevel
	logPath       string
	logName       string
	errLogName    string
	maxSize       uint64
	logFileObj    *os.File
	errLogFileObj *os.File
}

func NewFileLog(lv, path, name string, size uint64) *FileLog {
	lev, err := parseLogLevel(lv)
	if err != nil {
		fmt.Printf("日志类型错误, %v\n", err)
		panic(err)
	}

	tmpF := &FileLog{
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

func (f *FileLog) fileInit() error {
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
func (f *FileLog) getFileInfo() string {
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

func (f *FileLog) SplitFile(file *os.File, fileName string) (*os.File, error) {
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

func (f *FileLog) checkFileSize(file *os.File) (bool, error) {
	fileStat, err := file.Stat()
	if err != nil {
		return false, err
	}

	//如果大于maxSize
	return fileStat.Size() > int64(f.maxSize), nil
}

func (f *FileLog) bufMsg(ch chan<- *Buffer, funcLv logLevel, msg string, arg ...interface{}) {
	if funcLv >= f.level {
		msg := fmt.Sprintf(msg, arg...)
		fileInfo := f.getFileInfo()
		// fileInfo := fmt.Sprintf("%s %s %d ", funcName, fileName, lineNo)
		s := time.Now().Format("2006-01-02 15:04:05") + " [" + logLevel2string(funcLv) + "] " + fileInfo + msg
		//拿到信息s，写到channel中
		tmp := &Buffer{con: s, lv: funcLv}
		ch <- tmp

	}
}

func (f *FileLog) WriteIntoFile(ch <-chan *Buffer) {
	//循环遍历通道，将通道中的信息取出放到s中，将s写入文件
	for v := range ch {
		s := v.con

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

		if v.lv >= ERROR {
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

func (f *FileLog) CloseLogFile() {
	f.logFileObj.Close()
	f.errLogFileObj.Close()
}

func (f *FileLog) Debug(ch chan<- *Buffer, msg string, arg ...interface{}) {
	f.bufMsg(ch, DEBUG, msg, arg...)
}

func (f *FileLog) Trace(ch chan<- *Buffer, msg string, arg ...interface{}) {
	f.bufMsg(ch, TRACE, msg, arg...)
}

func (f *FileLog) Warning(ch chan<- *Buffer, msg string, arg ...interface{}) {
	f.bufMsg(ch, WARNING, msg, arg...)
}

func (f *FileLog) Info(ch chan<- *Buffer, msg string, arg ...interface{}) {
	f.bufMsg(ch, INFO, msg, arg...)
}

func (f *FileLog) Error(ch chan<- *Buffer, msg string, arg ...interface{}) {
	f.bufMsg(ch, ERROR, msg, arg...)
}

func (f *FileLog) Fatal(ch chan<- *Buffer, msg string, arg ...interface{}) {
	f.bufMsg(ch, FATAL, msg, arg...)
}
