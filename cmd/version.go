package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typomedia/mssql/app"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version",
	Long:  `Show the current version in the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Logo())
	},
}
