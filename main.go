package main

//#include<Windows.h>
import "C"
import (
	"fmt"
	"log"

	"github.com/go-learning/win32"
)

type POINT C.POINT

// 模块初始化函数，在main 之前调用
func init() {
	log.SetPrefix("[INFO]")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("log")
}

func main() {
	// var h win32.Handle
	fmt.Printf("%p\n", win32.CreateToolhelp32Snapshot(win32.Th32csSnapProcess, 0))
	// C.MessageBox(nil, C.CString("H"), C.CString("T"), 0)
	// var curpos C.POINT
	// var pId C.DWORD
	// for {
	// 	C.GetCursorPos(&curpos)
	// 	hWnd := C.HWND(C.WindowFromPoint(curpos))
	// 	C.GetWindowThreadProcessId(hWnd, &pId)
	// 	log.Printf("%d", pId)
	// 	time.Sleep(time.Duration(100) * time.Millisecond)
	// }
}
