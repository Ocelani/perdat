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

	"github.com/Ocelani/perdat/pkg/counter"
	"github.com/spf13/cobra"
)

var (
	counterNewF    bool
	counterListF   bool
	counterEditF   bool
	counterRemoveF bool
	counterAddF    bool
	counterSubF    bool
)

// counterCmd represents the commands for Counter entity.
var counterCmd = &cobra.Command{
	Use:   "counter",
	Short: "activity to be incremented or decremented like a counter",
	Example: fmt.Sprintf(
		"\t%s", "perdat counter --new 'headache medicine' --set 10", // TODO
	),
	Run: func(cmd *cobra.Command, args []string) {
		var op string
		switch {
		case counterNewF:
			op = "new"

		case counterAddF:
			op = "add"
			up := counter.MakeMapUpdateInt(args)

		case counterSubF:
			op = "sub"
			up := counter.MakeMapUpdateInt(args)

		case counterListF:
			op = "list"
		case counterEditF:
			op = "edit"
		case counterRemoveF:
			op = "remove"
		default:
			op = ""
		}
	},
}

func init() {
	counterCmd.Flags().BoolVarP(&counterNewF, "new", "n", counterNewF, "add new counters")
	counterCmd.Flags().BoolVarP(&counterListF, "list", "l", counterListF, "list counters")
	counterCmd.Flags().BoolVarP(&counterEditF, "edit", "e", counterEditF, "edit a counter")
	counterCmd.Flags().BoolVarP(&counterRemoveF, "remove", "r", counterRemoveF, "remove counters")
	root.AddCommand(counterCmd)
}
