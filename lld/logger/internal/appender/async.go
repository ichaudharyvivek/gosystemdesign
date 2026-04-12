package appender

import (
	"sync"

	"lld-logger/internal/model"
)

type AsyncAppender struct {
	appender Appender
	quit     chan struct{}
	wg       sync.WaitGroup
	ch       chan model.Record
}

func NewAsyncAppender(appender Appender, buffer int, workers int) *AsyncAppender {
	if workers <= 0 {
		workers = 1
	}

	if buffer <= 0 {
		buffer = 1000
	}

	a := &AsyncAppender{
		ch:       make(chan model.Record, buffer),
		appender: appender,
		quit:     make(chan struct{}),
	}

	a.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go a.worker()
	}

	return a
}

func (a *AsyncAppender) worker() {
	defer a.wg.Done()

	for {
		select {
		case msg := <-a.ch:
			a.appender.Append(msg)

		// When quit, drain the channel before exiting
		case <-a.quit:
			for {
				select {
				case msg := <-a.ch:
					a.appender.Append(msg)
				default:
					return
				}
			}
		}
	}
}

func (a *AsyncAppender) Append(record model.Record) error {
	select {
	case a.ch <- record:
		return nil
	default:
		// Drop logs when buffer is full
		// Can be extended to add a user defined buffer policy
		return nil
	}
}

func (a *AsyncAppender) Close() error {
	close(a.quit)
	a.wg.Wait()
	return nil
}
