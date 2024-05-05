package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "plasma-harness",
	Short: "Development harness for Plasma Core",
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
