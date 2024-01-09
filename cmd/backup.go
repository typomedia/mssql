package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a database",
	Long:  `Backup a database to a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		user, _ := cmd.Flags().GetString("user")
		pass, _ := cmd.Flags().GetString("pass")
		port, _ := cmd.Flags().GetInt("port")
		database, _ := cmd.Flags().GetString("database")

		dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, pass, port, database)

		db, err := sql.Open("sqlserver", dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// set backup file name
		name := fmt.Sprintf("%s.bak", database)

		if len(args) > 0 {
			name = args[0]
		}

		// backup database
		query := fmt.Sprintf("BACKUP DATABASE %s TO DISK='%s' WITH INIT, FORMAT", database, name)
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}

		var result string

		// get backup file path
		query = fmt.Sprintf("SELECT TOP 1 physical_device_name FROM msdb.dbo.backupmediafamily ORDER BY media_set_id DESC")
		row := db.QueryRow(query)
		err = row.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)

		fmt.Println("Database backup completed successfully!")

		// return exit code 0
		os.Exit(0)
	},
}

func init() {
	backupCmd.Flags().StringP("server", "s", "localhost", "specifiy server")
	backupCmd.Flags().IntP("port", "p", 1433, "specifiy port")
	backupCmd.Flags().StringP("user", "u", "", "specifiy user")
	backupCmd.Flags().StringP("pass", "w", "", "specifiy password")
	backupCmd.Flags().StringP("database", "d", "", "specifiy database")

	backupCmd.MarkFlagRequired("user")
	backupCmd.MarkFlagRequired("pass")
	backupCmd.MarkFlagRequired("database")

	// set args
	backupCmd.Args = cobra.MinimumNArgs(0)
	backupCmd.Args = cobra.MaximumNArgs(1)

	//backupCmd.SetUsageTemplate(
	//	//pflag.PrintDefaults() +
	//	fmt.Sprintf("Usage:\n  %s %s [flags] [path]\n", app.App.Name, backupCmd.Name()),
	//)
}
