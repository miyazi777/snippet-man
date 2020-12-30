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
	"fmt"
	"strings"

	"github.com/miyazi777/snippet-man/snippet"
	"github.com/miyazi777/snippet-man/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Register the command in a snippets.",
	Long:  "",
	RunE:  newCommand,
}

func checkAlias(line string) bool {
	var snippets snippet.Snippets
	snippets.Load()
	s := snippets.AliasFilter(line)
	if s != nil {
		fmt.Println(color.CyanString("Failed add command. Alias is alreay in use."))
		return false
	}
	return true
}

func newCommand(cmd *cobra.Command, args []string) error {
	// input command
	command, err := util.Scan(color.YellowString("Command >>> "), true, nil)
	if err != nil {
		return err
	}

	// input description
	description, err := util.Scan(color.GreenString("description >>> "), true, nil)
	if err != nil {
		return err
	}

	tagsInput, err := util.Scan(color.RedString("tags >>> "), false, nil)
	if err != nil {
		return err
	}

	alias, err := util.Scan(color.CyanString("alias >>> "), false, checkAlias)
	if err != nil {
		return err
	}

	var snippets snippet.Snippets
	snippets.Load()

	tags := strings.Fields(tagsInput)
	newSni := snippet.Snippet{
		Command:     command,
		Description: description,
		Tag:         tags,
		Alias:       alias,
	}
	snippets.Add(newSni)
	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
