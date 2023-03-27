package Boolean

import (
	"fmt"
	"log"
)

var (
	a1     int8
	b1     int8
	c1     int8
	IfTrue bool
)

func Boolean13(bool) (IfTrue bool) {

	fmt.Printf("Boolean13\n Даны три целых числа: A, B, C. Проверить истинность высказывания: " +
		"«Хотя бы одно из чисел A, B, C положительное».\n")
	fmt.Printf("Введите три целых числа: \n")
	_, err := fmt.Scan(&a1, &b1, &c1)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных, нужно целое число\n", err.Error())
		return
	}

	if a1 < 0 && b1 < 0 && c1 < 0 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
	return
}
