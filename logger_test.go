package logger

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const path = "./test_logfile.log"

func BenchmarkSubstring(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	log := New(&Config{
		AppName: "test Logger module",
		Debug:   false,
		LogFile: path,
	})
	for i := 0; i < b.N; i++ {
		log.Println(RandStringRunes(10))
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lenLetters = len(letterRunes)

func RandStringRunes(n int) string {
	builder := &strings.Builder{}
	for i := 0; i < n; i++ {
		builder.WriteRune(letterRunes[rand.Intn(lenLetters)])
	}
	return builder.String()
}
