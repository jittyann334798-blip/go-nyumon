/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type DogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "dog-cli",
		Short: "Dog API CLI",
	}

	var randomCmd = &cobra.Command{
		Use:   "random",
		Short: "ランダムな犬画像を取得",
		Run: func(cmd *cobra.Command, args []string) {

			resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			var dog DogResponse
			if err := json.NewDecoder(resp.Body).Decode(&dog); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(dog.Message)
		},
	}

	rootCmd.AddCommand(randomCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
