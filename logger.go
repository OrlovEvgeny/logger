package logger

import (
	"fmt"
	"log"
	"os"
)

//Config
type Config struct {
	AppName string
	Debug   bool
	LogFile string

	logChannel chan string
}

//New
func New(lw *Config) *log.Logger {
	lw.logChannel = make(chan string, 500)
	go lw.loop()

	prefix := fmt.Sprintf("<%s> - ", lw.AppName)
	return log.New(lw, prefix, log.LstdFlags)
}

//Write
func (lw *Config) Write(p []byte) (int, error) {
	lw.logChannel <- string(p)
	return 0, nil
}

//loop
func (lw *Config) loop() {
	logfile, err := os.OpenFile(lw.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	for {
		textLog := fmt.Sprintf("%s", <-lw.logChannel)
		if lw.Debug {
			fmt.Println(textLog)
		}
		logfile.WriteString(textLog)
	}
}
