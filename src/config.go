package main

import (
	"fmt"
	"github.com/go-ini/ini" // INI reader (im lazy)
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Config struct {
	Resx    int
	Resy    int
	Ratio   int
	Rate    int
	Vsfx    int
	Vbgm    int
	P1up    sdl.Keycode
	P1down  sdl.Keycode
	P1left  sdl.Keycode
	P1right sdl.Keycode
	P1a     sdl.Keycode
	P1b     sdl.Keycode
	P1c     sdl.Keycode
	P1d     sdl.Keycode
}

var cnf Config

func Load_config(name string) Config {
	cfg, err := ini.Load(name)
	if err != nil {
		fmt.Printf("Fail to read configuration: %v", err)
		os.Exit(1)
	}
	cnf.Resx = cfg.Section("video").Key("res_x").MustInt()
	cnf.Resy = cfg.Section("video").Key("res_y").MustInt()
	cnf.Vbgm = cfg.Section("audio").Key("vol_bgm").MustInt()
	cnf.Vsfx = cfg.Section("audio").Key("vol_sfx").MustInt()
	cnf.Rate = cfg.Section("audio").Key("rate").MustInt()
	cnf.P1up = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_up").String())
	cnf.P1down = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_down").String())
	cnf.P1left = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_left").String())
	cnf.P1right = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_right").String())
	cnf.P1a = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_A").String())
	cnf.P1b = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_B").String())
	cnf.P1c = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_C").String())
	cnf.P1d = sdl.GetKeyFromName(cfg.Section("controls").Key("p1_D").String())
	cnf.Ratio = -1
	switch float32(cnf.Resx) / float32(cnf.Resy) {
	case 4.0 / 3.0:
		cnf.Ratio = 0
	case 16.0 / 9.0:
		cnf.Ratio = 1
	default:
		println("Invalid aspect ratio. Check your resolution settings")
		os.Exit(2)
	}
	return cnf
}

func Get_config() Config {
	return cnf
}
