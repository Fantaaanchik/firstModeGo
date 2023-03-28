package Boolean

import (
	"fmt"
	"log"
)

func Boolean20() {
	var (
		numa      int16
		firstNum  int16
		secondNum int16
		thirdNum  int16
	)
	fmt.Printf("Boolean20\n Дано трехзначное число. Проверить истинность высказывания: " +
		"«Все цифры данного числа различны»\n")
	fmt.Printf("Введите трехзначное число: \n")
	_, err := fmt.Scan(&numa)
	if err != nil {
		log.Println("Введенные данные не являются числом\n", err.Error())
		return
	}
	firstNum = numa / 100
	secondNum = numa % 100 / 10
	thirdNum = numa % 10

	if firstNum == secondNum || secondNum == thirdNum || firstNum == thirdNum {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

	return
}
