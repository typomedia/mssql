package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all databases",
	Long:  `List all databases in the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		user, _ := cmd.Flags().GetString("user")
		pass, _ := cmd.Flags().GetString("pass")
		port, _ := cmd.Flags().GetInt("port")

		dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, pass, port)

		db, err := sql.Open("sqlserver", dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// get server databases
		query := fmt.Sprintf("SELECT name, filename FROM master.dbo.sysdatabases WHERE name NOT IN ('master', 'tempdb', 'model', 'msdb')")
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		headerFmt := color.New(color.Underline).SprintfFunc()

		tbl := table.New("Name", "Source")
		tbl.WithHeaderFormatter(headerFmt)

		for rows.Next() {
			var name string
			var filename string
			err = rows.Scan(&name, &filename)
			if err != nil {
				log.Fatal(err)
			}

			tbl.AddRow(name, filename)
		}

		query = fmt.Sprintf("SELECT name, data_source FROM master.sys.servers WHERE is_linked = 1")
		rows, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			var name string
			var filename string
			err = rows.Scan(&name, &filename)
			if err != nil {
				log.Fatal(err)
			}

			tbl.AddRow(name, filename)
		}

		tbl.Print()
	},
}

func init() {
	listCmd.Flags().StringP("server", "s", "localhost", "specifiy server")
	listCmd.Flags().IntP("port", "p", 1433, "specifiy port")
	listCmd.Flags().StringP("user", "u", "", "specifiy user")
	listCmd.Flags().StringP("pass", "w", "", "specifiy password")
}
