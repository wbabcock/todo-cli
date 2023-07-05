package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/wbabcock/todo-cli/internals"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "Add task to todo list")
	a := flag.Bool("a", false, "Add task to todo list")

	complete := flag.Int("complete", 0, "Complete task at index")
	c := flag.Int("c", 0, "Complete task at index")

	remove := flag.Int("remove", 0, "Remove task at index")
	r := flag.Int("r", 0, "Remove task at index")

	list := flag.Bool("list", false, "List all tasks")
	l := flag.Bool("l", false, "List all tasks")

	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch {
	case *add || *a:
		if len(flag.Args()) < 1 {
			fmt.Println("Missing task")
			os.Exit(0)
		}
		todos.Add(*&flag.Args()[0])
	case *complete > 0:
		if err := todos.Complete(*complete); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case *c > 0:
		if err := todos.Complete(*c); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case *remove > 0:
		if err := todos.Delete(*remove); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case *r > 0:
		if err := todos.Delete(*r); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case *list || *l:
		todos.List()
	default:
		todos.List()
	}

	if err := todos.Save(todoFile); err != nil {
		fmt.Println("Error saving todos")
		os.Exit(1)
	}
}
