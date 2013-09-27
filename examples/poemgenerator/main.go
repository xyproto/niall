package main

import (
	"fmt"
	"github.com/xyproto/niall"
	"io/ioutil"
	"os"
	"strings"
)

// Learn all the sentences in the given text file. The first and three last lines are ignored.
func learnPoem(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read %s.\n", filename);
		os.Exit(1);
	}
	lines := strings.Split(string(data), "\n");
	for i, line := range lines {
		if i > 1 && i < (len(lines)-3) {
			niall.Learn(line);
		}
	}
}

// Generate a poem with the given number of strophes and lines per strophe.
func generatePoem(strophes int, linesPerStrophe int) {
	// Find a catchy title
	title := strings.Trim(strings.Title(niall.Talk()), ".,;:!?_")
	if strings.Count(title, " ") > 4 {
		words := strings.Split(title, " ")
		title = strings.Join(words[:4], " ")
	}
	// Generate an appropriate amount of dashes
	dashes := ""
	for i := 0; i < len(title); i++ {
		dashes += "-"
	}
	// Output the header
	fmt.Printf("\n%s\n%s\n\n", title, dashes);
	// Generate and output all the srophes
	for reps := 0; reps < strophes; reps++ {
		for i := 0; i < linesPerStrophe; i++ {
			fmt.Println(niall.Talk());
		}
		fmt.Println()
	}
}

func main() {
	niall.Init()
	learnPoem("poem1.txt")
	learnPoem("poem2.txt")
	learnPoem("poem3.txt")
	generatePoem(3, 7)
	niall.Quit()
}
