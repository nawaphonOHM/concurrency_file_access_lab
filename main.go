package main

import "nawaphon.com/concurrency_1"

func main() {
	concurrency_1.Run()

	concurrency_1.Run() // should error
}
