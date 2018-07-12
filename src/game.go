package main

const PZM_VER string = "v0.001"

var Conf = Load_config("pizzamino.ini")

func run() int {
	var window, renderer = openwindow(Conf, PZM_VER)
	getevents(Conf, window, renderer)
	return 0
}
