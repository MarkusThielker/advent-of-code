package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ReadInput(path string) (lines []string) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
		return
	}

	file, err := os.Open(filepath.Dir(filename) + "/" + path + ".txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}
