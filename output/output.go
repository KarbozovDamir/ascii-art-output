package output

import "fmt"

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
