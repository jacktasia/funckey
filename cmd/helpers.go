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
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func validArgPercent(cmd *cobra.Command, args []string) error {
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
}

func forceInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func extractTextByRegex(text string, regex string) string {
	//matches := []string{}
	r := regexp.MustCompile(regex)

	// if !strings.Contains(regex, "(") && !strings.Contains(regex, ")") {
	// 	log.Warnf("Regex pattern: %s appears to have no capture group!", regex)
	// }

	rawMatches := r.FindAllStringSubmatch(text, -1)
	for _, m := range rawMatches {
		if len(m) > 1 {
			return strings.Trim(m[1], "\n")
			//matches = append(matches, m[1])
		}
	}

	return "" // joinMatches(matches)
}

func runCmdStringOutput(app string, args []string) string {
	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	if err != nil {
		return err.Error()
	}

	return strings.Trim(string(stdout), "\n")
}

func runCmdIntOutput(app string, args []string) int {
	result := runCmdStringOutput(app, args)
	return forceInt(result)
}
