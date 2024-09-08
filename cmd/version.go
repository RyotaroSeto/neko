/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string = "1.1.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of neko",
	Run: func(cmd *cobra.Command, args []string) {
		showVersion()
	},
}

func showVersion() {
	fmt.Println("neko version:", version)
}
