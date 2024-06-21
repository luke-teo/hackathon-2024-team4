package main

import (
	"first_move/config"
	"first_move/internal/webserver"
)

func main() {
	a := config.NewApp()
	ws := webserver.NewWebserver(a)

	ws.Start()
}
