package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

// Main runner
func run() error {
	flag.Parse()
	arguments := flag.Args()
	CheckArgsCount(arguments)

	updatedArgs := SetFlag(arguments)
	updatedPath := BuildNewPath(updatedArgs.path, updatedArgs.port)
	updatedArgs.path = updatedPath

	RunCommand(updatedArgs)
	return nil
}

// New path builder
func BuildNewPath(path string, port string) string {
	slicedPath := strings.Split(path, ":")
	res := fmt.Sprintf("ssh://%s:%s/%s", slicedPath[0], port, slicedPath[1])
	return res
}

// Command runner
func RunCommand(args ArgumentsType) {
	git := "git"

	cmd := exec.Command(git, args.command, args.path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
