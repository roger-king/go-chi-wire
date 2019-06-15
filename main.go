package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var replace string

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("*.go", fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		fmt.Println(path)

		newContents := strings.Replace(string(read), "roger-king/go-chi-wire", replace, -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func main() {
	replace = os.Args[1]
	err := filepath.Walk("pkg", visit)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk("cmd", visit)
	if err != nil {
		panic(err)
	}
}
