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
	"fmt"
	"os/exec"
	"strconv"
)

func NewBrightness() *Brightness {
	br := &Brightness{
		MaxVal:     getMaxScreenBrightness(),
		CurrentVal: getCurrentScreenBrightness(),
	}

	return br
}

type Brightness struct {
	MaxVal     int
	CurrentVal int
}

func (b *Brightness) up() {
	currentPercent := b.getPercent()

	levels := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, level := range levels {
		if level > currentPercent {
			b.setPercent(level)
			break
		}
	}
}

func (b *Brightness) down() {
	currentPercent := b.getPercent()

	levels := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}
	for _, level := range levels {
		if level < currentPercent {
			b.setPercent(level)
			break
		}
	}
}

func (b *Brightness) getPercent() int {
	return int((float64(b.CurrentVal) / float64(b.MaxVal)) * 100)
}

func (b *Brightness) setPercent(newPercent int) {
	if newPercent < 5 {
		newPercent = 5
	} else if newPercent > 100 {
		newPercent = 100
	}

	newVal := int(float64(newPercent) / float64(99.9) * float64(b.MaxVal))
	setScreenBrightness(newVal)
}

func getCurrentScreenBrightness() int {
	return runCmdIntOutput("cat", []string{"/sys/class/backlight/intel_backlight/actual_brightness"})
}

func getMaxScreenBrightness() int {
	return runCmdIntOutput("cat", []string{"/sys/class/backlight/intel_backlight/max_brightness"})
}

func setScreenBrightness(val int) {
	// `sudo` before `tee`
	args := []string{"-c", fmt.Sprintf("echo %s |  tee /sys/class/backlight/intel_backlight/brightness", strconv.Itoa(val))}
	cmd := exec.Command("/bin/bash", args...)
	cmd.Output()
}
