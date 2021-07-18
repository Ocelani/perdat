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
	"os"

	"github.com/Ocelani/perdat/internal"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	newF    bool
	listF   bool
	editF   bool
	removeF bool
)

// root represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:   "pdt",
	Short: "Know yourself",
	Long: `
	A CLI tool that registers information about your daily life privately.
	The user provides an input and the app stores in a sqlite file.
	With that data, you can get some insights about yourself.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var op string
		switch {
		case newF:
			op = "new"
		case listF:
			op = "list"
		case editF:
			op = "edit"
		case removeF:
			op = "remove"
		default:
		}
		internal.FactHandler(op, args)
	},
}

func Execute() {
	root.Flags().BoolVarP(&newF, "new", "n", newF, "add new facts")
	root.Flags().BoolVarP(&listF, "list", "l", listF, "list facts")
	root.Flags().BoolVarP(&editF, "edit", "e", editF, "edit a fact")
	root.Flags().BoolVarP(&removeF, "remove", "r", removeF, "remove facts")
	cobra.CheckErr(root.Execute())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".perdat")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
