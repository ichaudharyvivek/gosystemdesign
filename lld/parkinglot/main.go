package main

import (
	"fmt"
	"time"
)

func main() {
	parkingLot := GetInstance()
	parkingLot.AddLevel(NewLevel(1, 10))
	parkingLot.AddLevel(NewLevel(2, 10))

	// Vehicles
	car := NewCar("CAR-101")
	bike := NewBike("BIKE-202")
	truck := NewTruck("TRUCK-303")

	// Fee strategy
	feeStrategy := NewHourlyFeeStrategy()

	fmt.Println("Initial Availability:")
	parkingLot.DisplayAvailability()

	// Park vehicles
	fmt.Println("\n--- Parking Vehicles ---")
	parkingLot.ParkVehicle(car)
	parkingLot.ParkVehicle(bike)
	parkingLot.ParkVehicle(truck)

	fmt.Println("\nAvailability After Parking:")
	parkingLot.DisplayAvailability()

	// Simulate time
	fmt.Println("\n--- Simulating parking time (sleeping for 2 seconds) ---")
	time.Sleep(2 * time.Second)

	// Unpark vehicles
	fmt.Println("\n--- Unparking Vehicles ---")
	if receipt := parkingLot.UnparkVehicle(car, feeStrategy); receipt != nil {
		receipt.Print()
	}

	fmt.Println("\nAvailability After Unparking:")
	parkingLot.DisplayAvailability()
}
