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

type DogResponseMulti struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

func main() {

	var image int

	var rootCmd = &cobra.Command{
		Use: "dog-cli",
	}

	var randomCmd = &cobra.Command{
		Use:   "random",
		Short: "ランダムな犬画像を取得",
		Run: func(cmd *cobra.Command, args []string) {

			var url string

			if image <= 1 {
				url = "https://dog.ceo/api/breeds/image/random"
			} else {
				url = fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%d", image)
			}

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			if image <= 1 {
				var dog DogResponse
				if err := json.NewDecoder(resp.Body).Decode(&dog); err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(dog.Message)
			} else {
				var dogs DogResponseMulti
				if err := json.NewDecoder(resp.Body).Decode(&dogs); err != nil {
					fmt.Println(err)
					return
				}
				for _, img := range dogs.Message {
					fmt.Println(img)
				}
			}
		},
	}

	// 🔥 ここを images に変更
	randomCmd.Flags().IntVarP(&image, "images", "i", 1, "取得する画像の件数")

	rootCmd.AddCommand(randomCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
