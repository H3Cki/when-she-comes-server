package main

import (
	"fmt"
	"time"

	"github.com/H3Cki/wscsrv/control"
)

func main() {
	mouse := &control.Mouse{}

	for {
		time.Sleep(2 * time.Second)

		fmt.Println("scrollin down")
		mouse.ScrollDown(1)
		fmt.Println("scrollin down ok")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("scrollin up")
		mouse.ScrollUp(1)
		fmt.Println("scrollin up ok")
	}
}
