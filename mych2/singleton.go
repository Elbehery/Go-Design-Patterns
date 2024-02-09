package mych2

type Singleton interface {
	AddOne()
	GetCounter() int
}

type singleton struct {
	counter int
}

func (s *singleton) AddOne() {
	s.counter++
}

func (s *singleton) GetCounter() int {
	return s.counter
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
