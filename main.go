package main

import (
	"fmt"
	"os"
	"strings"
)

type data struct {
	word   string
	banner string
	flag   string
	isFlag bool
}

var dat data

const (
	fileLen = 855
)

// check amount of arguments
func main() {
	if len(os.Args[1:]) < 4 {
		fmt.Println("Non-valid amount of arguments")
		return
	}
	args := os.Args[1:]
	initValues(args)

	// if !isValid(args[0]) {
	// 	fmt.Println("Non-valid characters")
	// 	return
	// }

	// text := args[0]    // "hello" == [0]
	// font := "standard" //base font
	// if len(args) == 3 {
	// 	switch args[1] {
	// 	case "shadow":
	// 		font = "shadow"
	// 	case "thinkertoy":
	// 		font = "thinkertoy"
	// 	case "standard":
	// 		font = "standard"
	// 	default:
	// 		fmt.Println("Non-valid font")
	// 		return
	// 	}
	// }

	// // Read the content of the file
	// argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	// arr := []string{}
	// readFile, err := os.Open("fonts/" + font + ".txt")
	// defer readFile.Close()

	// if err != nil {
	// 	log.Fatalf("failed to open file: %s", err)
	// }

	// fileScanner := bufio.NewScanner(readFile)
	// fileScanner.Split(bufio.ScanLines)

	// for fileScanner.Scan() {
	// 	arr = append(arr, fileScanner.Text())
	// }

	// if len(arr) != fileLen {
	// 	fmt.Println("File is corrupted")
	// 	return
	// }
	// larg := len(argsArr)

	// if larg >= 2 {
	// 	if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
	// 		argsArr = argsArr[:larg-1]
	// 	}
	// }
	// ifoutput := false
	// if len(args) == 3 {
	// 	ifoutput = true
	// }
	// if !ifoutput {
	// 	printBanners(argsArr, arr)
	// 	return
	// }
	// nf := args[2][9:]
	// if len(args[2][:9]) != 9 {
	// 	fmt.Println("no valid len")
	// 	os.Exit(0)
	// }

	// if !strings.HasSuffix(nf, ".txt") {
	// 	fmt.Println("ERROR: must be txt file")
	// 	os.Exit(0)
	// }
	// if strings.HasSuffix(nf, "standard.txt") || strings.HasSuffix(nf, "shadow.txt") || strings.HasSuffix(nf, "thinkertoy.txt") {
	// 	fmt.Println("ERROR")
	// 	os.Exit(0)
	// }
	// // Creating a new file
	// f, e := os.Create(nf)
	// if e != nil {
	// 	fmt.Println("Please, write the flag \"--output=\", followed by name of the file to create")
	// 	os.Exit(0)
	// }
	// if nf == "--output=" {
	// 	fmt.Println("No filename")
	// }
	// output := output(argsArr, arr)
	// f.WriteString(output)
	// defer f.Close()
}

func initValues(s []string) {
	if !isValid(s[0]) {
		fmt.Println("Not Valid character")
		return
	}
	dat.word = s[0]
	for i := 1; i < len(s); i++ {
		if !isValid(s[i]) {
			fmt.Println("Not Valid character")
			return
		}
		if strings.HasPrefix(s[i], "--output") {
			if i < len(s)-1 { // hello --output   shadow
				fmt.Println("Error order") // hello shadow --output
				return
			}
			if s[i] == "--output" || s[i] == "--output=" { //--output
				fmt.Println("output: needs argument")
				return
			}
			temp := s[i]        // --output=asdsadad
			if temp[8] != '=' { //=
				fmt.Println("wrong operator")
				return
			}
			dat.flag = temp[9:]
			dat.isFlag = true
			continue
		}
		dat.banner = s[i]
	}
}

//Write string
func output(banners, arr []string) string {
	res := ""
	for _, ch := range banners {
		if ch == "" {
			res += ""
			continue
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				res += (arr[int(n)+i])
			}
			res += "\n"
		}
		res += "\n"
	}
	return res
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

// Print the full outcome
func printBanners() {
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
