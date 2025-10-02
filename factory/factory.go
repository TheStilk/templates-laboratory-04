package factory

import (
	"lab04/transport"
)

type TransportFactory interface {
	CreateTransport(model string, weight int, horsepower int) transport.Transport
}

type (
	CarFactory        struct{}
	MotorcycleFactory struct{}
	PlaneFactory      struct{}
	BicycleFactory    struct{}
)

func (f *CarFactory) CreateTransport(model string, weight int, horsepower int) transport.Transport {
	return transport.NewCar(model, weight, horsepower)
}

func (f *MotorcycleFactory) CreateTransport(model string, weight int, horsepower int) transport.Transport {
	return transport.NewMotorcycle(model, weight, horsepower)
}

func (f *PlaneFactory) CreateTransport(model string, weight int, horsepower int) transport.Transport {
	return transport.NewPlane(model, weight, horsepower)
}

func (f *BicycleFactory) CreateTransport(model string, weight int, horsepower int) transport.Transport {
	speed := weight / 10
	return transport.NewBicycle(model, speed)
}
