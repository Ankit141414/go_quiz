package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	read := csv.NewReader(file)
	result, err := read.ReadAll()

	buffer := bufio.NewReader(os.Stdin)
	input := make(chan string)

	go func() {
		value, _ := buffer.ReadString('\n')
		input <- value
	}()

	send(result, input)

}

func send(result [][]string, input chan string) {

	for i := range result {
		fmt.Printf("#Problem %d:%s =  \n", i+1, result[i][0])
		select {
		case chvalue := <-input:
			trimmed := strings.TrimSpace(chvalue)
			answer := strings.TrimSpace(result[i][1])

			if trimmed == answer {
				fmt.Println("Correct")
			} else {
				fmt.Println("wrong")
			}

		case <-time.After(4 * time.Second):
			fmt.Println("\ntime limit!")
		}

	}
}
