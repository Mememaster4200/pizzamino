package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	// run everything in our main SDL thread
	// (SDL crashes with multiple threads)
	var exitcode int
	sdl.Main(func() {
		exitcode = run()
	})

	os.Exit(exitcode) // pass exit code on return
}
