package cmd

import (
	"fmt"

	"github.com/burakiscoding/cli-todo/stores"
	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delte",
		Long:  "Delete todo",
		Run: func(cmd *cobra.Command, args []string) {
			if id == -1 {
				fmt.Println("You forgot to write id. Example:")
				fmt.Println("./cli-todo delete --id=<todo id>")
				return
			}
			if err := stores.Delete(id); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Deleted successfully")
		},
	}
)

func init() {
	deleteCmd.Flags().IntVar(&id, "id", -1, "use for identifying todo")
	rootCmd.AddCommand(deleteCmd)
}
