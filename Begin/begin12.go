package Begin

import (
	"fmt"
	"log"
	"math"
)

var a, b float32
var ResC float32
var ResP float32

func Beginner12(float32, float32) (ResC, ResP float32) {
	fmt.Printf("Даны катеты прямоугольного треугольника a и b. Найти его гипотенузу c и периметр P:\n" +
		"c = sqrt (a^2 + b^2), P = a + b + c.\n")
	fmt.Printf("Введите два числа: \n")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Println("Вы не правильно ввели данные ", err.Error())
		return
	}
	ResC = float32(math.Sqrt(float64((a * a) + (b * b))))
	ResP = a + b + ResC
	return
}
