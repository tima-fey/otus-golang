package main

import (
	"errors"
	"fmt"
	"time"
)

func taskExecutor(tc, ec, sc chan func() error, finish chan bool, M int) {
	for aTask := range tc {
		if len(ec) >= M {
			break
		}
		err := aTask()
		if err != nil {
			ec <- aTask
			if len(ec) >= M {
				break
			}
		} else {
			sc <- aTask
		}
	}
	finish <- true
}

func run(task []func() error, N int, M int) (int, int, error) {
	taskChannel := make(chan func() error, len(task))
	errorsChannel := make(chan func() error, len(task))
	successfulChannel := make(chan func() error, len(task))
	excexutorChannel := make(chan bool, N)

	for _, aTask := range task {
		taskChannel <- aTask
	}
	close(taskChannel)
	for counter := 0; counter < N; counter++ {
		go taskExecutor(taskChannel, errorsChannel, successfulChannel, excexutorChannel, M)
	}
	finishedTaskCounter := 0
	for range excexutorChannel {
		finishedTaskCounter++
		if finishedTaskCounter == N {
			break
		}
	}
	if len(errorsChannel) >= M {
		return len(successfulChannel), len(errorsChannel), errors.New("too many failed tasks. exiting")
	}
	return len(successfulChannel), len(errorsChannel), nil
}

func Run(task []func() error, N int, M int) error {
	_, _, err := run(task, N, M)
	return err
}

func st() error {
	time.Sleep(3 * time.Millisecond)
	return nil
}

func ft() error {
	time.Sleep(4 * time.Millisecond)
	return errors.New("error")
}

func main() {
	myTasks := []func() error{st, ft, st}
	_ = Run(myTasks, 2, 1)
	fmt.Println("Done")
}
