package Begin

import (
	"fmt"
	"log"
)

var (
	xSho      float32
	aSho      float32
	yIr       float32
	bIr       float32
	OneKgIr   float32
	OneKgSho  float32
	Different float32
)

func Beginner34(float32, float32, float32) (OneKgSho, OneKgIr, Different float32) {
	fmt.Printf("Begin34\n Известно, что X кг шоколадных конфет стоит A рублей, а Y кг ирисок стоит B рублей. " +
		"Определить, сколько стоит 1 кг шоколадных конфет, 1 кг ирисок, а также во сколько раз шоколадные конфеты дороже ирисок.\n")
	fmt.Println("Введите 4 числа, X, A, Y, B: ")
	_, err := fmt.Scan(&xSho, &aSho, &yIr, &bIr)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных\n", err.Error())
		return
	}

	OneKgSho = aSho / xSho
	OneKgIr = bIr / yIr
	Different = OneKgSho / OneKgIr
	fmt.Printf("Один килограмм шоколадныых конфет стоит %v рублей\n", OneKgSho)
	fmt.Printf("Один килограмм ирисок стоит %v рублей\n", OneKgIr)
	fmt.Printf("Шоколадные конфеты дороже ирисок в %v раз\n", Different)
	return
}
