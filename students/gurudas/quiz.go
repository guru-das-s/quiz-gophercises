package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"encoding/csv"
	"flag"
	"time"
)

func main() {
	var ans, correct, total int

	csvFile := flag.String("f", "problems.csv", "A CSV file formatted as question,answer")
	timerDuration := flag.Int("t", 30, "Timeout (in seconds) for duration of quiz")
	flag.Parse()

	f, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	total = len(data)

	timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)

	for _, fields := range data {
		fmt.Print(fields[0], " = ")
		userInputChan := make(chan int)
		go func() {
			var user int
			fmt.Scan(&user)
			userInputChan <- user
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d questions right!\n", correct, total)
			return
		case user := <-userInputChan:
			ans, _ = strconv.Atoi(fields[1])
			if user == ans {
				correct++
			}
		}
	}
}
