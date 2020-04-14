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
	"github.com/spf13/cobra"
	"strconv"
)

// setPercentCmd represents the setPercent command
var setPercentCmd = &cobra.Command{
	Use:   "set-percent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a percent argument")
		}

		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Errorf("invalid percent specified: %s", args[0])
		}

		if i < 5 || i > 100 {
			return fmt.Errorf("invalid percent specified: %d", i)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		newPercent, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		b := NewBrightness()
		b.setPercent(newPercent)
	},
}

func init() {
	// rootCmd.AddCommand(setPercentCmd)
	brightnessCmd.AddCommand(setPercentCmd)
}
