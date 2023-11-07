package main

import (
	"flag"
	"fmt"
	"time"
)

func TestGoroutineFunc() {
	timeLimit := flag.Int("timer", 3, "Time Limit for the Quiz in seconds")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
loop:
	for i := 0; i < 5; i++ {
		answerChannel := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChannel <- ans
		}()
		select {
		case value := <-answerChannel:
			if value == "Hey" {
				print("Correct Value\n")
			}
		case <-timer.C:
			print("Time's UP\n")
			break loop
		}

	}
}
