package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Devuelve la version",
	Long:  "algo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mi version es 1.0")
	},
}
