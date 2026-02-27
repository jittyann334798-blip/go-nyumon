package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "dog-cli",
		Short: "Dog API CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dog-cli 実行成功")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
