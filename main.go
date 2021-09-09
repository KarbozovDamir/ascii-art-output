package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileLen = 855
)

func main() {
	if len(os.Args) < 1 { //если нет аргументов выход
		return
	}
	font := "standard"               //базовый шрифт
	switch os.Args[len(os.Args)-1] { //последний элемент в массиве, font
	case "shadow":
		font = "shadow"
	case "thinkertoy":
		font = "thinkertoy"
	}

	//отсекаем font, если есть
	pool := os.Args[1:]                  //скипаем путь
	if os.Args[len(os.Args)-1] == font { // если последний элемент это font, тогда
		pool = pool[:len(pool)-1] //исключаем font из списка строк
	}

	pool, err := makepool(pool)
	if err != nil {
		fmt.Println("Non-valid characters")
		return
	}

	readFile, err := os.Open("fonts/" + font + ".txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	arr := []string{}
	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted")
		return
	}
	printBanners(pool, arr)
}

func makepool(mass []string) (result []string, err error) {
	for _, v := range mass {
		if !isValid(v) {
			return nil, errors.New("Error")
		}
		mass := strings.Split(strings.ReplaceAll(v, "\\n", "\n"), "\n")
		result = append(result, mass...) // mass... - массив mass с N-количеством эл-тов
	}
	return
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
