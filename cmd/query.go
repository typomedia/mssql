package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Execute a query",
	Long:  `Execute a query in the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		user, _ := cmd.Flags().GetString("user")
		pass, _ := cmd.Flags().GetString("pass")
		port, _ := cmd.Flags().GetInt("port")
		query, _ := cmd.Flags().GetString("query")

		dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, pass, port)

		db, err := sql.Open("sqlserver", dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if len(args) > 0 {
			file := args[0]
			content, err := os.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			query = string(content)
		}

		// get server version
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		format := color.New(color.Underline).SprintfFunc()

		cols, _ := rows.Columns()

		var columnHeaders []interface{}
		for _, col := range cols {
			columnHeaders = append(columnHeaders, col)
		}

		tbl := table.New(columnHeaders...)
		tbl.WithHeaderFormatter(format)

		for rows.Next() {
			// Initialize a slice to store pointers to values in the current row
			var values []interface{}

			// Loop through columns and append pointers to the values slice
			for i := 0; i < len(cols); i++ {
				var value interface{}
				values = append(values, &value)
			}

			// Scan the row values into the pointers
			err = rows.Scan(values...)
			if err != nil {
				log.Fatal(err)
			}

			// Print the values of all columns in the row
			var row []interface{}
			for _, v := range values {
				row = append(row, *v.(*interface{}))
			}
			tbl.AddRow(row...)
		}

		tbl.Print()
	},
}

func init() {
	queryCmd.Flags().StringP("server", "s", "localhost", "specifiy server")
	queryCmd.Flags().IntP("port", "p", 1433, "specifiy port")
	queryCmd.Flags().StringP("user", "u", "", "specifiy user")
	queryCmd.Flags().StringP("pass", "w", "", "specifiy password")
	queryCmd.Flags().StringP("query", "q", "", "specifiy query")
}
