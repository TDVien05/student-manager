package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetString(welcome string) string {
	fmt.Print(welcome)

	input, _ := reader.ReadString('\n')
	// TrimSpace removes the newline character (\n or \r\n)
	return strings.TrimSpace(input)
}

func GetInt(welcome string) int {
	for {
		stringValue := GetString(welcome)

		value, err := strconv.Atoi(stringValue)

		if err != nil {
			fmt.Println("Error: That's not a valid number. Try again.")
			continue
		}

		return value
	}
}

func GetUInt(welcome string) uint {
	for {
		stringValue := GetString(welcome)

		value, err := strconv.ParseUint(stringValue, 10, 0)

		if err != nil {
			fmt.Println("Error: Please enter a valid positive number.")
			continue
		}

		// Convert the result to uint before returning
		return uint(value)
	}
}

func GetFloat(welcome string) float64 {
	for {
		stringValue := GetString(welcome)

		value, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			fmt.Println("Error: That's not a valid number. Try again.")
			continue
		}

		return value
	}
}

func UpdateString(welcome string, oldValue string) string {
	// Step 1: Print welcome string
	fmt.Print(welcome)

	// Step 2: Read input
	newValue, _ := reader.ReadString('\n')
	newValue = strings.TrimSpace(newValue)

	// Step 3: Check empty input
	if newValue == "" {
		return oldValue
	}

	// Step 4: Replace oldValue with newValue
	return strings.ReplaceAll(oldValue, oldValue, newValue)
}

func UpdateInt(welcome string, oldValue int) int {
	// Step 1: Print prompt
	fmt.Print(welcome)

	// Step 2: Read input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Step 3: If empty → keep old value
	if input == "" {
		return oldValue
	}

	// Step 4: Convert string → int
	newValue, err := strconv.Atoi(input)
	if err != nil {
		// If invalid number, keep old value
		return oldValue
	}

	return newValue
}
