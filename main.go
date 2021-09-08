package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileLen = 855
)

func main() {
	args := os.Args[1:]
	if !isValid(args[1]) {
		fmt.Println("Non-valid characters")
		return
	}
	//in order to pass the test - student$ go run . "" | cat -e
	if args[1] == "" {
		fmt.Println()
		return
	}
	// in order to pass the test - student$ go run . "Hello\n" | cat -e
	if args[1] == "\\n" {
		fmt.Println("\n")
		return
	}
	text := args[0]     // "hello" == [0]
	font := ""          // shadow || thinkertoy
	if len(args) == 2 { // mean [0] == "hello" [1] == shadow or thinkertoy
		font = args[1]
		// fmt.Println(text)
	} else if len(args) == 1 { // without font(shadow and thin...)
		font = "standard"
	} else {
		fmt.Println("ERROR") // when len(args) != 2 & !=1
		return
	}

	//read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")

	arr := []string{}

	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	// var fileTextLines []string

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted")
		return
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	printBanners(argsArr, arr)
}

// check for valid of characters by runes from 32 to 126
func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}

//print the full outcome
func printBanners(banners, arr []string) {
	for _, ch := range banners {
		if ch == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Print(arr[int(n)+i])
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}
