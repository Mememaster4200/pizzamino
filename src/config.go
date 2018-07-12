package main

import (
	"fmt"
	"github.com/go-ini/ini" // INI reader (im lazy)
	"os"
	"strings"
)

type Config struct {
	Resx    int
	Resy    int
	Rate    int
	Vsfx    int
	Vbgm    int
	P1up    string
	P1down  string
	P1left  string
	P1right string
	P1a     string
	P1b     string
	P1c     string
	P1d     string
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
	cnf.P1up = strings.ToLower(cfg.Section("controls").Key("p1_up").String())
	cnf.P1down = strings.ToLower(cfg.Section("controls").Key("p1_down").String())
	cnf.P1left = strings.ToLower(cfg.Section("controls").Key("p1_left").String())
	cnf.P1right = strings.ToLower(cfg.Section("controls").Key("p1_right").String())
	cnf.P1a = strings.ToLower(cfg.Section("controls").Key("p1_A").String())
	cnf.P1b = strings.ToLower(cfg.Section("controls").Key("p1_B").String())
	cnf.P1c = strings.ToLower(cfg.Section("controls").Key("p1_C").String())
	cnf.P1d = strings.ToLower(cfg.Section("controls").Key("p1_D").String())
	return cnf
}

func Get_config() Config {
	return cnf
}
