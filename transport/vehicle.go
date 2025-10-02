package transport

import "fmt"

type Transport interface {
	Move()
	FuelUp()
	GetModel() string
}

type Motorized interface {
	GetHorsePower() int
	GetWeight() int
}

type Car struct {
	Model      string
	Weight     int
	HorsePower int
}

func NewCar(model string, weight int, horsepower int) *Car {
	if horsepower <= 0 {
		panic("horsepower must be > 0")
	}
	return &Car{
		Model:      model,
		Weight:     weight,
		HorsePower: horsepower,
	}
}

func (c *Car) Move() {
	powerToWeight := float64(c.HorsePower) / float64(c.Weight)
	var description string
	if powerToWeight < 0.05 {
		description = "slowly"
	} else if powerToWeight < 0.1 {
		description = "middle"
	} else if powerToWeight < 0.2 {
		description = "fast"
	} else {
		description = "very fast"
	}
	fmt.Printf("The car %s can accelerate to 100 km/h %s | power-to-weight: %.4f\n", c.Model, description, powerToWeight)
}

func (c *Car) FuelUp() {
	fmt.Printf("Filling the car %s with petrol\n", c.Model)
}

func (c *Car) GetModel() string   { return c.Model }
func (c *Car) GetHorsePower() int { return c.HorsePower }
func (c *Car) GetWeight() int     { return c.Weight }

type Motorcycle struct {
	Model      string
	Weight     int
	HorsePower int
}

func NewMotorcycle(model string, weight int, horsepower int) *Motorcycle {
	if horsepower <= 0 {
		panic("horsepower must be > 0")
	}
	return &Motorcycle{
		Model:      model,
		Weight:     weight,
		HorsePower: horsepower,
	}
}

func (m *Motorcycle) Move() {
	powerToWeight := float64(m.HorsePower) / float64(m.Weight)
	var description string
	if powerToWeight < 0.1 {
		description = "slowly"
	} else if powerToWeight < 0.2 {
		description = "middle"
	} else if powerToWeight < 0.4 {
		description = "fast"
	} else {
		description = "very fast"
	}
	fmt.Printf("The moto %s can accelerate to 100 km/h %s | power-to-weight: %.4f)\n", m.Model, description, powerToWeight)
}

func (m *Motorcycle) FuelUp() {
	fmt.Printf("Filling the moto %s with petrol\n", m.Model)
}

func (m *Motorcycle) GetModel() string   { return m.Model }
func (m *Motorcycle) GetHorsePower() int { return m.HorsePower }
func (m *Motorcycle) GetWeight() int     { return m.Weight }

type Plane struct {
	Model      string
	Weight     int
	HorsePower int
}

func NewPlane(model string, weight int, horsepower int) *Plane {
	if horsepower <= 0 {
		panic("horsepower must be > 0")
	}
	return &Plane{
		Model:      model,
		Weight:     weight,
		HorsePower: horsepower,
	}
}

func (p *Plane) Move() {
	powerToWeight := float64(p.HorsePower) / float64(p.Weight)
	description := "extremely fast"
	if powerToWeight < 0.2 {
		description = "fast"
	}
	fmt.Printf("The plane %s flies %s | power-to-weight: %.4f\n", p.Model, description, powerToWeight)
}

func (p *Plane) FuelUp() {
	fmt.Printf("Refueling plane %s with aviation fuel\n", p.Model)
}

func (p *Plane) GetModel() string   { return p.Model }
func (p *Plane) GetHorsePower() int { return p.HorsePower }
func (p *Plane) GetWeight() int     { return p.Weight }

type Bicycle struct {
	Model string
	Speed int
}

func NewBicycle(model string, speed int) *Bicycle {
	return &Bicycle{Model: model, Speed: speed}
}

func (b *Bicycle) Move() {
	fmt.Printf("Bicycle %s is rolling down the track at a speed of %d km/h\n", b.Model, b.Speed)
}

func (b *Bicycle) FuelUp() {
	fmt.Printf("Why do you need fuel when you have muscles?\n")
}

func (b *Bicycle) GetModel() string { return b.Model }
