package main

import (
	"fmt"
	"os"
	"task-tracker/cmd"
	"task-tracker/internal/storage"
)

func main() {
	if err := storage.Initialize(); err != nil {
		fmt.Println("Ошибка инициализации:", err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		cmd.StartREPL()
	} else {
		command := os.Args[1]
		args := os.Args[2:]
		cmd.Router(command, args)
	}
}
