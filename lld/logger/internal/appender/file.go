package appender

import (
	"fmt"
	"os"
	"sync"
)

type FileAppender struct {
	mu   sync.Mutex
	file *os.File
}

func NewFileAppender(path string) (*FileAppender, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error in processing file: %w", err)
	}

	return &FileAppender{
		file: file,
	}, nil
}

func (a *FileAppender) Append(data []byte) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	_, err := a.file.Write(data)
	return err
}
