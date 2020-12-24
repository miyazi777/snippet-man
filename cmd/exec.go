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
	"os"
	"regexp"
	"strings"

	"github.com/miyazi777/snippet-man/snippet"
	"github.com/miyazi777/snippet-man/util"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

type Placeholder struct {
	Match string
	Param string
}

func inputPlaceholder(command string) string {
	re := `<([\S].+?[\S])>`
	r := regexp.MustCompile(re)
	params := r.FindAllStringSubmatch(command, -1)

	var placeholders []Placeholder
	if len(params) > 0 {
		for _, p := range params {
			placeholders = append(placeholders, Placeholder{Match: p[0], Param: p[1]})
		}
	}

	newCommand := command
	for _, placeholder := range placeholders {
		prompt := fmt.Sprintf("%s > ", placeholder.Param)
		input, _ := util.Scan(color.HiYellowString(prompt), true, nil)
		newCommand = strings.Replace(newCommand, placeholder.Match, input, 1)
	}

	return newCommand
}

func exec(cmd *cobra.Command, args []string) error {
	var snippets snippet.Snippets
	snippets.Load()

	tag, _ := cmd.Flags().GetString("tag")

	displaySnippets := snippets.TagFilter(tag)

	// fazzy finder
	idx, err := fuzzyfinder.Find(
		displaySnippets,
		func(i int) string {
			return displaySnippets[i].Command
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			tag := strings.Join(displaySnippets[i].Tag, ",")
			CRString := strings.Repeat("\n", h-10)
			return fmt.Sprintf("%sCommand: %s\nAlias: %s\nTag: %s\nDescription:\n%s",
				CRString,
				displaySnippets[i].Command,
				displaySnippets[i].Alias,
				tag,
				displaySnippets[i].Description,
			)
		}),
	)

	if err != nil {
		return err
	}
	originalCommand := displaySnippets[idx].Command
	fmt.Printf("%s %s\n", color.GreenString("Selected:"), color.HiYellowString(originalCommand))

	// place holder
	command := inputPlaceholder(originalCommand)

	// exec command
	fmt.Printf("%s %s\n", color.GreenString("Exec:"), color.HiYellowString(command))
	return util.Run(command, os.Stdin, os.Stdout)
}

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: exec,
}

func init() {
	execCmd.Flags().StringP("tag", "t", "", "search tag")
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
