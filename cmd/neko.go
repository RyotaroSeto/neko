/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// nekoCmd represents the neko command
var nekoCmd = &cobra.Command{
	Use:   "neko",
	Short: "A brief description of your command",
	Long:  `A.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("neko called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nekoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nekoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nekoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
