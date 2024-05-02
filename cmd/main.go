package main

import (
	"fmt"
	"runtime"

	"github.com/H3Cki/wscsrv/control"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	mouse := &control.Mouse{}
	for {
		fmt.Println(mouse.Pointer())
	}

}
