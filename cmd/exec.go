package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a statement",
	Long:  `Execute a statement in the console.`,
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

		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	execCmd.Flags().StringP("server", "s", "localhost", "specifiy server")
	execCmd.Flags().IntP("port", "p", 1433, "specifiy port")
	execCmd.Flags().StringP("user", "u", "", "specifiy user")
	execCmd.Flags().StringP("pass", "w", "", "specifiy password")
	execCmd.Flags().StringP("query", "q", "", "specifiy query")
}
