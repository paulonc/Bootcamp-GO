package main

import "fmt"

const ebulicaoKelvin = 373.2


func kelvinParaCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	tempKelvin := ebulicaoKelvin
	tempCelsius := kelvinParaCelsius(tempKelvin)

	fmt.Printf("Ponto de ebulição da água em °K é: %.2f\n", tempKelvin)
	fmt.Printf("Ponto de ebulição da água em °C é: %.2f\n", tempCelsius)
}