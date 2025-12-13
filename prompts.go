package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func PromptRequired(prompt string) string {
	for {
		fmt.Printf("%s: ", prompt)
		input, err := reader.ReadString('\n')

		if err != nil {
			// TODO: Complete this
			fmt.Println("Error: read")
			os.Exit(1) // Check this error code
		}

		input = strings.TrimSpace(input)

		if input != "" {
			return input
		}

		fmt.Println("This value is required. Enter a value.")
	}
}

func PromptOptional(prompt string) string {
	fmt.Printf("%s (Optional): ", prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: Unable to read input, check and try again: '%s'", input)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

func PromptYesNo(prompt string) bool {
	fmt.Printf("%s [y/N]: ", prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: Unable to read input, check and try again: '%s'", input)
		os.Exit(1)
	}

	value := strings.ToLower(strings.TrimSpace(input))

	return value == "y"
}

func PromptWithDefault(prompt string, defaultValue string) string {
	fmt.Printf("%s (Default: %s): ", prompt, defaultValue)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: Unable to read input, check and try again: '%s'", input)
		os.Exit(1)
	}

	if input == "" {
		return defaultValue
	}

	return strings.TrimSpace(input)
}

func CloseReader() {}
