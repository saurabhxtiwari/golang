// Workers
// Channel push
// everybody is waiting for something on channel

package main

import (
	"fmt"
	"strconv"
	"time"
)

func initPool(jobChannel chan string, resultChannel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Created Thread : ", strconv.Itoa(i))
		go func(j int) {
			for {

				x := <-jobChannel
				time.Sleep(time.Second * 3)
				fmt.Println("Job executed by ", strconv.Itoa(j))
				resultChannel <- x
			}
		}(i)
	}
}

func execute(job string, jobChannel chan string, resultChannel chan string) {
	jobChannel <- job
	x := <-resultChannel
	fmt.Println("Received response ", x)
}

func main() {
	jobChannel := make(chan string)
	resultChannel := make(chan string)

	initPool(jobChannel, resultChannel)

	time.Sleep(time.Second * 10)
	go execute("One", jobChannel, resultChannel)
	go execute("Two", jobChannel, resultChannel)
	go execute("Three", jobChannel, resultChannel)
	go execute("Four", jobChannel, resultChannel)
	go execute("Five", jobChannel, resultChannel)
	go execute("Six", jobChannel, resultChannel)
	go execute("Seven", jobChannel, resultChannel)
	go execute("Eight", jobChannel, resultChannel)
	go execute("Nine", jobChannel, resultChannel)
	go execute("Ten", jobChannel, resultChannel)

	time.Sleep(time.Second * 15)
	close(jobChannel)
	close(resultChannel)
}
