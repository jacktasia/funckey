/*
Copyright © 2020 Jack Angers <jacktasia@gmail.com>

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
	"github.com/spf13/cobra"
	"strconv"
)

var setVolumePercentCmd = &cobra.Command{
	Use:   "set-percent",
	Short: "set the current system volume by percent",
	Args:  validArgPercent,
	Run: func(cmd *cobra.Command, args []string) {

		newPercent, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		b := NewVolume()
		b.setPercent(newPercent)
	},
}

func init() {
	volumeCmd.AddCommand(setVolumePercentCmd)
}
