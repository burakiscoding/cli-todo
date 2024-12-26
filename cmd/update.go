package cmd

import (
	"fmt"

	"github.com/burakiscoding/cli-todo/stores"
	"github.com/spf13/cobra"
)

var (
	checked   bool
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update",
		Long:  "Update todo",
		Run: func(cmd *cobra.Command, args []string) {
			if id == -1 {
				fmt.Println("You forgot to write id. Example:")
				fmt.Println("./cli-todo update --id=<todo id>")
				return
			} else {
				if err := stores.Update(id, title, checked); err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Updated successfully")
			}
		},
	}
)

func init() {
	updateCmd.Flags().IntVar(&id, "id", -1, "use for identifying todo")
	updateCmd.Flags().StringVar(&title, "title", "", "use for updating title value")
	updateCmd.Flags().BoolVar(&checked, "checked", true, "use for updating checked value")
	rootCmd.AddCommand(updateCmd)
}
