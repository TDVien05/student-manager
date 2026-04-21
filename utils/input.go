package util

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
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
