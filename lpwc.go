package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 || (os.Args[1] != "-c" && os.Args[1] != "-l" && os.Args[1] != "-w") {
		fmt.Println("Usage:\tlpwc\t-c <file_name>\n\t\t-l <file_name>\n\t\t-w <file_name>")
		os.Exit(1)
	}

	option := os.Args[1]
	fileName := os.Args[2]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int
	if option == "-l" {
		for scanner.Scan() {
			count++
		}
		fmt.Printf("%10d %s\n", count, fileName)
	} else if option == "-c" {
		fileInfo, err := file.Stat()

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("%10d %s\n", fileInfo.Size(), fileName)
	} else if option == "-w" {
		var wordCount int
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			wordCount += len(words)
		}
		fmt.Printf("%10d %s\n", wordCount, fileName)
	}

}
