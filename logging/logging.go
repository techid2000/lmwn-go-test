package logging

import (
	"log"
	"os"
)

func GetInfoLogger() *log.Logger {
	return log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
}
