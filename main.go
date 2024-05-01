package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"main/Structures"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	//! STEP 1 : Get file name using CSV Package
	csvFileName := flag.String("csv", "./utils/a.csv", "A csv file in the format 'Questions, Answers'")
	timeLimit := flag.Int("limit", 30, "Time Limit for the Quiz in seconds")

	flag.Parse()
	//! STEP 2 :  Open File using OS Package
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(
			fmt.Sprintf("Failed to open CSV File : %s\n", *csvFileName))
	}
	//! STEP 3 :  Make a CSV Reader using Encoding/csv Package
	reader := csv.NewReader(file)
	//! STEP 4 :  Parse CSV : Read or get all lines
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV File")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//! STEP 5 :  Use variable however you want
	rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	questionaires(problems, timer)
}

// ! Returns an Array of Problem type with length same as length of lines variable
func parseLines(lines [][]string) []Structures.Problem {
	ret := make([]Structures.Problem, len(lines))
	for i, line := range lines {
		ret[i] = Structures.Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// ! Asks Questions
func questionaires(problems []Structures.Problem, timer *time.Timer) {
	correct := 0
problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d : %s =\t", i+1, problem.Question)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop

		case ans := <-answerCh:

			if ans == problem.Answer {
				correct += 1
			}
		}

	}
	fmt.Printf("Correct    :\t%d\n", correct)
	fmt.Printf("Incorrect  :\t%d\n", len(problems)-correct)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

/*
Get file name using flag package
Open File using OS package
Make a CSV Reader
Parse the CSV File to get all Lines and convert these lines to a structure using an array variable
Use this array as you want
*/
