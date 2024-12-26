package cmd

import (
	"fmt"

	"github.com/burakiscoding/cli-todo/helpers"
	"github.com/burakiscoding/cli-todo/stores"
	"github.com/spf13/cobra"
)

var (
	id      int
	readCmd = &cobra.Command{
		Use:   "read",
		Short: "Read",
		Long:  "Reading todos",
		Run: func(cmd *cobra.Command, args []string) {
			if id == -1 {
				todos, err := stores.ReadAll()
				if err != nil {
					fmt.Println(err)
					return
				}
				helpers.WriteTodos(todos)
			} else {
				todo, err := stores.ReadById(id)
				if err != nil {
					fmt.Println(err)
					return
				}
				helpers.WriteTodo(todo)
			}
		},
	}
)

func init() {
	readCmd.Flags().IntVar(&id, "id", -1, "use for identifying todo")
	rootCmd.AddCommand(readCmd)
}
