package main

import (
	"time"
)

type ParkingSpot struct {
	spotNumber    int
	spotType      VehicleType
	parkedVehicle Vehicle
	parkedAt      time.Time
}

func NewParkingSpot(spotNumber int, spotType VehicleType) *ParkingSpot {
	return &ParkingSpot{spotNumber: spotNumber, spotType: spotType}
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps.parkedVehicle == nil
}

func (ps *ParkingSpot) ParkVehicle(vehicle Vehicle) {
	if ps.IsAvailable() && vehicle.GetType() == ps.spotType {
		ps.parkedVehicle = vehicle
		ps.parkedAt = time.Now()
	}
}

func (ps *ParkingSpot) UnparkVehicle(strategy FeeStrategy) *Receipt {
	if ps.parkedVehicle == nil {
		return nil
	}

	now := time.Now()
	fee := strategy.CalculateFee(ps.parkedAt, now, ps.spotType)
	receipt := &Receipt{
		Vehicle:    ps.parkedVehicle,
		Spot:       ps,
		ParkedAt:   ps.parkedAt,
		UnparkedAt: now,
		Fee:        fee,
	}
	ps.parkedVehicle = nil
	return receipt
}

func (ps *ParkingSpot) GetSpotNumber() int {
	return ps.spotNumber
}

func (ps *ParkingSpot) GetVehicleType() VehicleType {
	return ps.spotType
}

func (ps *ParkingSpot) GetParkedVehicle() Vehicle {
	return ps.parkedVehicle
}
