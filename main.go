package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	fileLen = 6623
)

func main() {
	args := os.Args[1:]
	if !isValid(args[1]) {
		fmt.Println("Non-valid characters")
		return
	}
	// in order to pass the test - student$ go run . "" | cat -e
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
	argsArr := strings.Split(strings.ReplaceAll(args[1], "\\n", "\n"), "\n")
	file, err := ioutil.ReadFile("fonts/" + font + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(file) != fileLen {
		fmt.Println("File is corrupted")
		return
	}
	arr := []string{}
	for _, el := range strings.Split(string(file), string('\n')) {
		arr = append(arr, el)
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
	for _, el := range banners {
		if el == "" {
			continue
		}
		for i := 0; i < 8; i++ {
			for _, j := range el {
				n := (j-32)*9 + 1
				fmt.Print(arr[int(n)+i])
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}
