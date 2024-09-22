package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tl",
	Short: "tl is a very fast todo list",
	Long: `Todo list es mi proyecto de prueba en go usando cobra para hacer
    aplicaciones de cli de manera sencilla
  `,

	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
