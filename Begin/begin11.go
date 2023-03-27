package Begin

import (
	"fmt"
	"log"
)

var first, second float32
var Sum, Dif, Proiz, Chast float32

func Beginner11(float32, float32, float32, float32) (sum, dif, proiz, chast float32) {
	fmt.Printf("Begin11.\n Даны два ненулевых числа. Найти сумму, разность, произведение и частное их модулей.\n")
	fmt.Printf("Введите два ненулевых числа: \n")
	_, err := fmt.Scan(&first, &second)
	if err != nil {
		log.Println("Вы ввели неправильные данные: ", err.Error())
		return
	}
	sum = first + second
	if sum < 0 {
		sum = sum * (-1)
	} else {
		sum = first + second
	}

	dif = first - second
	if dif < 0 {
		dif = dif * (-1)
	} else {
		dif = first - second
	}

	proiz = first * second
	if proiz < 0 {
		proiz = proiz * (-1)
	} else {
		proiz = first * second
	}

	chast = first / second
	if chast < 0 {
		chast = chast * (-1)
	} else {
		chast = first / second
	}
	return
}
