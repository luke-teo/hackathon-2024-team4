package main

import (
	"go_chi_template/config"
	"go_chi_template/internal/app/task"
	"go_chi_template/internal/webserver"
	"log"
	"os"
)

func main() {
	command := ""
	if len(os.Args[1:]) > 0 {
		command = os.Args[1]
	}

	a := config.NewApp()

	if command == "routes:list" {
		ws := webserver.NewWebserver(a)
		ws.PrintRoutes()
		return
	}

	if command == "tenant:cleanup" {
		err := task.DispatchTenantCleanup(a, os.Args[2])

		if err != nil {
			log.Panic(err)
		}
		return
	}
}
