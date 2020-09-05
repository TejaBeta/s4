build:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/s4 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/s4 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/s4 main.go