package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	b := 0
	c := 1
	println("Введите номер N чисел Фибоначчи:")
	fmt.Scan(&a)
	n, e := strconv.Atoi(a)
	if e != nil {
		println("Ошибка ввода")
		return
	}
	for i := 0; i < n+1; i++ {
		print(b, " ")
		b, c = c, b+c
	}
}