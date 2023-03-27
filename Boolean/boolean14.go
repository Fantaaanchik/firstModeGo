package Boolean

import (
	"fmt"
	"log"
)

var (
	a2    int8
	b2    int8
	c2    int8
	ITrue bool
)

func Boolean14(bool) (ITrue bool) {

	fmt.Printf("Boolean14\n Даны три целых числа: A, B, C. Проверить истинность высказывания: " +
		"«Ровно одно из чисел A, B, C положительное».\n")
	fmt.Printf("Введите три целых числа: \n")
	_, err := fmt.Scan(&a2, &b2, &c2)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных, нужно целое число\n", err.Error())
		return
	}

	if a2 > 0 && b2 <= 0 && c2 <= 0 || a2 <= 0 && b2 > 0 && c2 <= 0 || a2 <= 0 && b2 <= 0 && c2 > 0 {
		fmt.Println(true)
	}

	return
}
