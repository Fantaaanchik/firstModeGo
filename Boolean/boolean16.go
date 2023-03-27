package Boolean

import (
	"fmt"
	"log"
)

var (
	oneNum         int8
	ResOfOperation bool
)

func Boolean16(bool) (ResOfOperation bool) {

	fmt.Printf("Boolean16\n Дано целое положительное число. Проверить истинность высказывания: " +
		"«Данное число является четным двузначным». \n")
	fmt.Printf("Введите целое положительное число: \n")
	_, err := fmt.Scan(&oneNum)
	if err != nil {
		log.Println("Введенные вами данные не является числом\n", err.Error())
		return
	}

	if oneNum%2 == 0 && oneNum > 9 && oneNum < 100 {
		println(true)
	}

	return
}
