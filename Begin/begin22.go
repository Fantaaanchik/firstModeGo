package Begin

import (
	"fmt"
	"log"
)

var (
	Res     int8
	firstN  int8
	secondN int8
)

func Beginner22(c, i int8) (A, B int8) {

	fmt.Println("Begin22\n Поменять местами содержимое переменных A и B и вывести новые значения A и B.")
	fmt.Printf("Введите А и В: \n")
	_, err := fmt.Scan(&A, &B)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных:", err.Error())
		return
	}
	Res = firstN
	firstN = secondN
	secondN = Res
	return
}
