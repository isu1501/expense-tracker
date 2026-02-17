package main

import (
	"expense-tracker/src/router"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	command := os.Args
	if len(command) < 2 {
		fmt.Println("Please provide a valid command")
		return
	}
	router.SetupRoutes(command)
}
