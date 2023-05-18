/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var url string = "https://green.adam.ne.jp/roomazi/cgi-bin/randomname.cgi"

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Naming a cat",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := NewHttpClient(url, http.MethodGet)
		apiResponse, err := client.Execute()
		if err != nil {
			return err
		}

		var englishName string
		if len(apiResponse.Name) > 0 && len(apiResponse.Name[0]) == 3 {
			englishName = apiResponse.Name[0][2]
		}

		fmt.Println(englishName)
		return nil
	},
}
