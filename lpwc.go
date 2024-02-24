package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: lpwc -c <file_name>")
		os.Exit(1)
	}

	fileName := os.Args[2]
	fileInfo, err := os.Stat(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("%10d %s\n", fileInfo.Size(), fileName)

}
