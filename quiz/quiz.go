package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type testMath struct {
	question string
	answer   string
}

func main() {

	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("time", 30, "the time limit for the quiz (in second)")
	isShuffle := flag.Bool("shuffle", false, "the flag to show whether the question will be shown in random order")

	flag.Parse()

	tests, err := readFile(*fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *isShuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(tests), func(i, j int) {
			tests[i], tests[j] = tests[j], tests[i]
		})
	}

	fmt.Println("press enter to start the quiz!!!")
	bufio.NewScanner(os.Stdout).Scan()

	res := runQuiz(tests, *timeLimit)

	fmt.Printf("You get %d correct answers out of %d\n", res, len(tests))

}

func readFile(fileName string) ([]testMath, error) {

	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("cannot open " + fileName)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("error reading file " + fileName)
	}

	tests := make([]testMath, len(lines))
	for i, line := range lines {
		tests[i] = testMath{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return tests, nil

}

func runQuiz(tests []testMath, timeLimit int) int {

	var res int

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	for i, t := range tests {

		done := make(chan bool)
		go func() {
			var answer string
			fmt.Printf("Problem #%d: %s = ", i+1, t.question)
			fmt.Scanf("%s\n", &answer)
			if answer == t.answer {
				res++
			}
			done <- true
		}()

		select {
		case <-timer.C:
			fmt.Println()
			return res
		case <-done:
		}

		close(done)
	}
	return res
}
