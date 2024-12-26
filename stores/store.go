package stores

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"slices"

	"github.com/burakiscoding/cli-todo/models"
)

var errNotFound = errors.New("todo not found")

// Core function to read todos.json file
func readFile() ([]models.Todo, error) {
	b, err := os.ReadFile("todos.json")
	if err != nil {
		return []models.Todo{}, err
	}

	var t []models.Todo
	if err := json.Unmarshal(b, &t); err != nil {
		return []models.Todo{}, err
	}

	return t, nil
}

// Core function to write todos.json file
func writeFile(t []models.Todo) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	if err := os.WriteFile("todos.json", b, fs.FileMode(os.O_WRONLY)); err != nil {
		return err
	}

	return nil
}

func ReadAll() ([]models.Todo, error) {
	return readFile()
}

func ReadById(id int) (models.Todo, error) {
	t, err := readFile()
	if err != nil {
		return models.Todo{}, err
	}

	idx := slices.IndexFunc(t, func(todo models.Todo) bool {
		return todo.Id == id
	})

	if idx == -1 {
		return models.Todo{}, errNotFound
	}

	return t[idx], nil
}

func Create(title string) error {
	t, err := readFile()
	if err != nil {
		return err
	}

	n := len(t)
	id := 0
	if n > 0 {
		id = t[n-1].Id + 1
	}

	t = append(t, models.Todo{Id: id, Title: title})
	return writeFile(t)
}

func Delete(id int) error {
	t, err := readFile()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(t, func(todo models.Todo) bool {
		return todo.Id == id
	})

	if idx == -1 {
		return errNotFound
	}

	t = append(t[:idx], t[idx+1:]...)
	return writeFile(t)
}

func Update(id int, title string, checked bool) error {
	t, err := readFile()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(t, func(todo models.Todo) bool {
		return todo.Id == id
	})

	if idx == -1 {
		return errNotFound
	}

	if title != "" {
		t[idx].Title = title
	}
	t[idx].Checked = checked

	return writeFile(t)
}
