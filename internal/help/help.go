package help

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Command struct {
	Name        string
	Description string
	Usage       string
}

var Commands = []Command{
	{
		Name:        "add",
		Description: "Добавить новую задачу",
		Usage:       "add \"<описание задачи>\"",
	},
	{
		Name:        "update",
		Description: "Обновить описание задачи",
		Usage:       "update <ID> \"<новое описание>\"",
	},
	{
		Name:        "delete",
		Description: "Удалить задачу",
		Usage:       "delete <ID>",
	},
	{
		Name:        "mark-in-progress",
		Description: "Пометить задачу как 'в процессе'",
		Usage:       "mark-in-progress <ID>",
	},
	{
		Name:        "mark-done",
		Description: "Пометить задачу как 'выполнена'",
		Usage:       "mark-done <ID>",
	},
	{
		Name:        "list",
		Description: "Показать все задачи",
		Usage:       "list",
	},
	{
		Name:        "list-todo",
		Description: "Показать невыполненные задачи",
		Usage:       "list-todo",
	},
	{
		Name:        "list-in-progress",
		Description: "Показать задачи в процессе выполнения",
		Usage:       "list-in-progress",
	},
	{
		Name:        "list-done",
		Description: "Показать выполненные задачи",
		Usage:       "list-done",
	},
	{
		Name:        "pin-task",
		Description: "Закрепить задачу",
		Usage:       "pin-task <ID>",
	},
	{
		Name:        "unpin-task",
		Description: "Открепить задачу",
		Usage:       "unpin-task <ID>",
	},
	{
		Name:        "list-pinned",
		Description: "Вывести закрепленные задачи",
		Usage:       "list-pinned",
	},
	{
		Name:        "help",
		Description: "Показать эту справку",
		Usage:       "help [команда]",
	},
	{
		Name:        "exit",
		Description: "Выйти из программы",
		Usage:       "exit",
	},
}

func PrintGeneralHelp() {
	fmt.Println("Task Tracker CLI - Управление задачами\n")
	fmt.Println("Доступные команды:\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Команда", "Описание", "Использование"})
	table.SetBorder(false)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnSeparator(" ")
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	for _, cmd := range Commands {
		table.Append([]string{cmd.Name, cmd.Description, cmd.Usage})
	}

	table.Render()

	fmt.Println("\nПримеры:")
	fmt.Println("  > add \"Купить молоко\"")
	fmt.Println("  > mark-done 3")
	fmt.Println("  > list-todo")

}
