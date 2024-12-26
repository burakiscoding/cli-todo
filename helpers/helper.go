package helpers

import (
	"os"

	"github.com/burakiscoding/cli-todo/models"
	"github.com/jedib0t/go-pretty/table"
)

func WriteTodos(todos []models.Todo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Titled", "Checked"})
	for _, todo := range todos {
		t.AppendRow(table.Row{todo.Id, todo.Title, todo.Checked})
	}
	t.Render()
}

func WriteTodo(todo models.Todo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Titled", "Checked"})
	t.AppendRow(table.Row{todo.Id, todo.Title, todo.Checked})
	t.Render()
}
