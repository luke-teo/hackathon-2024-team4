package main

import (
	"fmt"
	"os"

	"first_move/config"
	"first_move/internal/console"
)

func main() {
	command := ""
	if len(os.Args[1:]) > 0 {
		command = os.Args[1]
	}

	opts := []string{}
	if len(os.Args[1:]) > 1 {
		opts = os.Args[2:]
	}

  fmt.Println(opts)

	app := config.NewApp()
	c := console.NewConsole(app)

	switch command {
	case "task:parse_text_chat_from_csv":
		if err := c.ParseTextChatFromCsv(opts[0]); err != nil {
			panic(err)
		}
	}
}
