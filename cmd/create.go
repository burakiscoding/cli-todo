package cmd

import (
	"fmt"

	"github.com/burakiscoding/cli-todo/stores"
	"github.com/spf13/cobra"
)

var (
	title     string
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create",
		Long:  "Create todo",
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" {
				fmt.Println("You forgot to write id. Example:")
				fmt.Println("./cli-todo create --title=\"<todo title>\" ")
				return
			}
			if err := stores.Create(title); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Created successfully")
		},
	}
)

func init() {
	createCmd.Flags().StringVar(&title, "title", "", "use for creating new title")
	rootCmd.AddCommand(createCmd)
}
