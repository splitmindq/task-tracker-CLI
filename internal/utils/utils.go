package utils

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func SetTableWriterOptions() *tablewriter.Table {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description", "Status", "Created", "Updated"})
	table.SetRowLine(true)
	table.SetBorder(false)

	return table

}
