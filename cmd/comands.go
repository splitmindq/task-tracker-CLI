package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker/internal/help"
	"task-tracker/internal/service"
	"task-tracker/internal/storage"
)

func Router(command string, args []string) {
	commands := map[string]func([]string){
		"help":             runHelp,
		"add":              runAdd,
		"list":             runList,
		"delete":           runDelete,
		"mark-done":        runMarkDone,
		"mark-in-progress": runMarkInProgress,
		"update":           runUpdate,
		"list-todo":        runListTodo,
		"list-in-progress": runListInProgress,
		"list-done":        runListDone,
		"pin-task":         runPinTask,
		"unpin-task":       runUnpinTask,
		"list-pinned":      runListPinned,
	}

	if handler, ok := commands[command]; ok {
		handler(args)
	} else {
		fmt.Printf("Неизвестная команда: %s\n", command)
		help.PrintGeneralHelp()
		os.Exit(1)
	}
}

func runHelp(args []string) {
	help.PrintGeneralHelp()
}

func runAdd(args []string) {
	requireArgs(args, 1, "Описание задачи обязательно.\nПример: task-tracker add \"Купить молоко\"")

	task, err := service.NewTask(args[0])
	if err != nil {
		fmt.Println("Ошибка создания задачи:", err)
		os.Exit(1)
	}
	fmt.Printf("Задача добавлена (ID: %d)\n", task.ID)
}

func runList(args []string) {
	service.GetAllTasks()
}

func runDelete(args []string) {
	id := parseIDArg(args)
	if err := service.DeleteTask(id); err != nil {
		fmt.Println("Ошибка удаления задачи:", err)
		os.Exit(1)
	}
}

func runMarkDone(args []string) {
	id := parseIDArg(args)
	if err := service.MarkTaskAsDone(id); err != nil {
		fmt.Println("Ошибка изменения статуса:", err)
		os.Exit(1)
	}
}

func runMarkInProgress(args []string) {
	id := parseIDArg(args)
	if err := service.MarkTaskAsInProgress(id); err != nil {
		fmt.Println("Ошибка изменения статуса:", err)
		os.Exit(1)
	}
}

func runUpdate(args []string) {
	requireArgs(args, 2, "Требуется ID и новое описание задачи.\nПример: task-tracker update 1 \"Купить хлеб\"")

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Ошибка: ID должен быть числом")
		os.Exit(1)
	}

	if err := service.UpdateTask(id, args[1]); err != nil {
		fmt.Println("Ошибка обновления задачи:", err)
		os.Exit(1)
	}
}

func runListTodo(args []string) {
	storage.GetNotDoneTasks()
}

func runListInProgress(args []string) {
	storage.GetInProgressTasks()
}

func runListDone(args []string) {
	storage.GetDoneTasks()
}

func runPinTask(args []string) {
	id := parseIDArg(args)
	if err := service.PinTask(id); err != nil {
		fmt.Println("Ошибка: неверный ID")
		os.Exit(1)
	}

}

func runUnpinTask(args []string) {
	id := parseIDArg(args)
	if err := service.UnpinTask(id); err != nil {
		fmt.Println("Ошибка: неверный ID")
		os.Exit(1)
	}
}

func runListPinned(args []string) {
	service.GetPinnedTasks()
}

// ---------------- Утилиты ----------------

func requireArgs(args []string, count int, usage string) {
	if len(args) < count {
		fmt.Println("Ошибка:", usage)
		os.Exit(1)
	}
}

func parseIDArg(args []string) int {
	requireArgs(args, 1, "Требуется указать ID задачи.")

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Ошибка: ID должен быть числом.")
		os.Exit(1)
	}
	return id
}
func StartREPL() {
	fmt.Println("Task Tracker CLI. Введите команду(help) или 'exit' для выхода.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		command := args[0]
		Router(command, args[1:])
	}
}
