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
	habitNewF    bool
	habitListF   bool
	habitEditF   bool
	habitRemoveF bool
)

// habit represents the commands for Habit entity.
var habit = &cobra.Command{
	Use:   "habit",
	Short: "scheduled routine activity to be marked as done",
	Example: fmt.Sprintf(
		"\t%s\n\t%s\n\t%s\n\t%s",
		"perdat habit --new exercise",
		"perdat habit -n sing --daily 3",       // TODO
		"perdat habit -n reviews --weekly 10",  // TODO
		"perdat habit -n lesson --monday 10am", // TODO
	),
	Run: func(cmd *cobra.Command, args []string) {
		var op string
		switch {
		case habitNewF:
			op = "new"
		case habitListF:
			op = "list"
		case habitEditF:
			op = "edit"
		case habitRemoveF:
			op = "remove"
		default:
			op = ""
		}
		internal.FactHandler(op, args)
	},
}

func init() {
	habit.Flags().BoolVarP(&habitNewF, "new", "n", habitNewF, "add new habits")
	habit.Flags().BoolVarP(&habitListF, "list", "l", habitListF, "list habits")
	habit.Flags().BoolVarP(&habitEditF, "edit", "e", habitEditF, "edit a habit")
	habit.Flags().BoolVarP(&habitRemoveF, "remove", "r", habitRemoveF, "remove habits")
	root.AddCommand(habit)
}
