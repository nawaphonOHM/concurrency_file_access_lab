package concurrency_2

import (
	"log"
	"math/rand"
	"os"
	"syscall"
	"time"
)

func Run() error {

	myPid := os.Getpid()

	file, err := os.OpenFile("./test2.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	for {
		log.Printf("[%v]: I will lock the file", myPid)
		errLock := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)

		if errLock != nil {
			log.Printf("[%v]: Error locking file. It seems the another process is using it.", myPid)
			ranNum := rand.Intn(10)

			time.Sleep(time.Duration(ranNum) * time.Second)

			continue
		}

		log.Printf("[%v]: Successfully locked the file", myPid)

		break
	}

	holdFileSecond := rand.Intn(10)

	log.Printf("[%v]: I will sleep for %v seconds", myPid, holdFileSecond)

	time.Sleep(time.Duration(holdFileSecond) * time.Second)

	log.Printf("[%v]: I will unlock the file", myPid)

	errUnlock := syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	if errUnlock != nil {
		return errUnlock
	}

	return nil
}
