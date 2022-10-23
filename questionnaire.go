package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type T struct {
	Questions []struct {
		Q string
		A []string
	}
}

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Error! Please provide a file path.")
	}

	var filePath = os.Args[1]
	var data, err = os.ReadFile(filePath)
	check(err)

	t := T{}

	check(yaml.Unmarshal([]byte(data), &t))
	fmt.Printf("\n%v questions found!\n\n", len(t.Questions))
}
