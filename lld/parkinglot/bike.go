package main

func NewBike(licensePlate string) Vehicle {
	return &BaseVehicle{licensePlate: licensePlate, vehicleType: BIKE}
}
