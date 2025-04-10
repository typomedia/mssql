package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database",
	Long:  `Restore a database to a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		user, _ := cmd.Flags().GetString("user")
		pass, _ := cmd.Flags().GetString("pass")
		port, _ := cmd.Flags().GetInt("port")
		database, _ := cmd.Flags().GetString("database")

		dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, pass, port)

		db, err := sql.Open("sqlserver", dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// set backup file name
		file := fmt.Sprintf("%s.bak", database)

		if len(args) > 0 {
			file = args[0]
		}

		fmt.Println(file)

		// backup database
		query := fmt.Sprintf("RESTORE DATABASE %s FROM DISK = '%s' WITH REPLACE", database, file)
		fmt.Println(query)
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Database restored successfully!")

		// return exit code 0
		os.Exit(0)
	},
}

func init() {
	restoreCmd.Flags().StringP("server", "s", "localhost", "specifiy server")
	restoreCmd.Flags().IntP("port", "p", 1433, "specifiy port")
	restoreCmd.Flags().StringP("user", "u", "", "specifiy user")
	restoreCmd.Flags().StringP("pass", "w", "", "specifiy password")
	restoreCmd.Flags().StringP("database", "d", "", "specifiy database")
}
