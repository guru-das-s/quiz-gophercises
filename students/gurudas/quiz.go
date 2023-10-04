package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"encoding/csv"
	"flag"
)

func main() {
	var user, ans, correct, total int

	csvFile := flag.String("f", "problems.csv", "A CSV file formatted as question,answer")
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

	for _, fields := range data {
		fmt.Print(fields[0], " = ")
		ans, _ = strconv.Atoi(fields[1])
		fmt.Scan(&user)
		if user == ans {
			correct++
		}
	}

	fmt.Printf("You got %d out of %d questions right!\n", correct, total)
}
