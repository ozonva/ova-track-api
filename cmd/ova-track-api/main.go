package main

import (
	"fmt"
	"github.com/ozonva/ova-task-api/internal/utils"
	"os"
)

func main() {
	fmt.Println("Hi, i am ova-track-api!")
	if len (os.Args) != 2{
		fmt.Println("Path to config is strictly required")
		return
	}
	path := os.Args[1]
	utils.InitLibraryFromFile(path)
}
