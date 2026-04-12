package appender

import (
	"fmt"
	"lld-logger/internal/formatter"
	"lld-logger/internal/model"
	"os"
	"sync"
)

type FileAppender struct {
	mu        sync.Mutex
	file      *os.File
	formatter formatter.Formatter
}

func NewFileAppender(path string, formatter formatter.Formatter) (*FileAppender, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error in processing file: %w", err)
	}

	return &FileAppender{
		file:      file,
		formatter: formatter,
	}, nil
}

func (a *FileAppender) Append(record model.Record) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	data := a.formatter.Format(record)
	_, err := a.file.Write(data)
	return err
}

func (a *FileAppender) Close() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.file.Close()
}
