package app

import "fmt"

type Application struct {
	Name        string
	Version     string
	Author      string
	Description string
	Explanation string
	Banner      string
}

var App = Application{
	Name:        "mssql",
	Version:     "0.2.0",
	Author:      "Philipp Speck <philipp@typo.media>",
	Description: "Microsoft SQL Server CLI",
	Explanation: "Microsoft SQL Server Command Line Interface",
	Banner: `┳┳┓┏┓┏┓┏┓┓ 
┃┃┃┗┓┗┓┃┃┃ 
┛ ┗┗┛┗┛┗┻┗┛`,
}

func Logo() string {
	banner := fmt.Sprintf("%s\n", App.Banner)
	banner += fmt.Sprintf("%s %s\n", App.Name, App.Version)
	banner += App.Author

	return banner
}
