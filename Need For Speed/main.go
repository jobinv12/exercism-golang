package main

import "fmt"

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {

	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {

	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

	var distanceDriven int

	if car.battery < car.batteryDrain {
		distanceDriven = car.distance
	} else {
		car.battery -= car.batteryDrain
		distanceDriven += (car.distance + car.speed)
	}

	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      car.battery,
		distance:     distanceDriven,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	distanceCovered := (car.battery / car.batteryDrain) * car.speed

	return distanceCovered >= track.distance
}

func main() {

	speed := 5
	battery := 100
	batteryDrain := 2
	distance := 10

	trackDistance := 20

	fmt.Println(NewCar(speed, batteryDrain))
	fmt.Println(NewTrack(distance))
	fmt.Println(Drive(Car{battery: battery, speed: speed, batteryDrain: batteryDrain, distance: distance}))
	fmt.Println(CanFinish(Car{battery: battery, speed: speed, batteryDrain: batteryDrain, distance: distance}, Track{distance: trackDistance}))

}
