package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ParseInt(text string) (int, bool) {
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0, false
	}
	return number, true
}

func MustParseInt(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		panic(1)
	}
	return number
}

func ReadFile(day string) []string {
	fileName := fmt.Sprintf("../../../assets/puzzles/%s.txt", day)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not read file '%s': %v", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading file '%s': %v", fileName, err)
	}

	return input
}

func OpenFile(day string) *os.File {
	fileName := fmt.Sprintf("../../../assets/puzzles/%s.txt", day)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not read file '%s': %v", fileName, err)
	}

	return file
}
