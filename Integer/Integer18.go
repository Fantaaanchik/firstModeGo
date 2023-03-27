package Integer

import (
	"fmt"
	"log"
)

var (
	sauNum     int16
	NeededNum2 int16
)

func Integer18(int16) (NeededNum2 int16) {

	fmt.Printf("Integer18\n Дано целое число, большее 999. Используя одну операцию деления нацело и одну операцию " +
		"взятия остатка от деления, найти цифру, соответствующую разряду тысяч в записи этого числа.\n")
	fmt.Printf("Введите целое число больше 999: \n")
	_, err := fmt.Scan(&sauNum)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных", err.Error())
		return
	}
	if sauNum < 1000 {
		log.Printf("Число %v не является больше чем 999", sauNum)
	}
	NeededNum2 = sauNum % 10000 / 1000
	return
}
