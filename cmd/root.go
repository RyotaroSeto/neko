/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const argsCountOne = 1

var nekoCmd = &cobra.Command{
	Use:   "neko",
	Short: "show cat",
	Long:  `show cat terminal.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		isArgsVersion, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}
		if isArgsVersion {
			showVersion()
			return nil
		}

		fmt.Println(simpleNeko.Simple)
		return nil
	},
}

func Execute() {
	err := nekoCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	nekoCmd.AddCommand(versionCmd)
	nekoCmd.AddCommand(nameCmd)
	nekoCmd.Flags().BoolP("version", "v", false, "")
}
