package win32

//#include<Windows.h>
import "C"
import "log"

// DWord 双字
type DWord C.DWORD

// Handle 对象句柄
// Handle 是代表系统的内核对象，如文件句柄，线程句柄，进程句柄
type Handle C.HANDLE

// HWindow 窗口句柄
// HWindow 是线程相关的，你可以通过HWindow 找到该窗口所属进程和线程
type HWindow C.HWND

// HModule 模块句柄
// HModule 是代表应用程序载入的模块，win32系统下通常是被载入模块的线性地址
type HModule C.HMODULE

// HInstance 实例句柄
// HInstance 的本质是模块基地址，他仅仅在同一进程中才有意义
type HInstance C.HINSTANCE

func init() {
	log.Println("Windows.h loaded")
}

// TestWindows 测试Windows 包
func TestWindows() {
	log.Println("测试Windows.h")
}
