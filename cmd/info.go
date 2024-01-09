package cmd

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/typomedia/mssql/app/structs"
	"log"
	"regexp"

	_ "github.com/denisenkom/go-mssqldb"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show server information",
	Long:  `Show server information in the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		user, _ := cmd.Flags().GetString("user")
		pass, _ := cmd.Flags().GetString("pass")
		port, _ := cmd.Flags().GetInt("port")

		dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, pass, port)

		// connect to database
		db, err := sql.Open("sqlserver", dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var result string

		// get server version
		query := fmt.Sprintf("SELECT @@VERSION")
		err = db.QueryRow(query).Scan(&result)
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`Microsoft SQL Server \d{4}`)
		version := re.FindString(result)

		// get server edition
		query = fmt.Sprintf("SELECT SERVERPROPERTY('Edition')")
		err = db.QueryRow(query).Scan(&result)
		if err != nil {
			log.Fatal(err)
		}

		mssql := structs.Mssql{
			Version: version,
			Edition: result,
		}

		headerFmt := color.New(color.Underline).SprintfFunc()

		tbl := table.New("Version", "Edition")
		tbl.WithHeaderFormatter(headerFmt)
		tbl.AddRow(mssql.Version, mssql.Edition)

		tbl.Print()
	},
}

func init() {
	infoCmd.Flags().StringP("server", "s", "localhost", "specifiy server")
	infoCmd.Flags().IntP("port", "p", 1433, "specifiy port")
	infoCmd.Flags().StringP("user", "u", "", "specifiy user")
	infoCmd.Flags().StringP("pass", "w", "", "specifiy password")
}
