package main

import (
	"fmt"
	"os"
	// "encoding/csv"
	// "fmt"
	// "os"
	// "strconv"
	// "text/tabwriter"
	// "github.com/spf13/cobra"
	// "github.com/mergestat/timediff"
	// "time"
)

func about() {
	fmt.Println("A todo list for the terminal")

	fmt.Println("Usage:")
	fmt.Println("  tasks [command]")
	fmt.Println("")
	fmt.Println("Available Commands:")
	fmt.Println("  add		    Add a new task to the todo list")
	fmt.Println("  complete	    Set a task as being completed")
	fmt.Println("  completions  Generate the autocompletion script for the specified shell")
	fmt.Println("  delete		Removes a task for the todo list by it's id")
	fmt.Println("  help		    Help about any command")
	fmt.Println("  list		    Lists all of the tasks in your todo list")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -h, --help     Help for tasks")
	fmt.Println("  -t, --toggle   Help message for toggle")
	fmt.Println("")
	fmt.Println("Use \"tasks [command] --help\" for more information about a command")
}

func add(args []string) {

}

func complete(args []string) {

}

func completion(args []string) {

}

func delete(args []string) {

}

func help(args []string) {

}

func list(args []string) {

}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./tasks <command>")
		fmt.Println("Use \"tasks --about\" for information about the app")
		return
	}

	switch args[1] {
	case "--about":
		about()
	case "add":
		add(args[2:])
	case "complete":
		complete(args[2:])
	case "completions":
		completion(args[2:])
	case "delete":
		delete(args[2:])
	case "help":
		help(args[2:])
	case "list":
		list((args[2:]))
	default:
		fmt.Printf("Invalid args: %s\n", args[1])
		fmt.Println("Type --about for more info about the app")
		return
	}
}
