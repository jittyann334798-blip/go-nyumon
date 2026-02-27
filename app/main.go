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

type BreedListResponse struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

func main() {
	var rootCmd = &cobra.Command{
		Use: "dog-cli",
	}

	var breedListCmd = &cobra.Command{
		Use:   "breed-list",
		Short: "犬種一覧を取得",
		Run: func(cmd *cobra.Command, args []string) {

			resp, err := http.Get("https://dog.ceo/api/breeds/list")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			var result BreedListResponse
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, breed := range result.Message {
				fmt.Println(breed)
			}
		},
	}

	rootCmd.AddCommand(breedListCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
