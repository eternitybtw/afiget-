package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	x := 0
	y := 1
	fmt.Println("Введите количество чисел Фибоначчи, которое вы хотите увидеть")
	fmt.Scan(&a)
	n, e := strconv.Atoi(a)
	if e != nil {
		println("Ошибка ввода")
		return
	}
	for i := 0; i < n+1; i++ {
		print(x, " ")
		x, y = y, x+y
	}
}
