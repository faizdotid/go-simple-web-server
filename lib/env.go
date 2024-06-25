package lib

import (
	"os"
	"strings"
)

func LoadEnv() {
	fileBytes, err := os.ReadFile("lib/.env")
	if err != nil {
		panic(err)
	}
	fileContent := string(fileBytes)

	eachLine := strings.Split(fileContent, "\n")

	for _, line := range eachLine {
		env := strings.Split(line, "=")
		if len(env) == 2 {
			os.Setenv(env[0], env[1])
		}
	}
}
