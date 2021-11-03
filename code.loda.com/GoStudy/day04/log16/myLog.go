package mylog

import (
	"fmt"
	"os"
)

func Trace(x string) {
	if x != "stdout" {
		data := "Trace file"
		os.WriteFile(x, []byte(data), 0644)
		// file, err := os.OpenFile(x, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		// if err != nil {
		// 	fmt.Printf("open file %s error, %v\n", x, err)
		// 	return
		// }
		// defer file.Close()
	} else {
		fmt.Fprintln(os.Stdout, "trace stdout")
	}
}

func Debug() {
	fmt.Fprintln(os.Stdout, "this is debug func")
}

func Warning() {
	fmt.Fprintln(os.Stdout, "this is warning func")
}

func Info() {
	fmt.Fprintln(os.Stdout, "this is info func")
}

func Panic() {
	fmt.Fprintln(os.Stdout, "this is panic func")
}
