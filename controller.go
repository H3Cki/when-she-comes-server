package wscsrv

type VolumeController interface {
	SetVolume(vol float64) error
	SetMute(isMuted bool) error
}

type MouseController interface {
	PressLeft() error
	ReleaseLeft() error
	ClickLeft() error
	ClickRight() error
	ScrollUp(something int) error
	ScrollDown(something int) error
	Move(xOffset, yOffset int) error
	SetPosition(x, y int) error
}

type KeyController interface {
	// Alphabet
	Press(byte) error
	Release(byte) error
	SendKey(byte) error
	SendKeys([]byte) error

	// F keys
	SendF1() error
	SendF2() error
	SendF3() error
	SendF4() error
	SendF5() error
	SendF6() error
	SendF7() error
	SendF8() error
	SendF9() error
	SendF10() error
	SendF11() error
	SendF12() error

	// Navigation keys
	SendEscape() error

	SendTilde() error
	SendTab() error
	SendCapslock() error
	SendShift() error
	SendControl() error
	SendAlt() error
	SendEnter() error
	SendBackspace() error

	SendPrintScreen() error
	SendScreenLock() error
	SendPauseBreak() error

	SendInsert() error
	SendHome() error
	SendPageUp() error
	SendPageDown() error
	SendDelete() error
	SendEnd() error

	SendArrowUp() error
	SendArrowDown() error
	SendArrowLeft() error
	SendArrowRight() error
}

type ZoomController interface {
	EnableZoom() error
	DisableZoom() error
	ZoomIn() error
	ZoomOut() error
}

type PointerController interface {
	IsCaret() error
}

type MediaController interface {
	Play() error
	Pause() error
	Next() error
	Previous() error
}

const (
	ClipboardContentTypeText = iota
	ClipboardContentTypeImage
)

type ClipboardContent struct {
	Type int
	Data any
}

type ClipboardController interface {
	Copy() error
	Paste() error
	GetContent() (ClipboardContent, error)
}
