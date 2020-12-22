package main

import "C"
import (
	"log"
	"time"

	"github.com/go-learning/win32"
)

// 模块初始化函数，在main 之前调用
func init() {
	log.SetPrefix("[INFO]")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("Start: ...")
}

func main() {

	for {
		pos, _ := win32.GetCursorPos()
		hWnd := pos.WindowFromPoint()
		pId, tId := hWnd.GetWindowThreadProcessID()
		hWnd.CloseHandle()

		log.Printf("Process Id: %d, Thread Id: %d", pId, tId)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}
