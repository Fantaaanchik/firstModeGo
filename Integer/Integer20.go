package Integer

import (
	"fmt"
	"log"
)

var (
	seconds      int32
	ResultInHour int32
)

func Integer20(int32) (ResultInHour int32) {

	fmt.Printf("Integer19\n С начала суток прошло N секунд (N — целое). Найти количество полных часов, прошедших " +
		"с начала суток\n")
	fmt.Printf("Введите челое число, количество секунд: \n")
	_, err := fmt.Scan(&seconds)
	if err != nil {
		log.Println("Вы ввели не правильные данные \n", err.Error())
		return
	}

	ResultInHour = seconds / 3600
	fmt.Printf("С начала суток прошло %v часов\n", ResultInHour)
	return
}
