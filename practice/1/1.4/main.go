/**
- 1.4：修改 dup2 程序，输出出现重复行的文件的名称。
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// go run main.go file1 file2 file3
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		log.Fatal("Input filename please")
	}

	for _, filename := range files {
		func(filename string, counts map[string]int) {
			f, err := os.Open(filename)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				return
			}
			countLines(f, filename, counts)
		}(filename, counts)
	}

	for context, n := range counts {
		if n > 1 {
			filename, content := getNameAndContent(context)
			fmt.Printf("filename %s: content %q has duplicate count - %d\n", filename, content, n)
		}
	}
}

func getNameAndContent(context string) (filename string, content string) {
	strs := strings.Split(context, "|")
	return strs[0], strs[1]
}

func countLines(f *os.File, filename string, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[filename+"|"+input.Text()]++
	}
}
