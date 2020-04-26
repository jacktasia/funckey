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

func NewVolume() *Volume {
	br := &Volume{}

	return br
}

type Volume struct {
}

func (b *Volume) up() {
	currentPercent := getVolumePercent()

	levels := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110}
	for _, level := range levels {
		if level > currentPercent {
			b.setPercent(level)
			break
		}
	}
}

func (b *Volume) toggleMute() {
	c := getVolumePercent()

	if c == -1 {
		unmuteVolume()
	} else {
		muteVolume()
	}
}

func (b *Volume) down() {
	currentPercent := getVolumePercent()

	if currentPercent == -1 { //muted
		unmuteVolume()
		b.down()
		return
	}

	levels := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0}
	for _, level := range levels {
		if level < currentPercent {
			b.setPercent(level)
			break
		}
	}
}

func (b *Volume) getPercent() int {
	return getVolumePercent()
}

func (b *Volume) setPercent(newPercent int) {
	if newPercent > 100 {
		newPercent = 100
	} else if newPercent < 0 {
		newPercent = 0
	}

	curPercent := getVolumePercent()
	if curPercent == -1 {
		unmuteVolume()
	} else if newPercent == 0 {
		muteVolume()
	} else {
		setVolumePercent(newPercent)
	}
}

///
func muteVolume() {
	runCmdStringOutput("pactl", []string{"set-sink-mute", "0", "1"})
}

func unmuteVolume() {
	runCmdStringOutput("pactl", []string{"set-sink-mute", "0", "0"})
}

func setVolumePercent(per int) {
	runCmdStringOutput("pactl", []string{"set-sink-volume", "0", fmt.Sprintf("%d%%", per)})
}

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
