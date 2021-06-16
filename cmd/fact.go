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

	"github.com/Ocelani/perdat/internal"
	"github.com/spf13/cobra"
)

var (
	factNewF    bool
	factListF   bool
	factEditF   bool
	factRemoveF bool
)

// fact represents the commands for Fact entity.
var fact = &cobra.Command{
	Use:   "fact",
	Short: "any default life event to be registered",
	Example: fmt.Sprintf(
		"\t%s\n\t%s",
		"perdat fact --new headache",
		"perdat fact -n cinema --date 27/12/20", // TODO
	),
	Run: func(cmd *cobra.Command, args []string) {
		var op string
		switch {
		case factNewF:
			op = "new"
		case factListF:
			op = "list"
		case factEditF:
			op = "edit"
		case factRemoveF:
			op = "remove"
		default:
			op = ""
		}
		internal.FactHandler(op, args)
	},
}

func init() {
	fact.Flags().BoolVarP(&factNewF, "new", "n", factNewF, "add new facts")
	fact.Flags().BoolVarP(&factListF, "list", "l", factListF, "list facts")
	fact.Flags().BoolVarP(&factEditF, "edit", "e", factEditF, "edit a fact")
	fact.Flags().BoolVarP(&factRemoveF, "remove", "r", factRemoveF, "remove facts")
	root.AddCommand(fact)
}
