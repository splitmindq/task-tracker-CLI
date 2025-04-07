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
	}
}

func runHelp(args []string) {
	help.PrintGeneralHelp()
}

func runAdd(args []string) {
	if !requireArgs(args, 1, "Описание задачи обязательно.\nПример: task-tracker add \"Купить молоко\"") {
		return
	}

	task, err := service.NewTask(args[0])
	if err != nil {
		fmt.Println("Ошибка создания задачи:", err)
		return
	}
	fmt.Printf("Задача добавлена (ID: %d)\n", task.ID)
}

func runList(args []string) {
	service.GetAllTasks()
}

func runDelete(args []string) {
	id, err := parseIDArg(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = service.DeleteTask(id); err != nil {
		fmt.Println("Ошибка удаления задачи:", err)
	}
}

func runMarkDone(args []string) {
	id, err := parseIDArg(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = service.MarkTaskAsDone(id); err != nil {
		fmt.Println("Ошибка изменения статуса:", err)
	}
}

func runMarkInProgress(args []string) {
	id, err := parseIDArg(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = service.MarkTaskAsInProgress(id); err != nil {
		fmt.Println("Ошибка изменения статуса:", err)
	}
}

func runUpdate(args []string) {
	if !requireArgs(args, 2, "Требуется ID и новое описание задачи.\nПример: update 1 \"Купить хлеб\"") {
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Ошибка: ID должен быть числом")
		return
	}

	if err := service.UpdateTask(id, args[1]); err != nil {
		fmt.Println("Ошибка обновления задачи:", err)
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
	id, err := parseIDArg(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = service.PinTask(id); err != nil {
		fmt.Println("Ошибка: неверный ID")
	}
}

func runUnpinTask(args []string) {
	id, err := parseIDArg(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = service.UnpinTask(id); err != nil {
		fmt.Println("Ошибка: неверный ID")
	}
}

func runListPinned(args []string) {
	service.GetPinnedTasks()
}

// ---------------- Утилиты ----------------

func requireArgs(args []string, count int, usage string) bool {
	if len(args) < count {
		fmt.Println("Ошибка:", usage)
		return false
	}
	return true
}

func parseIDArg(args []string) (int, error) {
	if len(args) < 1 {
		return 0, fmt.Errorf("требуется указать ID задачи")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, fmt.Errorf("ID должен быть числом")
	}
	return id, nil
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
