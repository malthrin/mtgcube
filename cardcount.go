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
		"multi",
		"nocolor",
		"land",
	}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arg")
		return
	}
	arg := os.Args[1]
	if arg == "all" {
		total := 0
		for _, filename := range Files {
			cf, err := ReadFile(filename)
			if err != nil {
				fmt.Printf("Read error: %v\n", err.Error())
				continue
			}
			fmt.Printf("%v\t%v\n", cf.Name, cf.CardCount())
		}
		fmt.Printf("\nTotal\t%v\n", total)
	} else {
		found := false
		for _, filename := range Files {
			if arg == filename {
				found = true
				cf, err := ReadFile(filename)
				if err != nil {
					fmt.Printf("Read error: %v\n", err.Error())
				} else {
					PrintFile(cf)
				}
				break
			}
		}
		if !found {
			fmt.Printf("Unknown arg '%v'\n", arg)
		}
	}
}

func ReadFile(filename string) (*CardFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return &CardFile{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sections := make([]*Section, 0)
	var current *Section
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			header := strings.TrimSpace(line[1:])
			current = &Section{Header: header, Lines: make([]string, 0)}
			sections = append(sections, current)
			continue
		}
		current.Lines = append(current.Lines, line)
	}
	return &CardFile{Name: filename, Sections: sections}, nil
}

func PrintFile(file *CardFile) {
	fmt.Printf("%v\t%v\n", file.Name, file.CardCount())
	for _, section := range file.Sections {
		l := len(section.Lines)
		if l > 0 {
			fmt.Printf("\t%v\t%v\n", section.Header, l)
		}
	}
}

type CardFile struct {
	Name     string
	Sections []*Section
}

type Section struct {
	Header string
	Lines  []string
}

func (file *CardFile) CardCount() int {
	total := 0
	for _, s := range file.Sections {
		total += len(s.Lines)
	}
	return total
}
