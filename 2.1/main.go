package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Celsius = 0
	FreezingK     Celsius = 1
	BoilingK      Celsius = 373.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(CToK(FToC(f))) }

func KToC(k Kelvin) Celsius    { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

func main() {
	fmt.Printf("100F in K: %v\n", FToK(100))
	fmt.Printf("100K in F:%v\n", KToF(100))
}
