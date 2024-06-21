package main

import (
	"first_move/config"
	"first_move/internal/worker"
)

func main() {
	a := config.NewApp()
	w := worker.NewWorker(a)
	defer a.Queue().Client.Close()

	w.Start()

}
