build:
	go mod tidy
	go build -ldflags "-s -w" -o dist/ .

run:
	go mod tidy
	go run .

cross:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/mssql_linux_amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/mssql_macos_amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/mssql_windows_amd64.exe .
