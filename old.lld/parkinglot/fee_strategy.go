package main

import "time"

type FeeStrategy interface {
	CalculateFee(parkedAt, unparkedAt time.Time, vehicleType VehicleType) float64
}

type HourlyFeeStrategy struct {
	rates map[VehicleType]float64
}

func NewHourlyFeeStrategy() *HourlyFeeStrategy {
	return &HourlyFeeStrategy{
		rates: map[VehicleType]float64{
			CAR:   3.0,
			BIKE:  2.0,
			TRUCK: 1.5,
		},
	}
}

func (h *HourlyFeeStrategy) CalculateFee(parkedAt, unparkedAt time.Time, vehicleType VehicleType) float64 {
	duration := unparkedAt.Sub(parkedAt)
	hours := float64(duration.Hours())
	if hours < 1.0 {
		return 1.0
	}

	rate, exists := h.rates[vehicleType]
	if !exists {
		return 0.0
	}

	return hours * rate
}
