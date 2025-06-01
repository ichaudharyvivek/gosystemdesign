package models

type UserRole int

const (
	Unregistered UserRole = iota
	Registered
)

func (r UserRole) String() string {
	return [...]string{"unregistered", "registered"}[r]
}
