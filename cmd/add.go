/*
Copyright © 2020 perdat - github.com/Ocelani/perdat

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

	"github.com/spf13/cobra"
)

// TODO: perdat add created pull request -d yesterday

// add represents the add command
var add = &cobra.Command{
	Use:   "add",
	Short: "add - Registers a new fact",
	Example: fmt.Sprintf(
		"  %s\n  %s",
		"perdat add headache",
		"perdat add cinema 27/12/20"),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	root.AddCommand(add)
	// * Flags and configuration settings * //
	// Persistent Flags work for this command and all subcommands
	add.PersistentFlags().Args()
	// Local flags which will only run when this command is called directly
	add.Flags().StringP(
		"date", "d", "time.Now()", "registers a datetime of a fact")

}
