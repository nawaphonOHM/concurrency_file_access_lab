package main

import (
	"log"
	"nawaphon.com/concurrency_1"
	"nawaphon.com/concurrency_2"
	"os"
)

func main_1() {
	err := concurrency_1.Run()

	if err != nil {
		log.Fatal("I don't expected error at this point.")
	}

	err = concurrency_1.Run() // should error

	if err == nil {
		log.Printf("I expected error at this point.")
	}
}

func main_2() {
	err := concurrency_2.Run()

	if err != nil {
		log.Fatal("I don't expected error at this point.")
	}
}

func main() {

	params := os.Args[1:]

	if len(params) > 1 {
		log.Fatal("I don't expected more than one parameter.")
	}

	if params[0] == "1" {
		main_1()
		return
	} else if params[0] == "2" {
		main_2()
		return
	} else {
		log.Fatal("I expected only 1 or 2 as parameter.")
	}
}
