package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	var user, ans, correct, total int;

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		total++

		split := strings.Split(line, ",")
		ans, _ = strconv.Atoi(split[1])
		fmt.Print(split[0], " = ")

		fmt.Scan(&user)

		if user == ans {
			correct++;
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You got %d out of %d questions right!\n", correct, total)
}
