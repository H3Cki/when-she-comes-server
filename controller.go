package wscsrv

import "image"

type SystemInfo struct {
	Name              string
	SystemName        string
	NetworkAddress    string
	Displays          []Display
	Volume            int
	VolumeMuted       bool
	ZoomEnabled       bool
	Actions           []Action
	MouseSensitivity  float64
	ScrollSensitivity float64
}

type Display struct {
	Idx    int
	Width  int
	Height int
}

type SystemInfoController interface {
	SystemInfo() (SystemInfo, error)
}

type VolumeController interface {
	GetVolume() (int, error)
	GetMute() (bool, error)
	SetVolumePct(p int) error
	SetMute(isMuted bool) error
}

type MouseController interface {
	PressLeft() error
	ReleaseLeft() error
	ClickLeft() error
	ClickRight() error
	ScrollUp(scrolls int) error
	ScrollDown(scrolls int) error
	MoveBy(x, y int) error
	MoveTo(x, y int) error
	Pointer() (MousePointer, error)
}

type MousePointer int

const (
	MousePointerNormal MousePointer = iota
	MousePointerText
	MousePointerLink
	MousePointerUnknown
)

type KeyboardController interface {
	Type(string) error
}

type ZoomController interface {
	EnableZoom() error
	DisableZoom() error
	ZoomIn() error
	ZoomOut() error
}

type MediaController interface {
	Play() error
	Pause() error
	Stop() error
	Next() error
	Previous() error
}

type ClipboardController interface {
	// Copy() error
	// Paste() error
	SetContent(ClipboardContent) error
	GetContent() (ClipboardContent, error)
}

type ClipboardContentType int

const (
	ClipboardContentTypeText ClipboardContentType = iota
	ClipboardContentTypeImage
)

type ClipboardContent struct {
	Type    ClipboardContentType
	Content []byte
}

type ActionController interface {
	OpenURL(string) error
}

type ScreenController interface {
	Screenshot() (image.Image, error)
}
