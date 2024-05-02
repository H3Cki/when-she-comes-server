package control

import "github.com/go-vgo/robotgo"

type Keyboard struct{}

func (k *Keyboard) Type(s string) error {
	robotgo.TypeStr(s)
	return nil
}
