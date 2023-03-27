package Integer

import (
	"fmt"
	"log"
)

var (
	SauNum    int16
	NeededNum int16
)

func Integer17(int16) (NeededNum int16) {

	fmt.Printf("Integer17\n Дано целое число, большее 999. Используя одну операцию деления нацело и одну операцию " +
		"взятия остатка от деления, найти цифру, соответствующую разряду сотен в записи этого числа.\n")
	fmt.Printf("Введите целое число больше 999: \n")
	_, err := fmt.Scan(&SauNum)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных", err.Error())
		return
	}
	if SauNum < 1000 {
		log.Printf("Число %v не является больше чем 999", SauNum)
	}
	NeededNum = SauNum % 1000 / 100
	return
}
