package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ArgumentsType struct {
	command string
	path    string
	port    string
}

func main() {
	flag.Parse()
	arguments := flag.Args()
	checkArgsCount(arguments)

	updatedArgs := SetFlag(arguments)
	updatedPath := BuildNewPath(updatedArgs.path, updatedArgs.port)
	updatedArgs.path = updatedPath

	RunCommand(updatedArgs)
}

func checkArgsCount(args []string) {
	if len(args) != 2 {
		fmt.Println("\n\t################\n\t# gitea custom #\n\t################")
		fmt.Println("\nYou have to use two arguments (excluding flags).")
		fmt.Println("Exmaples:")
		fmt.Println("\t1. gitea clone git@domain.com:user/test.git")
		fmt.Println("\t2. gitea -port=22 clone git@domain.com:user/test.git")
		os.Exit(1)
	}
}

func BuildNewPath(path string, port string) string {
	slicedPath := strings.Split(path, ":")
	res := fmt.Sprintf("ssh://%s:%s/%s", slicedPath[0], port, slicedPath[1])
	return res
}

func RunCommand(args ArgumentsType) {
	git := "git"
	cmd := exec.Command(git, args.command, args.path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (obj *ArgumentsType) defaultPort() {
	if obj.port == "" {
		obj.port = "2221"
	}
}

//find flag
func SetFlag(str []string) ArgumentsType {
	res := ArgumentsType{command: str[0], path: str[1]}
	if !port.set {
		res.defaultPort()
	} else {
		res.port = port.value
	}
	return res
}

//  with flag magic
type StringFlag struct {
	set   bool
	value string
}

func (sf *StringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *StringFlag) String() string {
	return sf.value
}

var port StringFlag

func init() {
	flag.Var(&port, "port", "the port")
}
