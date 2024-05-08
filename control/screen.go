package control

import (
	"image"

	"github.com/go-vgo/robotgo"
)

type Screen struct{}

func (s *Screen) Screenshot() (image.Image, error) {
	img, err := robotgo.Capture()
	if err != nil {
		return nil, err
	}
	return img, nil
}
