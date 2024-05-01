package control

import (
	"errors"

	"github.com/go-vgo/robotgo"
)

type Mouse struct{}

func (m *Mouse) PressLeft() error {
	return robotgo.MouseDown(robotgo.Left)
}

func (m *Mouse) ReleaseLeft() error {
	return robotgo.MouseUp(robotgo.Left)
}

func (m *Mouse) ClickLeft() error {
	downErr := robotgo.MouseDown(robotgo.Left)
	upErr := robotgo.MouseUp(robotgo.Left)
	return errors.Join(downErr, upErr)
}

func (m *Mouse) ClickRight() error {
	downErr := robotgo.MouseDown(robotgo.Right)
	upErr := robotgo.MouseUp(robotgo.Right)
	return errors.Join(downErr, upErr)
}

func (m *Mouse) ScrollUp(n int) error {
	robotgo.ScrollSmooth(n, 1, 200)
	return nil
}
func (m *Mouse) ScrollDown(n int) error {
	robotgo.ScrollSmooth(-n, 1, 200)
	return nil
}
func (m *Mouse) Move(xOffset, yOffset int) error { return nil }
func (m *Mouse) SetPosition(x, y int) error      { return nil }
