package singleton

import "sync"

var instance *singleton
var once sync.Once

type singleton struct {
}

func GetInstance() *singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{}
		}
	})
	return instance
}
