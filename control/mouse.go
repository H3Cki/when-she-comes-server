package control

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/H3Cki/wscsrv"
	"github.com/go-vgo/robotgo"
)

var libUser32 syscall.Handle
var FuncGetCursorInfo uintptr

func init() {
	libUser32, _ = syscall.LoadLibrary("user32.dll")
	FuncGetCursorInfo, _ = syscall.GetProcAddress(syscall.Handle(libUser32), "GetCursorInfo")
}

type Mouse struct{}

func (m *Mouse) PressLeft() error {
	return robotgo.MouseDown(robotgo.Left)
}

func (m *Mouse) ReleaseLeft() error {
	return robotgo.MouseUp(robotgo.Left)
}

func (m *Mouse) ClickLeft() error {
	robotgo.Click(robotgo.Left)
	return nil
}

func (m *Mouse) ClickRight() error {
	robotgo.Click(robotgo.Right)
	return nil
}

func (m *Mouse) ScrollUp(n int) error {
	robotgo.ScrollDir(n, robotgo.Up)
	return nil
}

func (m *Mouse) ScrollDown(n int) error {
	robotgo.ScrollDir(n, robotgo.Down)
	return nil
}

func (m *Mouse) MoveBy(xOffset, yOffset int) error {
	robotgo.MoveRelative(xOffset, yOffset)
	return nil
}

func (m *Mouse) MoveTo(x, y int) error {
	robotgo.Move(x, y)
	return nil
}

type (
	HANDLE  uintptr
	DWORD   uint32
	LONG    int32
	HCURSOR HANDLE
)

type POINT struct {
	X LONG
	Y LONG
}

type CURSORINFO struct {
	CbSize      DWORD
	Flags       DWORD
	HCursor     HCURSOR
	PTScreenPos POINT
}

func (m *Mouse) Pointer() (wscsrv.MousePointer, error) {
	var curInfo CURSORINFO
	curInfo.CbSize = DWORD(unsafe.Sizeof(curInfo))
	ok, _, err := syscall.SyscallN(FuncGetCursorInfo, uintptr(unsafe.Pointer(&curInfo)), 0, 0)
	if ok == 0 {
		return 0, errors.New(err.Error())
	}

	switch curInfo.HCursor {
	case 65539: //normal
		return wscsrv.MousePointerNormal, nil
	case 65541: //caret
		return wscsrv.MousePointerText, nil
	case 65567: //pointer
		return wscsrv.MousePointerLink, nil
	}

	return wscsrv.MousePointerUnknown, nil
}
