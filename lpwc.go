package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 || (os.Args[1] != "-c" && os.Args[1] != "-l") {
		fmt.Println("Usage:\tlpwc\t-c <file_name>\n\t\t-l <file_name>")
		os.Exit(1)
	}

	option := os.Args[1]
	fileName := os.Args[2]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	var count int
	if option == "-l" {
		scanner := bufio.NewScanner(file)

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
	}

}
