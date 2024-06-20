package main

import (
	"go_chi_template/config"
	"go_chi_template/internal/worker"
)

func main() {
	a := config.NewApp()
	w := worker.NewWorker(a)
	defer a.Queue().Client.Close()

	w.Start()

}
