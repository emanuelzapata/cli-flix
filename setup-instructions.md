install cobra-cli 
go install github.com/spf13/cobra-cli@latest

init go project
go mod init cli-flix

init cobra cli scaffolds
cobra-cli init

run using:
go run main.go

build using 
go build -o cli-flix
then run using
./cli-flix --help

add new command using 
cobra-cli add <command-name>