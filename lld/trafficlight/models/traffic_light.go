package models

import (
	"fmt"
	"time"
)

type Duration struct {
	Red    time.Duration
	Yellow time.Duration
	Green  time.Duration
}

type TrafficLight struct {
	ID            string
	CurrentSignal Signal
	Duration      Duration
}

func NewTrafficLight(id string, duration *Duration) *TrafficLight {
	return &TrafficLight{ID: id, CurrentSignal: RED, Duration: *duration}
}

func (tl *TrafficLight) ChangeSignal(newSignal Signal) {
	tl.CurrentSignal = newSignal
	tl.NotifyAll()
}

func (tl *TrafficLight) NotifyAll() {
	fmt.Printf("Traffic Light '%s' changed to '%s'\n", tl.ID, tl.CurrentSignal)
}
