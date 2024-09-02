package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := filepath.Join("real.input")
	file, err := os.Open(filePath)
	check(err)

	var result = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var buffer bytes.Buffer

		for _, char := range scanner.Text() {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			if buffer.Len() == 0 {
				buffer.WriteString(string(char))
				continue
			}

			if buffer.Len() == 2 {
				buffer.Truncate(1)
			}

			buffer.WriteString(string(char))
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
