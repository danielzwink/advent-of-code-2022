package main

import (
	"advent-of-code-2022/cmd/puzzles/07/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	root := readFileSystem("07/input")

	calculateSizes := func(currentDirectory *types.Directory) {
		for _, file := range currentDirectory.Files {
			currentDirectory.Size += file.Size
		}
		for _, child := range currentDirectory.Directories {
			currentDirectory.Size += child.Size
		}
	}
	root.Visit(calculateSizes)

	fmt.Printf("Part 1: %v\n", part1(root))
	fmt.Printf("Part 2: %v\n", part2(root))
}

func part1(root *types.Directory) int {
	total := 0

	sumSmallDirectories := func(currentDirectory *types.Directory) {
		if currentDirectory.Size <= 100000 {
			total += currentDirectory.Size
		}
	}

	root.Visit(sumSmallDirectories)
	return total
}

func part2(root *types.Directory) int {
	remainingSpace := 70000000 - root.Size
	missingSpace := 30000000 - remainingSpace

	var smallestDirectory *types.Directory = nil

	findSmallestDirectory := func(currentDirectory *types.Directory) {
		currentDifference := currentDirectory.Size - missingSpace
		if (currentDifference >= 0) && (smallestDirectory == nil || currentDifference < (smallestDirectory.Size-missingSpace)) {
			smallestDirectory = currentDirectory
		}
	}

	root.Visit(findSmallestDirectory)
	return smallestDirectory.Size
}

func readFileSystem(day string) *types.Directory {
	lines := util.ReadFile(day)

	listDirectoryCommand := "$ ls"
	leaveDirectoryCommand := "$ cd .."
	enterDirectoryCommandPattern := regexp.MustCompile("\\$ cd ([a-z]+)")

	directoryPattern := regexp.MustCompile("dir ([a-z]+)")
	filePattern := regexp.MustCompile("([0-9]+) ([a-z\\.]+)")

	root := types.NewDirectory("/", nil)
	current := root

	for _, line := range lines[1:] {
		if line == listDirectoryCommand {
			continue
		} else if strings.HasPrefix(line, "$ cd") {
			if line == leaveDirectoryCommand {
				current = current.Parent
			} else {
				directoryName := enterDirectoryCommandPattern.FindStringSubmatch(line)
				current = current.Directories[directoryName[1]]
			}
		} else {
			if strings.HasPrefix(line, "dir") {
				directoryName := directoryPattern.FindStringSubmatch(line)
				current.AddDirectory(directoryName[1], current)
			} else {
				fileName := filePattern.FindStringSubmatch(line)
				current.AddFile(fileName[2], fileName[1])
			}
		}
	}
	return root
}
