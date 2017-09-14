package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file2.txt.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(gz)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		fmt.Print(line)
	}
}
