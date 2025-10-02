package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lab04/factory"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Transport Factory System")
	fmt.Println("Available types: car, motorcycle, plane, bicycle")
	fmt.Print("Enter transport type: ")
	input, _ := reader.ReadString('\n')
	transportType := strings.TrimSpace(strings.ToLower(input))

	fmt.Print("Enter model: ")
	model, _ := reader.ReadString('\n')
	model = strings.TrimSpace(model)

	var weight, horsepower int
	if transportType != "bicycle" {
		fmt.Print("Enter weight (kg): ")
		weightStr, _ := reader.ReadString('\n')
		weight, _ = strconv.Atoi(strings.TrimSpace(weightStr))

		fmt.Print("Enter horsepower: ")
		hpStr, _ := reader.ReadString('\n')
		horsepower, _ = strconv.Atoi(strings.TrimSpace(hpStr))
	} else {
		fmt.Print("Enter approximate weight (kg, used to estimate speed): ")
		weightStr, _ := reader.ReadString('\n')
		weight, _ = strconv.Atoi(strings.TrimSpace(weightStr))
		horsepower = 0
	}

	var f factory.TransportFactory
	switch transportType {
	case "car":
		f = &factory.CarFactory{}
	case "motorcycle":
		f = &factory.MotorcycleFactory{}
	case "plane":
		f = &factory.PlaneFactory{}
	case "bicycle":
		f = &factory.BicycleFactory{}
	default:
		fmt.Printf("Unknown transport type: %s\n", transportType)
		return
	}

	vehicle := f.CreateTransport(model, weight, horsepower)

	fmt.Println("\nCreated vehicle:")
	vehicle.FuelUp()
	vehicle.Move()
	fmt.Printf("Model: %s\n", vehicle.GetModel())
}
