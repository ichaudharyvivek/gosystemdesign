package main

import (
	"fmt"
	"time"
)

type Receipt struct {
	Vehicle    Vehicle
	Spot       *ParkingSpot
	ParkedAt   time.Time
	UnparkedAt time.Time
	Fee        float64
}

func (r *Receipt) Print() {
	fmt.Printf("Receipt for %s\n", r.Vehicle.GetLicensePlate())
	fmt.Printf("Spot: %d\n", r.Spot.GetSpotNumber())
	fmt.Printf("Type: %s\n", r.Spot.GetVehicleType())
	fmt.Printf("Parked At: %s\n", r.ParkedAt.Format(time.RFC1123))
	fmt.Printf("Unparked At: %s\n", r.UnparkedAt.Format(time.RFC1123))
	fmt.Printf("Total Fee: ₹%.2f\n", r.Fee)
}
