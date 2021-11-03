package main

import (
	"sync"

	"code.loda.com/GoStudy/day06/logChan"
)

var bufRC = make(chan *logChan.Buffer, 1024)

func bufChan(fl *logChan.FileLog, buf chan<- *logChan.Buffer) {
	defer wg.Done()
	//一直向通道中写入信息
	for {
		fl.Debug(buf, "today 是星晴二%s", "20211005")
		fl.Trace(buf, "today 是星晴二%s", "20211005")
		fl.Info(buf, "today 是星晴二%s", "20211005")
		fl.Warning(buf, "today 是星晴二%s", "20211005")
		fl.Error(buf, "today 是星晴二%s", "20211005")
		fl.Fatal(buf, "today 是星晴二%s", "20211005")
		// time.Sleep(time.Second)
	}

}

func writeChan(fl *logChan.FileLog, buf <-chan *logChan.Buffer) {
	defer wg.Done()
	//一直从通道中读取信息
	for {
		fl.WriteIntoFile(buf)
	}
}

var wg sync.WaitGroup

func main() {
	lc := logChan.NewFileLog("debug", "./", "mylog.txt", 10*1024*1024)

	wg.Add(1)
	go bufChan(lc, bufRC)

	wg.Add(1)
	go writeChan(lc, bufRC)

	wg.Wait()
}
