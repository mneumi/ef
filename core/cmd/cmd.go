package cmd

import (
	"bytes"
	"os/exec"
	"strings"
)

func RunCMD(name string, args ...string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	result := out.String()

	return trimResult(result), nil
}

func trimResult(result string) string {
	return strings.TrimFunc(result, func(r rune) bool {
		str := string(r)
		if str == " " || str == "\n" {
			return true
		}
		return false
	})
}
