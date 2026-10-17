package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
		return
	}

	filePath := filepath.Join(homeDir, ".iwantto")
	content, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer content.Close()

	myList := make([]string, 0)
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		myList = append(myList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("You want to:")
		for i, str := range myList {
			fmt.Printf("%d: %s\n", i+1, str)
		}
		return
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("You want to:")
		for i, str := range myList {
			fmt.Printf("%d: %s\n", i+1, str)
		}
		return
	}

	switch args[1] {
	case "h", "help", "-h", "--help":
		fmt.Println("Usage: iwantto [l/ist|a/add|r/remove] [String|removal number]")
		return
	case "l", "list":
		break
	case "r", "remove":
		if len(args) < 3 {
			fmt.Println("Invalid argument: need item number to remove")
			return
		}
		if args[2] == "all" {
			myList = myList[0:0]
			break
		}
		n, err := strconv.Atoi(args[2])
		n -= 1
		if err != nil || n < 0 || n >= len(myList) {
			fmt.Println("Invalid argument: please provide valid item number")
			return
		}
		myList = append(myList[:n], myList[n+1:]...)
	case "add", "a":
		if len(args) < 3 {
			fmt.Println("Invalid argument: need string to add")
			return
		}
		myList = append(myList, strings.Join(args[2:], " "))
	default:
		if len(args) < 2 {
			fmt.Println("Unknown command. Use -h or --help for usage information")
			return
		} else {
			myList = append(myList, strings.Join(args[1:], " "))
		}
	}
	readList(myList)

	if err := content.Truncate(0); err != nil {
		fmt.Fprintf(os.Stderr, "Error truncating file: %v\n", err)
		return
	}
	if _, err := content.Seek(0, 0); err != nil {
		fmt.Fprintf(os.Stderr, "Error seeking file: %v\n", err)
		return
	}

	writer := bufio.NewWriter(content)
	for _, line := range myList {
		if _, err := fmt.Fprintln(writer, line); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			return
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing writer: %v\n", err)
		return
	}
}
func readList(myList []string) {
	fmt.Println("You want to:")
	for i, str := range myList {
		fmt.Printf("%d: %s\n", i+1, str)
	}
}
