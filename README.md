# Easy logger

[![Build Status](https://travis-ci.org/OrlovEvgeny/logger.svg?branch=master)](https://travis-ci.org/OrlovEvgeny/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/OrlovEvgeny/logger?v1)](https://goreportcard.com/report/github.com/OrlovEvgeny/logger)
[![GoDoc](https://godoc.org/github.com/OrlovEvgeny/logger?status.svg)](https://godoc.org/github.com/OrlovEvgeny/logger)


easy logger golang mini module

# Use

````go
	//New return *log.Logger (https://godoc.org/log#Logger)
	log := New(&Config{
		AppName: "you application name",
		Debug:   true,
		LogFile: "./error.log",
	})

	log.Println("custom Error message")
````


# License:

[MIT](LICENSE)