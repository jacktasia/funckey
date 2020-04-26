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
	"bufio"
	"fmt"
	"strings"
)

func getVolumePercent() int {
	raw := runCmdStringOutput("pactl", []string{"list", "sinks"})

	scanner := bufio.NewScanner(strings.NewReader(raw))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Volume:") {
			rawPercent := extractTextByRegex(raw, "\\s([0-9]+)%")
			return forceInt(rawPercent)
		} else if strings.Contains(line, "Mute:") && strings.Contains(line, "yes") {
			return -1
		}
	}

	fmt.Println(raw)
	return -2 // unkown
}

func setVolumePercent() int {
	return 0
}
