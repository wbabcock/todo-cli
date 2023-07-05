package main

import (
	"flag"

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
		panic(err)
	}

	switch {
	case *add || *a:
		if len(flag.Args()) < 1 {
			panic("Missing task")
		}
		todos.Add(*&flag.Args()[0])
	case *complete > 0:
		if err := todos.Complete(*complete); err != nil {
			panic(err)
		}
	case *c > 0:
		if err := todos.Complete(*c); err != nil {
			panic(err)
		}
	case *remove > 0:
		if err := todos.Delete(*remove); err != nil {
			panic(err)
		}
	case *r > 0:
		if err := todos.Delete(*r); err != nil {
			panic(err)
		}
	case *list || *l:
		todos.List()
	default:
		panic("Invalid option")
	}

	if err := todos.Save(todoFile); err != nil {
		panic(err)
	}
}
