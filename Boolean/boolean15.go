package Boolean

import (
	"fmt"
	"log"
)

var (
	a3      int8
	b3      int8
	c3      int8
	TrueAns int8
)

func Boolean15(int8) (TrueAns bool) {

	fmt.Printf("Boolean15\n Даны три целых числа: A, B, C. Проверить истинность высказывания: " +
		"«Ровно два из чисел A, B, C являются положительными \n")
	fmt.Printf("Введите три числа: \n")
	_, err := fmt.Scan(&a3, &b3, &c3)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных. Нужно целое число\n", err.Error())
		return
	}

	if a3 > 0 && b3 > 0 && c3 <= 0 || a3 <= 0 && b3 > 0 && c3 > 0 || a3 > 0 && b3 <= 0 && c3 > 0 {
		fmt.Println(true)
	}
	return
}
