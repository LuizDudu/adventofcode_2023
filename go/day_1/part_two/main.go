package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumber(word string) (int, error) {
	if strings.Contains(word, "one") {
		return 1, nil
	}

	if strings.Contains(word, "two") {
		return 2, nil
	}

	if strings.Contains(word, "three") {
		return 3, nil
	}

	if strings.Contains(word, "four") {
		return 4, nil
	}

	if strings.Contains(word, "five") {
		return 5, nil
	}

	if strings.Contains(word, "six") {
		return 6, nil
	}

	if strings.Contains(word, "seven") {
		return 7, nil
	}

	if strings.Contains(word, "eight") {
		return 8, nil
	}

	if strings.Contains(word, "nine") {
		return 9, nil
	}

	return 0, errors.New("doesn't contain a number")
}

func main() {
	filePath := filepath.Join("..", "real.input")
	file, err := os.Open(filePath)
	check(err)

	var result = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var buffer bytes.Buffer
		word := ""

		for _, char := range scanner.Text() {
			character := string(char)
			word += character
			_, err := strconv.Atoi(character)

			number, getNumberErr := getNumber(word)
			if getNumberErr == nil {
				fmt.Println(word)
				character = strconv.Itoa(number)
			}

			if err != nil && getNumberErr != nil {
				continue
			}

			word = word[len(word)-1:]
			if buffer.Len() == 0 {
				buffer.WriteString(character)
				continue
			}

			if buffer.Len() == 2 {
				buffer.Truncate(1)
			}

			buffer.WriteString(character)
		}

		if buffer.Len() == 1 {
			buffer.WriteString(buffer.String())
		}

		number, err := strconv.Atoi(buffer.String())
		check(err)

		result += number
	}

	fmt.Println(result)
}
