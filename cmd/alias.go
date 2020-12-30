/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/miyazi777/snippet-man/snippet"
	"github.com/miyazi777/snippet-man/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func alias(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Failed exec. Parameter is required.")
	}
	alias := strings.Join(args, " ")

	var snippets snippet.Snippets
	snippets.Load()

	snippet := snippets.AliasFilter(alias)
	if snippet == nil {
		return errors.New("Not found alias.")
	}

	fmt.Printf("%s %s\n", color.GreenString("Command:"), color.HiYellowString(snippet.Command))

	// place holder
	command := inputPlaceholder(snippet.Command)

	// exec command
	fmt.Printf("%s %s\n", color.GreenString("Exec:"), color.HiYellowString(command))
	return util.Run(command, os.Stdin, os.Stdout)
}

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Execute the command with an alias.",
	Long:  "",
	RunE:  alias,
}

func init() {
	rootCmd.AddCommand(aliasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aliasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aliasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
