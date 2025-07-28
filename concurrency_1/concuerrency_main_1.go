package concurrency_1

import (
	"log"
	"os"
	"time"
)

func Run() error {

	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE, os.ModeExclusive)
	if err != nil {
		return err
	}

	log.Println("File opened!! next I will sleep for 10 seconds")

	time.Sleep(10 * time.Second)

	log.Println("I will exist the file should close.")

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return nil
}
