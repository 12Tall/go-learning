package win32

//#include<Windows.h>
import "C"
import "log"

// DWord 双字
type DWord C.DWORD

// Bool 布尔值
// type Bool C._Bool

// Point 鼠标
type Point C.POINT

// GetCursorPos 获取鼠标位置
func GetCursorPos() (Point, bool) {
	var point Point
	res := false
	if C.GetCursorPos((*C.POINT)(&point)) == 0 {
		res = true
	}
	return point, res
}

// GetWindowThreadProcessID 获取创建窗口的进程和线程的Id
func (hWnd *HWindow) GetWindowThreadProcessID() (DWord, DWord) {
	var pID DWord
	var tID = DWord(C.GetWindowThreadProcessId(C.HWND(hWnd.handle), (*C.DWORD)(&pID)))
	return pID, tID
}

// WindowFromPoint 获取鼠标下的窗口句柄
func (point *Point) WindowFromPoint() HWindow {
	var hWnd HWindow
	hWnd.handle = C.HANDLE(C.WindowFromPoint(C.POINT(*point)))
	return hWnd
}

// Handle 对象句柄：C.HANDLE。
// Handle 是代表系统的内核对象，如文件句柄，线程句柄，进程句柄
type Handle struct {
	handle C.HANDLE
}

// HWindow 窗口句柄：C.HWND。
// HWindow 是线程相关的，你可以通过HWindow 找到该窗口所属进程和线程
type HWindow struct {
	Handle
}

// HModule 模块句柄：C.HMODULE。
// HModule 是代表应用程序载入的模块，win32系统下通常是被载入模块的线性地址
type HModule struct {
	Handle
}

// HInstance 实例句柄：C.HINSTANCE。
// HInstance 的本质是模块基地址，他仅仅在同一进程中才有意义
type HInstance struct {
	Handle
}

// CloseHandle 关闭句柄
func (h *Handle) CloseHandle() bool {
	return bool(C.CloseHandle((C.HANDLE)(h.handle)) != 0)
}

func init() {
	log.Println("Windows.h loaded")
}
