package cmd

import (
	"log"
	"strconv"
	"todo_list/models"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateTodoCommand)
}

var updateTodoCommand = &cobra.Command{
	Use:   "up",
	Short: "Actualiza el nombre o descripcion de una tarea",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Panic("esto es el update", err)
		}
		models.UpdateTodo(id, args[1], args[2])
	},
}
