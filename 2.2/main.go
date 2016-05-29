package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Temperature
type Celsius float64
type Fahrenheit float64

// Length
type Feet float64
type Meter float64

// Weight
type Kilogram float64
type Pound float64

// Temperature
func (c Celsius) String() string               { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string            { return fmt.Sprintf("%g°F", f) }
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Length
func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string  { return fmt.Sprintf("%g meter", m) }
func FeetToMeter(f Feet) Meter  { return Meter(f * 0.3048) }
func MeterToFfeet(m Meter) Feet { return Feet(m / 0.3048) }

// Weight
func (k Kilogram) String() string      { return fmt.Sprintf("%g kilogram", k) }
func (p Pound) String() string         { return fmt.Sprintf("%g pound", p) }
func KilogramToPound(k Kilogram) Pound { return Pound(k / 0.453592) }
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.453592) }

func main() {
	flag.Parse()
	numbers := flag.Args()

	if len(numbers) >= 1 {
		convert(numbers)
	} else {
		for true {
			var number string
			_, err := fmt.Scanf("%s", &number)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			numbers := []string{number}
			convert(numbers)
		}
	}
}

func convert(args []string) {
	for _, n := range args {
		f, err := strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(Celsius(f).String(), " = ", CelsiusToFahrenheit(Celsius(f)).String())
		fmt.Println(Fahrenheit(f).String(), " = ", FahrenheitToCelsius(Fahrenheit(f)).String())
		fmt.Println()

		fmt.Println(Feet(f).String(), " = ", FeetToMeter(Feet(f)).String())
		fmt.Println(Meter(f).String(), " = ", MeterToFfeet(Meter(f)).String())
		fmt.Println()

		fmt.Println(Kilogram(f).String(), " = ", KilogramToPound(Kilogram(f)).String())
		fmt.Println(Pound(f).String(), " = ", PoundToKilogram(Pound(f)).String())
		fmt.Println()
	}
}
