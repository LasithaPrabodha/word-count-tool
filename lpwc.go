package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: lpwc [-c | -l | -w | -m] <file_name>")
		os.Exit(1)
	}

	options := []string{"-c", "-l", "-w", "-m"}
	arg1 := os.Args[1]
	if len(os.Args) == 3 {
		options = []string{arg1}
	}

	var f *os.File
	var err error
	var fileName string

	if len(os.Args) == 3 {
		fileName, f = openFile()
		defer f.Close()

	} else if len(os.Args) == 2 {
		if slices.Contains(options, arg1) {
			f = os.Stdin
		} else {
			fileName, f = openFile()
			defer f.Close()
		}
	}

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	scanner := bufio.NewScanner(f)

	counts["-c"] = int(fi.Size())
	for scanner.Scan() {
		line := scanner.Text()
		counts["-l"]++
		counts["-w"] += len(strings.Fields(line))
		counts["-m"] += len([]rune(line))
	}

	for _, option := range options {
		fmt.Printf("%10d ", counts[option])
	}
	fmt.Println("\t" + fileName)

}

func openFile() (string, *os.File) {
	fileName := os.Args[len(os.Args)-1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return fileName, f
}
