package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_DEFAULT,
		tablewriter.ALIGN_DEFAULT,
	})
	table.SetHeaderColor(
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.BgGreenColor,
		},
		tablewriter.Colors{
			tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor,
		},
		tablewriter.Colors{
			tablewriter.BgRedColor, tablewriter.FgWhiteColor,
		},
	)

	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetFooter([]string{"", "", "427.0"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
