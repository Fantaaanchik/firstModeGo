package Integer

import (
	"fmt"
	"log"
)

var (
	hour     int32
	ResInSec int32
)

func IntegerMy(int32) (ResInSec int32) {

	fmt.Printf("С начала суток прошло N часов (N — целое). Найти количество полных секунд, прошедших с начала суток. \n")
	fmt.Printf("Введите число: \n")
	_, err := fmt.Scan(&hour)
	if err != nil {
		log.Println("Вы ввели не правильные данные, нужно ввести число \n", err.Error())
		return
	}

	ResInSec = hour * 3600
	fmt.Printf("С начала суток прошло %v секунд \n", ResInSec)
	return
}
