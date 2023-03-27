package Begin

import (
	"fmt"
	"log"
)

var (
	x uint8
	Y uint8
)

func Beginner25(uint8) (Y uint8) {

	fmt.Println("Begin25\n Найти значение функции y = 3x^6 − 6x^2 − 7 при данном значении x.")
	fmt.Printf("Введите значение x: \n")
	_, err := fmt.Scan(&x)
	if err != nil {
		log.Println("Введен не правильный формат данных", err.Error())
		return
	}
	Y = 3*(x*x*x*x*x*x) - (6 * (x * x)) - 7
	return
}
