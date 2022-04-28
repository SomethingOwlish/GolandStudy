package main

import (
	"fmt"
	"math/rand"
)

// Объявляет функцию под названием main
func main() {

	fmt.Println("Hello, playground")
	fmt.Println("こんにちは Здравствуйте Hola")
	fmt.Print("Мой вес на поверхности Марса равен ")
	fmt.Print(55.0 * 0.3783) // В результате 20.8065
	fmt.Print(" килограммам, а мой возраст равен ")
	fmt.Print(41 * 365 / 687) // В результате 21
	fmt.Print(" годам.")
	fmt.Println("Мой вес на поверхности Марса равен", 55.0*0.3783, " килограммам, а мой возраст равен", 41*365/687, " годам.")
	fmt.Printf("а мой возраст равен %v годам.\n", 41*365/687)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	const lightSpeed = 299792 // км/с
	var distance = 56000000   // км

	fmt.Println(distance/lightSpeed, "секунд") // В результате 186 секунд

	distance = 401000000
	fmt.Println(distance/lightSpeed, "секунд") // В результате 1337 секунд

	distance = 96300000
	fmt.Println(distance/lightSpeed, "секунд") //Долетят
	var (
		speedBySec = distance / lightSpeed
		min        = speedBySec / 60
		hr         = min / 60
		day        = hr / 24
	)
	fmt.Printf("А в сутках это будет всего лишь %v дня", day)

	var weight = 10
	weight -= 2

	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	distance = rand.Intn(345000001) + 56000000
	fmt.Println(distance)

	const hoursPerDay = 24

	var days = 28

	fmt.Println(distance/(days*hoursPerDay), "км/ч")
}
