package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl" // SDL2 bindings
	"os"
	"pzmconfig"
	"strings"
)

func openwindow(conf pzmconfig.Config, version string) (*sdl.Window, *sdl.Renderer) {
	var winTitle string = "Pizzamino " + version
	var winWidth, winHeight int32 = int32(conf.Resx), int32(conf.Resy)
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	sdl.Do(func() {
		window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			winWidth, winHeight, sdl.WINDOW_OPENGL)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		panic(err)
	}

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}

	sdl.Do(func() {
		renderer.Clear()
	})

	return window, renderer
}

func getevents(conf pzmconfig.Config, window *sdl.Window, renderer *sdl.Renderer) {

	var event sdl.Event
	var FrameRate uint32 = 60

	running := true
	for running {
		sdl.Do(func() {
			for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					running = false
				case *sdl.KeyboardEvent:
					switch strings.ToUpper(string(t.Keysym.Sym)) {
					case strings.ToUpper(conf.P1up):
						print("Up")
					}
				}
				sdl.Delay(1000 / FrameRate) // pause every frame
				// this does not make a stable precise framerate,
				// but it's better than wasting laptop battery.
			}
		})
	}
}
