package logging

import (
	"log"
	"os"
)

func LogError(errs error) {
	data := errs.Error()
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0666))
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.Println(data)
}
