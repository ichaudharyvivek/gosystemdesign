package models

type Road struct {
	ID           string
	Name         string
	TrafficLight *TrafficLight
}

func NewRoad(id, name string) *Road {
	return &Road{ID: id, Name: name}
}

func (r *Road) SetTrafficLight(tl *TrafficLight) {
	r.TrafficLight = tl
}
