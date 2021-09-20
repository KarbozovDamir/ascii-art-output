unc main() {
	if len(os.Args[1:]) < 2 || len(os.Args[1:]) > 3 {
		fmt.Println("Non-valid amount of arguments")
		return
	}
	args := os.Args[1:]
	if !isValid(args[0]) {
		fmt.Println("Non-valid characters")
		return
	}

	text := args[0]    // "hello" == [0]
	font := "standard" //base font
	if len(args) == 3 {
		switch args[1] {
		case "shadow":
			font = "shadow"
		case "thinkertoy":
			font = "thinkertoy"
		case "standard":
			font = "standard"
		default:
			fmt.Println("Non-valid font")
			return
		}
	}

	// Read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

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
	ifoutput := false
	if len(args) == 3 {
		ifoutput = true
	}
	if !ifoutput {
		printBanners(argsArr, arr)
		return
	}
	nf := args[2][9:]
	if len(args[2][:9]) != 9 {
		fmt.Println("no valid len")
		os.Exit(0)
	}

	if !strings.HasSuffix(nf, ".txt") {
		fmt.Println("ERROR: must be txt file")
		os.Exit(0)
	}
	if strings.HasSuffix(nf, "standard.txt") || strings.HasSuffix(nf, "shadow.txt") || strings.HasSuffix(nf, "thinkertoy.txt") {
		fmt.Println("ERROR")
		os.Exit(0)
	}
	// Creating a new file
	f, e := os.Create(nf)
	if e != nil {
		fmt.Println("Please, write the flag \"--output=\", followed by name of the file to create")
		os.Exit(0)
	}
	if nf == "--output=" {
		fmt.Println("No filename")
	}
	output := output(argsArr, arr)
	f.WriteString(output)
	defer f.Close()
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
