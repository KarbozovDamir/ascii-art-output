package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		file, err := os.Open("fonts/" + "standard" + ".txt")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		scanner := bufio.NewScanner(file) //Adding new function in file
		b.StartTimer()
		fill(scanner)

		b.StopTimer()
		file.Close()
		b.StartTimer()
	}
}
