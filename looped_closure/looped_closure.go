package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	const routines = 25

	results := make(chan string)

	for i := 1; i <= routines; i++ {
		go func(i interface{}) {
			results <- fmt.Sprintf("Hello from Go Routine #%v", i)
		}(i)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	var resultCount = 0

	for {
		select {
		case <-c:
			fmt.Println("I've been interrupted!")
			os.Exit(0)
		case result := <-results:
			fmt.Println(result)
			resultCount++

			if resultCount == routines {
				fmt.Println("All done.")
				os.Exit(0)
			}
		}
	}
}
