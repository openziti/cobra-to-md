/*
Copyright NetFoundry, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	gendoc "github.com/openziti/cobra-to-md"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "example command root",
		Example: "your example here",
		Aliases: []string{"root", "main"},
		Long:    "this is a long example",
	}

	docCmd := gendoc.NewGendocCmd(rootCmd)
	// docCmd.Hidden = false -- here is how you'd show the gendoc command exists
	rootCmd.AddCommand(docCmd)
	sub1 := newSubCommand("subcmd1")
	sub1sub1 := newSubCommand("sub1sub1")
	sub1sub2 := newSubCommand("sub1sub2")
	sub1.AddCommand(sub1sub1)
	sub1.AddCommand(sub1sub2)
	rootCmd.AddCommand(sub1)

	sub2 := newSubCommand("subcmd2")
	sub2sub1 := newSubCommand("sub1sub1")
	sub2sub2 := newSubCommand("sub1sub2")
	sub2.AddCommand(sub2sub1)
	sub2.AddCommand(sub2sub2)
	rootCmd.AddCommand(sub2)

	sub3 := newSubCommand("subcmd3")
	rootCmd.AddCommand(sub3)

	e := rootCmd.Execute()
	if e != nil {
		logrus.Error(e)
	}
}

func newSubCommand(text string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     text + " subcommand",
		Example: text + " your example here",
		Aliases: []string{"subcmd." + text, "subcommand." + text},
		Long:    text + " this is a long example",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(text)
		},
	}
	return cmd
}
