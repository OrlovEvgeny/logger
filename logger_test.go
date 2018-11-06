package logger

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	fileName   = "./test_logfile.log"
	writeLines = 50
)

//TestWrite
func TestWrite(t *testing.T) {
	log := New(&Config{
		AppName: "test Logger module",
		Debug:   false,
		LogFile: fileName,
	})

	for i := 0; i < writeLines; i++ {
		log.Println(randSeq(10))
	}

	time.Sleep(time.Second * 1)
}

//TestCount
func TestCount(t *testing.T) {
	file, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Count the lines.
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	if count != writeLines {
		t.Errorf("count != %d\n", writeLines)
	}
}

//TestDeleteFile
func TestDeleteFile(t *testing.T) {
	err := os.Remove(fileName)
	if err != nil {
		t.Error(err)
	}
}

//letters
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//randSeq
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
