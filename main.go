package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	read := csv.NewReader(file)

	result, err := read.ReadAll()

	for i, _ := range result {

		fmt.Printf("#Problem %d:%s =  \n", i+1, result[i][0])

		buffer := bufio.NewReader(os.Stdin)
		value, err := buffer.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		trimmed := strings.TrimSpace(value)
		answer := strings.TrimSpace(result[i][1])

		if trimmed == answer {
			fmt.Println("Correct")
		} else {
			fmt.Println("wrong")
		}

	}

}
