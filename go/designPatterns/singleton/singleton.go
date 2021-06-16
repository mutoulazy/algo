package singleton

import "sync"

type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// 饿汉式
func GetInstance() *Singleton {
	return singleton
}

var (
	lazysingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazysingleton == nil {
		once.Do(func() {
			lazysingleton = &Singleton{}
		})
	}

	return lazysingleton
}
