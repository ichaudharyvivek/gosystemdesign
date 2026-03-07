package main

import "sync"

type ParkingLot struct {
	levels []*Level
}

var (
	instance *ParkingLot
	once     sync.Once
)

func GetInstance() *ParkingLot {
	once.Do(func() {
		instance = &ParkingLot{levels: []*Level{}}
	})

	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnparkVehicle(vehicle Vehicle, strategy FeeStrategy) *Receipt {
	for _, level := range p.levels {
		if receipt := level.UnparkVehicle(vehicle, strategy); receipt != nil {
			return receipt
		}
	}
	return nil
}

func (p *ParkingLot) DisplayAvailability() {
	for _, level := range p.levels {
		level.DisplayAvailability()
	}
}
