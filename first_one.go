// Assigning a function to a variable
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		fmt.Println(sensor())
		time.Sleep(time.Second)
	}
}

func main() {
	measureTemperature(3, fakeSensor)
	fmt.Println("-----------------")
	measureTemperature(3, realSensor)
}
