package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	Files = []string{
		"white",
		"blue",
		"black",
		"red",
		"green",
		"multicolor",
		"colorless",
		"land",
	}
)

func main() {
	arg := os.Args[1]
	if len(os.Args) < 2 || arg == "" {
		fmt.Println("No arg")
		return
	} else if arg == "all" {
		total := 0
		for _, filename := range Files {
			total += PrintFileStats(filename)
		}
		fmt.Printf("Total :%v\n", total)
	} else {
		found := false
		for _, filename := range Files {
			if arg == filename {
				PrintFileStats(filename)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Unknown arg '%v'", arg)
		}
	}
}

func PrintFileStats(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			header := strings.TrimSpace(line[1:])
			fmt.Println(header)
			continue
		}

	}
	return 0
}
