package singleton

import "sync"

var (
	once     sync.Once
	instance *singleton
)

// Non exportable field so no one can create direct object
type singleton struct{}

// Single exportable method to create object
func GetInstance() *singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{}
		}
	})

	return instance
}
