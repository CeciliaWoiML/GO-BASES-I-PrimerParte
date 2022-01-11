package main

import "fmt"

func main()  {
	var temperatura string
	var humedad string
	var presion string

	temperatura = "25 °C"
	humedad = "74%"
	presion = "1014 hPa"

	fmt.Println(temperatura, humedad, presion)

	// Les asignaría un string para poder indicar los datos con su símbolo o especificación correspondiente (°C, %, hPa)
}