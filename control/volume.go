package control

import (
	"runtime"

	"github.com/itchyny/volume-go"
)

type Volume struct {
	fChan chan func()
}

func NewVolume() *Volume {
	v := &Volume{
		fChan: make(chan func()),
	}
	go v.run()
	return v
}

func (v *Volume) SetVolumePct(p int) error {
	retC := make(chan error)

	v.runThreadLocked(func() {
		retC <- volume.SetVolume(p)
	})

	return <-retC
}

func (v *Volume) SetMute(mute bool) error {
	retC := make(chan error)

	v.runThreadLocked(func() {
		if mute {
			retC <- volume.Mute()
			return
		}
		retC <- volume.Unmute()
	})

	return <-retC
}

type getVolumeRet struct {
	v int
	e error
}

func (v *Volume) GetVolume() (int, error) {
	retC := make(chan getVolumeRet)

	v.runThreadLocked(func() {
		vol, err := volume.GetVolume()
		retC <- getVolumeRet{vol, err}
	})

	ret := <-retC

	return ret.v, ret.e
}

type getMuteRet struct {
	v bool
	e error
}

func (v *Volume) GetMute() (bool, error) {
	retC := make(chan getMuteRet)

	v.runThreadLocked(func() {
		mute, err := volume.GetMuted()
		retC <- getMuteRet{mute, err}
	})

	ret := <-retC

	return ret.v, ret.e
}

func (v *Volume) runThreadLocked(f func()) {
	v.fChan <- f
}

func (v *Volume) run() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	for f := range v.fChan {
		go f()
	}
}
