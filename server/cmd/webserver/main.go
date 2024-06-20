package main

import (
	"go_chi_template/config"
	"go_chi_template/internal/webserver"
)

func main() {
	a := config.NewApp()
	ws := webserver.NewWebserver(a)

	ws.Start()
}
