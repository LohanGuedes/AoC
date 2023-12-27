package year2023

import (
	"os"

	"github.com/lohanguedes/aoc/cmd/year2023/day1"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "Solutions for 2023",
	Long:  `go run main.go 2023 day1`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := Cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Cmd.AddCommand(day1.Cmd)
}
