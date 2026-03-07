package main

import (
	"fmt"
	m "lld/trafficlight/models"
	"sync"
	"time"
)

var (
	instance *TrafficController
	once     sync.Once
)

type TrafficController struct {
	roads map[string]*m.Road
	mu    sync.Mutex
}

func GetTrafficController() *TrafficController {
	once.Do(func() {
		if instance == nil {
			instance = &TrafficController{roads: make(map[string]*m.Road)}
		}
	})

	return instance
}

func (tc *TrafficController) AddRoad(road *m.Road) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.roads[road.ID] = road
}

func (tc *TrafficController) StartTraffic() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	for _, road := range tc.roads {
		trafficLight := road.TrafficLight
		go func(tl *m.TrafficLight) {
			for {
				time.Sleep(trafficLight.Duration.Red)
				tl.ChangeSignal(m.YELLOW)

				time.Sleep(tl.Duration.Yellow)
				tl.ChangeSignal(m.GREEN)

				time.Sleep(tl.Duration.Green)
				tl.ChangeSignal(m.RED)
			}
		}(trafficLight)
	}
}

func (tc *TrafficController) HandleEmergency(roadId string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	road, exists := tc.roads[roadId]
	if !exists {
		fmt.Println("Road don't exists..")
	}

	fmt.Printf("Emergency at road %s\n", road.ID)
	tl := road.TrafficLight
	tl.ChangeSignal(m.GREEN)
}
