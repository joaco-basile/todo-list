package cmd

import (
	"todo_list/models"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addTodoCommand)
}

var addTodoCommand = &cobra.Command{
	Use:   "add",
	Short: "añade una tarea ej: add Title Description",
	Run: func(cmd *cobra.Command, args []string) {
		models.WriteTodo(args)
	},
}
