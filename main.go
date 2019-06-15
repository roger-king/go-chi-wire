package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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
		//fmt.Println(string(read))
		fmt.Println(path)

		newContents := strings.Replace(string(read), "roger-king/go-chi-wire", "roger-king/new-repo", -1)

		fmt.Println(newContents)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func main() {
	// First edit go.mod
	err := filepath.Walk(".", visit)

	if err != nil {
		panic(err)
	}

	err = filepath.Walk("pkg", visit)
	if err != nil {
		panic(err)
	}
}
