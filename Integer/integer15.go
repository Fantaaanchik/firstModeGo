package Integer

import (
	"fmt"
	"log"
)

var (
	Chislo3znach int
	Result       int
	hundred      int
	ten          int
	one          int
)

func Integer15(int) (Result int) {
	fmt.Println("Integer15\n Дано трехзначное число. Вывести число, полученное при перестановке цифр сотен и десятков " +
		"исходного числа (например, 123 перейдет в 213).")
	fmt.Printf("Введите трехзначное число: \n")
	_, err := fmt.Scan(&Chislo3znach)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных ", err.Error())
		return
	}
	if Chislo3znach > 999 || Chislo3znach < 100 {
		fmt.Printf("Вы ввели не правильное число, нужно трехзначное\n")
	}
	hundred = Chislo3znach / 100
	ten = Chislo3znach / 10 % 10
	one = Chislo3znach % 10 % 10
	Result = ten*100 + hundred*10 + one

	return
}
