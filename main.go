package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"secret-hunter/modules"
)

func scanContent(content string, source string) {
	for name, pattern := range modules.Providers {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(content, -1)
		for _, m := range matches {
			fmt.Printf("[+] %s FOUND | %s | %s\n", name, source, m)
		}
	}
}

func scanFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err == nil {
		scanContent(string(data), path)
	}
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			scanContent(scanner.Text(), "STDIN")
		}
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: secret-hunter.exe <file|directory>")
		return
	}

	target := os.Args[1]

	filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			scanFile(path)
		}
		return nil
	})
}
