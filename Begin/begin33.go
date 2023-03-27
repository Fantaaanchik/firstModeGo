package Begin

import (
	"fmt"
	"log"
)

var (
	X     float32
	A     float32
	Ykg   float32
	OneKg float32
)

func Beginner33(float32, float32) (OneKg, Ykg float32) {
	fmt.Println("Begin33\n Известно, что X кг конфет стоит A рублей. Определить, сколько стоит 1 кг и Y кг этих же конфет")
	fmt.Println("Введите 2 числа, X и A: ")
	_, err := fmt.Scan(&X, &A)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных\n", err.Error())
		return
	}
	fmt.Println("Введите сколько кг конфет нужно посчитать: ")
	_, er := fmt.Scan(&Ykg)
	if er != nil {
		log.Println("\"Вы ввели не правильный формат данных\n", er.Error())
		return
	}

	OneKg = A / X
	Ykg = Ykg * OneKg
	fmt.Printf("Один килограмм конфет стоит %v рублей\n", OneKg)
	fmt.Printf("Y килограмм конфет стоит %v рублей\n", Ykg)
	return
}
