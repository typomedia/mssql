build:
	go mod tidy
	go build -ldflags "-s -w" -o dist/ .

run:
	go mod tidy
	go run .

compile:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/ .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/ .
