// Package tempconv performe Celsius and Fahrenheit temperature computations.package tempconv

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroe Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Println("%g\n", BoilingC-FreezingC) // "100" ˚C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" ˚F
	fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
}
