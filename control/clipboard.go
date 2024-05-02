package control

import (
	"fmt"

	"github.com/H3Cki/wscsrv"
	"golang.design/x/clipboard"
)

type Clipboard struct{}

func NewClipboard() *Clipboard {
	c := &Clipboard{}
	if err := clipboard.Init(); err != nil {
		panic(err)
	}
	return c
}

// func (c *Clipboard) Copy() error  {

// }

// func (c *Clipboard) Paste() error {
// 	content, err := c.GetContent()
// 	if err != nil {

// 	}
// 	robotgo.PasteStr()
// }

func (c *Clipboard) GetContent() (wscsrv.ClipboardContent, error) {
	imgBytes := clipboard.Read(clipboard.FmtImage)
	if imgBytes != nil {
		return wscsrv.ClipboardContent{
			Type:    wscsrv.ClipboardContentTypeImage,
			Content: imgBytes,
		}, nil
	}

	textBytes := clipboard.Read(clipboard.FmtText)
	if textBytes == nil {
		textBytes = []byte{}
	}
	return wscsrv.ClipboardContent{
		Type:    wscsrv.ClipboardContentTypeText,
		Content: textBytes,
	}, nil
}

func (c *Clipboard) SetContent(content wscsrv.ClipboardContent) error {
	var format clipboard.Format
	switch content.Type {
	case wscsrv.ClipboardContentTypeText:
		format = clipboard.FmtText
	case wscsrv.ClipboardContentTypeImage:
		format = clipboard.FmtImage
	default:
		return fmt.Errorf("unsupported format: %v", content.Type)
	}

	clipboard.Write(format, content.Content)
	return nil
}
