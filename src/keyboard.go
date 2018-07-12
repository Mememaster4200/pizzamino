package main

import "github.com/veandco/go-sdl2/sdl" // SDL2 bindings

func keyhandle(t *sdl.KeyboardEvent, conf Config) {
	switch t.Keysym.Sym {
	case conf.P1up:
	case conf.P1down:
	case conf.P1left:
	case conf.P1right:
	case conf.P1a:
	case conf.P1b:
	case conf.P1c:
	case conf.P1d:
		//	default:	println(sdl.GetKeyName(t.Keysym.Sym))
	}
}
