package Boolean

import (
	"fmt"
	"log"
)

var (
	a           int8
	b           int8
	c           int8
	TrueChecker bool
)

func Boolean12(bool) (TrueChecker bool) {

	fmt.Printf("Boolean12\n Даны три целых числа: A, B, C. Проверить истинность высказывания: " +
		"«Каждое из чисел A, B, C положительное».\n")
	fmt.Printf("Введите три целых числа: \n")
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных, нужно целое число\n", err.Error())
		return
	}

	if a > 0 && b > 0 && c > 0 {
		fmt.Println(true)
	}

	return
}
