package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type data struct {
	rawInput string
	banner   string
	flag     string
	isFlag   bool
}

var d data

func initValues(s []string) {
	d.banner = "standard"
	if !isValid(s[0]) {
		fmt.Println("Not Valid character")
		os.Exit(1)
	}
	d.rawInput = s[0]
	for i := 1; i < len(s); i++ {
		if !isValid(s[i]) {
			fmt.Println("Not Valid character")
			os.Exit(1)
		}
		if strings.HasPrefix(s[i], "--output") {
			if i < len(s)-1 { // hello --output   shadow
				fmt.Println("Error order") // hello shadow --output
				os.Exit(1)
			}
			if s[i] == "--output" || s[i] == "--output=" { //--output
				fmt.Println("output: needs argument")
				os.Exit(1)
			}
			temp := s[i]        // --output=asdsadad
			if temp[8] != '=' { //=
				fmt.Println("wrong operator")
				os.Exit(1)
			}
			d.flag = temp[9:]
			d.isFlag = true
			continue
		}
		d.banner = s[i]
	}
}

// check amount of arguments
func main() {
	if len(os.Args[1:]) > 3 {
		fmt.Println("Non-valid amount of arguments")
		return
	}
	args := os.Args[1:]

	initValues(args) //initiliaze data -> split arguments : word, banner, flag and adding in data struct

	splittedWord := strings.Split(d.rawInput, "\\n")

	file, err := os.Open("fonts/" + d.banner + ".txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file) //Adding new function in file

	arrData := make([]string, 0) // value is 0, but we will expand from 0 to 855 index

	for scanner.Scan() { //it will go until reach to the end of file
		arrData = append(arrData, scanner.Text()) // scan by line
	}

	if len(arrData) != 855 {
		fmt.Println("file is corrupted")
		return
	}

	temp := ""
	if !d.isFlag { // if there is no flag, then print words
		for _, word := range splittedWord {
			if word == "" {
				fmt.Println()
			} else {
				for i := 1; i < 9; i++ {
					for _, char := range word {
						temp += arrData[(int(char)-32)*9+i]
					}
					if i != 8 {
						temp += "\n"
					}
				}
				fmt.Println(temp)
				temp = ""
			}
		}
	} else { // else write file
		file, err := os.Create(d.flag + ".txt")
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, word := range splittedWord {
			if word == "" {
				file.Write([]byte("\n"))
			} else {
				for i := 1; i < 9; i++ {
					for _, char := range word {
						temp += arrData[(int(char)-32)*9+i]
					}
					temp += "\n"
				}
				file.Write([]byte(temp))
				temp = ""
			}
		}
	}

}

// Check for valid of characters by runes from 32 to 126
func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}
