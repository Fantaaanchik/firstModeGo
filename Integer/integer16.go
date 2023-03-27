package Integer

import (
	"fmt"
	"log"
)

var (
	сhislo3znach int
	Resulted     int
	hundreded    int
	tened        int
	oned         int
)

func Integer16(int) (Resulted int) {
	fmt.Println("Integer16\n Дано трехзначное число. Вывести число, полученное при перестановке цифр десятков и единиц исходного числа" +
		" (например, 123 перейдет в 132")
	fmt.Printf("Введите трехзначное число: \n")
	_, err := fmt.Scan(&сhislo3znach)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных ", err.Error())
		return
	}
	if сhislo3znach > 999 || сhislo3znach < 100 {
		fmt.Printf("Вы ввели не правильное число, нужно трехзначное\n")

	}
	hundreded = сhislo3znach / 100
	tened = сhislo3znach / 10 % 10
	oned = сhislo3znach % 10 % 10
	Resulted = hundreded*100 + oned*10 + tened

	return
}
