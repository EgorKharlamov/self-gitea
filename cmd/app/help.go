package main

import (
	"fmt"
	"os"
)

// Check for arguments count and print help
func CheckArgsCount(args []string) {
	if len(args) != 2 {
		fmt.Println("\n\t################\n\t# gitea custom #\n\t################")
		fmt.Println("\nYou have to use two arguments (excluding flags).")
		fmt.Println("Exmaples:")
		fmt.Println("\t1. gitea clone git@domain.com:user/test.git")
		fmt.Println("\t2. gitea -port=22 clone git@domain.com:user/test.git")
		os.Exit(1)
	}
}
