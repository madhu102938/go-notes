// Assigning a function to a variable
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func measureTemperature(samples int, sensor sensor) {
	for i := 0; i < samples; i++ {
		fmt.Println(sensor())
		time.Sleep(time.Second)
	}
}	

// anonymous functions or function literals
var min = func(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	rfunc := calibrate(realSensor, 5)
	// rfunc keeps reference to the realSensor function and the offset value
	measureTemperature(3, rfunc)

	fmt.Println("-----------")

	var k kelvin = 294
	sensor := func() kelvin {
		return k
	}

	fmt.Println(sensor())
	k++
	fmt.Println(sensor()) // sensor keeps reference to the k variable
}
