package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Dumb ass quiz. File content below ...")
	fmt.Println("------")
	file, err := os.Open("problems.csv")
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)
	questions, err := reader.ReadAll()
	if err != io.EOF && err != nil {
		panic(err)
	}

	fmt.Println(questions)
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	for _, value := range questions {
		fmt.Printf("What is", value[0])
	}
}
