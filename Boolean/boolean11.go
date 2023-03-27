package Boolean

import (
	"fmt"
	"log"
)

var (
	A    int8
	B    int8
	True bool
)

func Boolean11(bool) (True bool) {

	fmt.Printf("Boolean11\n Даны два целых числа: A, B. Проверить истинность высказывания: " +
		"«Числа A и B имеют одинаковую четность».\n")
	fmt.Printf("Введите 2 числа для сравнения: \n")
	_, err := fmt.Scan(&A, &B)
	if err != nil {
		log.Println("Вы ввели не правильные данные, нужно ввести число \n", err.Error())
		return
	}
	if A%2 == 0 && B%2 == 0 || A%2 != 0 && B%2 != 0 {
		fmt.Println(true)
	}

	return
}
