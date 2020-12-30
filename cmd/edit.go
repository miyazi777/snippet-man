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
	"path/filepath"

	"github.com/miyazi777/snippet-man/util"

	"github.com/spf13/cobra"
)

func edit(cmd *cobra.Command, args []string) error {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "snippet-man")
	fullFilePath := filepath.Join(configDir, "snippets.toml")

	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("Not found $EDITOR.")
	}

	command := fmt.Sprintf("%s %s\n", editor, fullFilePath)
	util.Run(command, os.Stdin, os.Stdout)
	return nil
}

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the snippet in an editor(vi).",
	Long:  "",
	RunE:  edit,
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
