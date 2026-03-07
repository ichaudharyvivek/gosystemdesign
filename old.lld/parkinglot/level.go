package main

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor int, numSpots int) *Level {
	level := &Level{floor: floor}
	spotNumber := 0

	carSpots := int(float64(numSpots) * 0.50)
	bikeSpots := int(float64(numSpots) * 0.40)
	truckSpots := int(float64(numSpots) * 0.10)

	for i := 0; i < carSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(spotNumber, CAR))
		spotNumber++
	}

	for i := 0; i < bikeSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(spotNumber, BIKE))
		spotNumber++
	}

	for i := 0; i < truckSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(spotNumber, TRUCK))
		spotNumber++
	}

	return level
}

func (l *Level) ParkVehicle(vehicle Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if spot.IsAvailable() && spot.GetVehicleType() == vehicle.GetType() {
			spot.ParkVehicle(vehicle)
			return true
		}
	}
	return false
}

func (l *Level) UnparkVehicle(vehicle Vehicle, strategy FeeStrategy) *Receipt {
	for _, spot := range l.parkingSpots {
		if !spot.IsAvailable() && spot.GetParkedVehicle().GetLicensePlate() == vehicle.GetLicensePlate() {
			receipt := spot.UnparkVehicle(strategy)
			return receipt
		}
	}
	return nil
}

func (l *Level) DisplayAvailability() {
	for _, spot := range l.parkingSpots {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied "
		}
		println("Level:", l.floor, ", Spot:", spot.GetSpotNumber(), ", Status:", status, ", Type:", spot.GetVehicleType())
	}
}
