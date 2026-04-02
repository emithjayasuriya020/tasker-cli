package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task add <title>")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		markDone(id)
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)
	default:
		printHelp()
	}
}
