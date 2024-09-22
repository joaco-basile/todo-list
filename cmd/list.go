package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"todo_list/models"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTodo)
}

var listTodo = &cobra.Command{
	Use:   "ls",
	Short: "Lista las tareas",
	Run: func(cmd *cobra.Command, _ []string) {
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintln(writer, "ID\tTitle\tDescription\tIsComplete")
		todos := models.ReadTodos()
		for _, todo := range todos {
			fmt.Fprintf(writer, "%v\t%v\t%v\t%v\n", todo.Id, todo.Title, todo.Description, todo.IsComplete)
		}
		writer.Flush()
	},
}
