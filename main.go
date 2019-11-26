package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("argument error")
		os.Exit(1)
	}
	p := os.Args[1]

	// check path exists
	_, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s not exists.\n", p)
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	// convert to abs path
	absPath, err := filepath.Abs(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print
	fmt.Println(absPath)

	os.Exit(0)
}
