package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

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

	filePath := flag.String("file", "", "path of the yaml file with questions")
	numQuestions := flag.Int("n", 10, "# of questions to ask")
	flag.Parse()

	if len(*filePath) == 0 {
		fmt.Println("Usage: questionnaire.go -file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	data, err := os.ReadFile(*filePath)
	check(err)

	t := T{}

	check(yaml.Unmarshal([]byte(data), &t))
	count := len(t.Questions)
	fmt.Printf("\n%v questions found!\n\n", count)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < *numQuestions; i++ {
		fmt.Println(t.Questions[r1.Intn(count)].Q)
	}
}
