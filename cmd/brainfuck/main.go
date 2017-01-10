package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/umayr/brainfuck"
)

var (
	isBrainFuckRegex    = regexp.MustCompile(`^[<|>|\+|,|-|\.|\[|\]].*$`)
	brainFuckCharsRegex = regexp.MustCompile(`[^<>\+,\-\.\[]]`)
)

func run(program string) {
	result, err := brainfuck.Exec(program)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error executing brainfuck program: %s", err.Error()))
		os.Exit(1)
	}

	fmt.Println(string(result))
}

func main() {

	if isBrainFuckRegex.MatchString(os.Args[1]) {
		run(os.Args[1])
		return
	}

	files := os.Args[1:]
	for _, f := range files {
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println("Error reading file")
			os.Exit(1)
		}

		p := brainFuckCharsRegex.ReplaceAllString(string(buf), "")

		run(p)
	}
}
