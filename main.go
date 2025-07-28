package main

import (
	"log"
	"nawaphon.com/concurrency_1"
	"nawaphon.com/concurrency_2"
)

func main() {
	err := concurrency_1.Run()

	if err != nil {
		log.Fatal("I don't expected error at this point.")
	}

	err = concurrency_1.Run() // should error

	if err == nil {
		log.Printf("I expected error at this point.")
	}

	err = concurrency_2.Run()

	if err != nil {
		log.Fatal("I don't expected error at this point.")
	}
}
