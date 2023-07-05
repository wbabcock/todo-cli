package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{Task: task, Done: false, CreatedAt: time.Now()}
	*t = append(*t, todo)
}

func (t *Todos) Complete(i int) error {
	if i > len(*t) || i < 1 {
		return errors.New("Invalid index")
	}

	(*t)[i-1].Done = true
	(*t)[i-1].CompletedAt = time.Now()

	return nil
}

func (t *Todos) Delete(i int) error {
	if i > len(*t) || i < 1 {
		return errors.New("Invalid index")
	}

	*t = append((*t)[:i-1], (*t)[i:]...)

	return nil
}

func (t *Todos) List() {
	// for i, todo := range *t {
	// 	if todo.Done {
	// 		println(i+1, "✅", todo.Task)
	// 	} else {
	// 		println(i+1, "⬜", todo.Task)
	// 	}
	// }

	if len(*t) == 0 {
		println("Nothing to do!")
		return
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Index"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	for i, todo := range *t {
		var done string
		if todo.Done {
			done = "✅"
		} else {
			done = "⬜"
		}

		var completeAt string
		if todo.CompletedAt.IsZero() {
			completeAt = ""
		} else {
			completeAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
		}

		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i+1)},
			{Text: todo.Task},
			{Align: simpletable.AlignCenter, Text: done},
			{Align: simpletable.AlignRight, Text: todo.CreatedAt.Format("2006-01-02 15:04:05")},
			{Align: simpletable.AlignRight, Text: completeAt},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 5, Text: fmt.Sprintf("Total: %d", len(*t))},
		},
	}

	fmt.Println(table.String())
}

func (t *Todos) Save(filename string) error {
	dirname, _ := os.UserHomeDir()
	file, err := os.OpenFile(dirname+"/Documents/"+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	if err := json.NewEncoder(file).Encode(t); err != nil {
		return err
	}

	return nil
}

func (t *Todos) Load(filename string) error {
	dirname, _ := os.UserHomeDir()
	if _, err := os.Stat(dirname+"/Documents/"+filename); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	file, err := ioutil.ReadFile(dirname+filename)
	if err != nil {
		return err
	}

	if len(file) > 0 {
		if err := json.Unmarshal(file, t); err != nil {
			return err
		}
	}

	return nil
}
