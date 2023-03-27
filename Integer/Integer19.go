package Integer

import (
	"fmt"
	"log"
)

var (
	second         int16
	ResultInMinute int16
)

func Integer19(int16) (ResultInMinute int16) {

	fmt.Printf("Integer19\n С начала суток прошло N секунд (N — целое). Найти количество полных минут, прошедших " +
		"с начала суток\n")
	fmt.Printf("Введите челое число, количество секунд: \n")
	_, err := fmt.Scan(&second)
	if err != nil {
		log.Println("Вы ввели не правильные данные \n", err.Error())
		return
	}

	ResultInMinute = second / 60
	fmt.Printf("С начала суток прошло %v минут\n", ResultInMinute)
	return
}
