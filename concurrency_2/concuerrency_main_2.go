package concurrency_2

import (
	"log"
	"math/rand"
	"os"
	"syscall"
	"time"
)

func Run() {

	myPid := os.Getpid()

	file, err := os.OpenFile("./test2.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	for {
		log.Printf("[%v]: I will lock the file", myPid)
		errLock := syscall.Flock(int(file.Fd()), syscall.LOCK_EX)

		if errLock != nil {
			log.Printf("[%v]: Error locking file. It seems the another process is using it.", myPid)
			ranNum := rand.Int()

			time.Sleep(time.Duration(ranNum) * time.Second)

		}

		log.Printf("[%v]: Successfully locked the file", myPid)

		break
	}

	holdFileSecond := rand.Int()

	log.Printf("[%v]: I will sleep for %v seconds", myPid, holdFileSecond)

	time.Sleep(time.Duration(holdFileSecond) * time.Second)

	log.Printf("[%v]: I will unlock the file", myPid)

	errUnlock := syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	if errUnlock != nil {
		panic(errUnlock)
	}
}
