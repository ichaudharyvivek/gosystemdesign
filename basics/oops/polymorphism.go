// go:build polymorphism
package main

import (
	"fmt"
	"time"
)

// Base logger object(class in java)
type Logger struct {
	Prefix string
}

func (l *Logger) Log(message string) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] %s: %s\n", timestamp, l.Prefix, message)
}

type Loggable interface {
	Log(message string)
}

// specialized logger - FileLogger
// overridden method for Log
type FileLogger struct {
	Logger
	FilePath string
}

func (fl *FileLogger) Log(message string) {
	fmt.Printf("Writing to file: %s\n", fl.FilePath)
	fl.Logger.Log(message)

}

// specialized logger - DBLogger
// overridden method for Log
type DBLogger struct {
	Logger
	Connection string
}

func (dbl *DBLogger) Log(message string) {
	fmt.Printf("Writing to db with connection: %s\n", dbl.Connection)
	dbl.Logger.Log(message)
}

func main() {
	fl := FileLogger{
		Logger:   Logger{Prefix: "FILE"},
		FilePath: "/c/some/path/",
	}
	fl.Log("Successfully Added to File.")

	fmt.Println()

	dl := DBLogger{
		Logger:     Logger{Prefix: "DB"},
		Connection: "jdbc://mysql:8080/testDB",
	}
	dl.Log("Successfully connected to Database.")

}
