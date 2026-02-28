package singleton

import "sync"

var (
	once     sync.Once
	instance *singleton
)

// Non exportable field so no one can create direct object
type singleton struct {
	host string
	port int
}

// Single exportable method to create object
func GetInstance(host string, port int) *singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{
				host: host,
				port: port,
			}
		}
	})

	return instance
}
