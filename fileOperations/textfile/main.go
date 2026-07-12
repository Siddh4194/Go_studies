package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WriteFile(txt string) {
	content := txt

	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Print(err)
	}
}

func ReadFile() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Default().Fatal(err)
	}
	fmt.Print(string(data))
}

func ReadFileLineByLine() {
	file, err := os.OpenFile("data.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Default().Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		fmt.Printf("line: %d : %s\n",lineNumber,scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Print("Error reading lines",err)
	}
}

func main() {
	fmt.Println("the arguments calculator")

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments. Use --help for usage details.")
		return
	}

	if os.Args[1] == "--help" {
		fmt.Println("========================")
		fmt.Println("Usage: go run . [option] [text]")
		fmt.Println("Options:")
		fmt.Println("  --help     Show this help message")
		fmt.Println("  --write    Write text to data.txt")
		fmt.Println("  --read     Read text from data.txt")
		fmt.Println("  --read-lines Read text from data.txt line by line")
		return
	}

	if os.Args[1] == "--write" {
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments for writing text")
			return
		}
		WriteFile(os.Args[2])
	}

	if os.Args[1] == "--read" {
		ReadFile()
	}

	
	if os.Args[1] == "--read-lines" {
		ReadFileLineByLine()
	}
}
