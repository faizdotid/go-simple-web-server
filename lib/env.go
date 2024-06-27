// Package: lib
package lib

import (
	"os"
	"strings"
)

// Init: Load environment variables from .env file
// and set them to the system environment variables
// for the application to use.
func init() {
	fileBytes, err := os.ReadFile("lib/.env")
	if err != nil {
		panic(err)
	}
	fileContent := string(fileBytes)

	eachLine := strings.Split(fileContent, "\n")

	for _, line := range eachLine {
		env := strings.Split(line, "=")
		if len(env) == 2 {
			os.Setenv(
				strings.TrimSpace(env[0]),
				strings.TrimSpace(env[1]),
			)
		}
	}
}
