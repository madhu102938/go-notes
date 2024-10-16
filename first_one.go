// reading directory, writing to file, handling errors
// safeWriter (best practice)

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// type celsius int <- declaring new type
// type celsius = int <- Type Alias

// safe way to write to file with minimal code :)
type SafeWriter struct {
	err error
	w io.Writer
}

func (sw *SafeWriter) WriteLn(text string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.w, text)
}

func safeWritingToFile(path string, words []string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("File could not be opened")
		return err
	}

	defer file.Close()

	sw := SafeWriter{w : file}
	for _, word := range words {
		sw.WriteLn(word)
	}

	return sw.err
}

func writingToFile(path string, words []string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("File could not be opened")
		return err
	}
	
	defer file.Close()
	// Go ensures that all deferred actions take place before the containing function returns
	// Even if some error occurs in future, file will closed, only then will be return
	
	for i, word := range words {
		_, err := fmt.Fprintln(file, word)
		if err != nil {
			fmt.Println("Failed after ", i, "lines")
			return err
		}
	}

	return err
}

func main() {
	files, err := os.ReadDir("./") // error type specially for errors :)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range files {
		fmt.Printf("%-20v %5v\n", info.Name(), info.IsDir())
	}

	proverbs := make([]string, 0, 13)
	scanner := bufio.NewScanner(os.Stdin)
	n := 13
	for i := 0; i < n; i++ {
		scanner.Scan()
		if scanner.Err() == nil {
			proverbs = append(proverbs, strings.Trim(scanner.Text(), " "))
		}
	}
	
	for _, proverb := range proverbs {
		fmt.Println(proverb)
	}

	err = writingToFile("./proverbs.txt", proverbs)
	if err == nil {
		fmt.Println("Writing to file successful 1 : )")
	}

	err = safeWritingToFile("./proverbs(Safe).txt", proverbs)
	if err == nil {
		fmt.Println("Writing to file successful 2 : )")
	}
}