package main

type VehicleType string

const (
	CAR, BIKE, TRUCK VehicleType = "CAR", "BIKE", "TRUCK"
)

type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

type BaseVehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

func (bv *BaseVehicle) GetLicensePlate() string {
	return bv.licensePlate
}

func (bv *BaseVehicle) GetType() VehicleType {
	return bv.vehicleType
}
