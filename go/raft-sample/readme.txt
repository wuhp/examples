go run main.go `mktemp -d` localhost:10001 localhost:10002 localhost:10003
go run main.go `mktemp -d` localhost:10002 localhost:10003 localhost:10001
go run main.go `mktemp -d` localhost:10003 localhost:10001 localhost:10002
