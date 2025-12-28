package main

import (
	"bufio"
	"encoding/csv"
	"flag"
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
		for {
			value, _ := buffer.ReadString('\n')
			input <- value
		}
	}()

	flags := flag.Int("time", 30, "timer")
	flag.Parse()

	send(result, input, *flags)

}

func send(result [][]string, input chan string, flags int) {

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

		case <-time.After(time.Duration(flags) * time.Second):
			fmt.Println("\ntime limit!")
		}

	}
}
