package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var filename = "output_response.txt"

//WriteFile used for write response to text
func WriteFile(msg string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = fmt.Fprintln(f, msg)
	if err != nil {
		return err
	}

	return err
}

//ReadFile used for read file response
func ReadFile() ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//TruncateFile used to truncate file
func TruncateFile() error {
	err := os.Truncate(filename, 0)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
