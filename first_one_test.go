package main

import (
    "testing"
	"math"
)

func TestKelvinToCelsius(t *testing.T) {
    tests := []struct {
        input    kelvin
        expected celsius
    }{
        {0, -273.15},
        {273.15, 0},
        {373.15, 100},
    }

    for _, test := range tests {
        result := test.input.celsius()
        if math.Abs(float64(result - test.expected)) > 1e-3 {
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}

func TestKelvinToFahrenheit(t *testing.T) {
    tests := []struct {
        input    kelvin
        expected fahrenheit
    }{
        {0, -459.67},
        {273.15, 32},
        {373.15, 212},
    }

    for _, test := range tests {
        result := test.input.fahrenheit()
        if math.Abs(float64(result - test.expected)) > 1e-3 {
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}

func TestCelsiusToFahrenheit(t *testing.T) {
    tests := []struct {
        input    celsius
        expected fahrenheit
    }{
        {0, 32},
        {100, 212},
        {-40, -40},
    }

    for _, test := range tests {
        result := test.input.fahrenheit()
        if math.Abs(float64(result - test.expected)) > 1e-3 {
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}

func TestCelsiusToKelvin(t *testing.T) {
    tests := []struct {
        input    celsius
        expected kelvin
    }{
        {0, 273.15},
        {100, 373.15},
        {-273.15, 0},
    }

    for _, test := range tests {
        result := test.input.kelvin()
        if math.Abs(float64(result - test.expected)) > 1e-3 {
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}

func TestFahrenheitToCelsius(t *testing.T) {
    tests := []struct {
        input    fahrenheit
        expected celsius
    }{
        {32, 0},
        {212, 100},
        {-40, -40},
    }

    for _, test := range tests {
        result := test.input.celsius()
        if math.Abs(float64(result - test.expected)) > 1e-3 {
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}

func TestFahrenheitToKelvin(t *testing.T) {
    tests := []struct {
        input    fahrenheit
        expected kelvin
    }{
        {32, 273.15},
        {212, 373.15},
        {-459.67, 0},
    }

    for _, test := range tests {
        result := test.input.kelvin()
        if math.Abs(float64(result - test.expected)) > 1e-3 {
            t.Errorf("Expected %v, but got %v", test.expected, result)
        }
    }
}