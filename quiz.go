package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("questions.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)

	var answer string
	var correct, wrong int
	answerCh := make(chan string)
	doneCh := make(chan struct{})
	questionCh := make(chan []string)
	fmt.Println("Press enter to start!")
	fmt.Scanln(&answer)
	ticker := time.NewTicker(time.Second * 10)
	defer func() {
		ticker.Stop()
		fmt.Println("Correct: ", correct, "Wrong :", wrong)
	}()

	go func() {
		select {
		case _ = <-ticker.C:
			fmt.Println("Time out")
			close(doneCh)
			return
		}
	}()

	go func() {
		for _, q := range records {
			questionCh <- q
		}
	}()

	go func() {
		for {
			q := <-questionCh
			fmt.Println(q[0], "=?")
			ans := <-answerCh
			a, _ := strconv.Atoi(strings.Trim(q[1], " "))
			b, _ := strconv.Atoi(strings.Trim(ans, " "))
			if b != a {
				wrong++
			} else {
				correct++
			}
		}
	}()

	go func() {
		var ans string
		fmt.Scanln(&ans)
		answerCh <- ans
	}()


