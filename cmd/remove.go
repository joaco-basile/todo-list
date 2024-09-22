package cmd

import (
	"log"
	"strconv"
	"todo_list/models"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeTodoCommand)
}

var removeTodoCommand = &cobra.Command{
	Use:   "rm",
	Short: "borra una tarea por el id",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		models.DeleteTodo(id)
	},
}
