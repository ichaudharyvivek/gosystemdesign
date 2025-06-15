package main

import (
	m "lld/trafficlight/models"
	"time"
)

func main() {
	tl1 := m.NewTrafficLight("TL1", &m.Duration{Red: 2 * time.Second, Yellow: 1 * time.Second, Green: 1 * time.Second})
	tl2 := m.NewTrafficLight("TL2", &m.Duration{Red: 1 * time.Second, Yellow: 1 * time.Second, Green: 2 * time.Second})
	tl3 := m.NewTrafficLight("TL3", &m.Duration{Red: 4 * time.Second, Yellow: 1 * time.Second, Green: 3 * time.Second})
	tl4 := m.NewTrafficLight("TL4", &m.Duration{Red: 3 * time.Second, Yellow: 1 * time.Second, Green: 4 * time.Second})

	road1 := m.NewRoad("RD1", "Brewery Street")
	road2 := m.NewRoad("RD2", "JP Street")
	road3 := m.NewRoad("RD3", "Financial Street")
	road4 := m.NewRoad("RD4", "Church Gate Street")

	road1.SetTrafficLight(tl1)
	road2.SetTrafficLight(tl2)
	road3.SetTrafficLight(tl3)
	road4.SetTrafficLight(tl4)

	tc := GetTrafficController()
	tc.AddRoad(road1)
	tc.AddRoad(road2)
	tc.AddRoad(road3)
	tc.AddRoad(road4)

	go tc.StartTraffic()

	// Start and emergency method
	time.Sleep(3 * time.Second)
	tc.HandleEmergency("RD4")

	// To hold the program forever
	select {}
}
