// Package win32 一些常用的Win32 API
// References:
//   https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/
package win32

//#include<Windows.h>
//#include<TlHelp32.h>
import "C"

// HeapEntry32 描述正在检查的堆的一个条目
type HeapEntry32 C.HEAPENTRY32

// HeapList32 描述列表中的条目，该条目枚举指定进程使用的堆
type HeapList32 C.HEAPLIST32

// ModuleEntry32 描述属于指定进程的模块列表中的条目
type ModuleEntry32 C.MODULEENTRY32

// ModuleEntry32W 描述属于指定进程的模块列表中的条目
type ModuleEntry32W C.MODULEENTRY32W

// ProcessEntry32 从快照列表中描述驻留在系统地址空间中的进程列表中的一项
type ProcessEntry32 C.PROCESSENTRY32

// ProcessEntry32W 从快照列表中描述驻留在系统地址空间中的进程列表中的一项
type ProcessEntry32W C.PROCESSENTRY32W

// ThreadEntry32 描述在拍摄快照时系统中正在执行的线程列表中的一项。
type ThreadEntry32 C.THREADENTRY32

// CreateToolhelp32Snapshot 拍摄指定进程以及这些进程使用的堆，模块和线程的快照。
func CreateToolhelp32Snapshot(dwFlags DwFlags, th32ProcessID DWord) Handle {
	return Handle(C.CreateToolhelp32Snapshot(C.DWORD(dwFlags), C.DWORD(th32ProcessID)))
}

// DwFlags 标志位
type DwFlags DWord

const (
	// Th32csInherit 指示快照句柄是可继承的。
	Th32csInherit DwFlags = 0x80000000
	// Th32csSnapAll 包括系统中的所有进程和线程，以及th32ProcessID中指定的进程的堆和模块
	Th32csSnapAll DwFlags = 0x80000000
	// Th32csSnapHeapList 包括系统中的所有进程和线程，以及th32ProcessID中指定的进程的堆和模块
	Th32csSnapHeapList DwFlags = 0x00000001
	// Th32csSnapModule 包括快照中th32ProcessID中指定的进程的所有模块
	Th32csSnapModule DwFlags = 0x00000008
	// Th32csSnapModule32 从64位进程调用时，包括快照中th32ProcessID中指定的进程的所有32位模块
	Th32csSnapModule32 DwFlags = 0x00000010
	// Th32csSnapProcess 在快照中包括系统中的所有进程
	Th32csSnapProcess DwFlags = 0x00000002
	// Th32csSnapThread 在快照中包括系统中的所有线程。
	Th32csSnapThread DwFlags = 0x00000004
)
