package main

import (
	"fmt"

	"github.com/tromper98/tempconv"
)

func main() {
	t := tempconv.Fahrenheit(100)
	c := tempconv.Celsius(32)
	fmt.Printf("Брррр! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Преобрауем Цельсии в Фаренгейты: %v : %v\n", c, tempconv.CToF(c))
	fmt.Printf("Преобразуем Фаренгейты в Кельвины: %v : %v\n", t, tempconv.FToK(t))
}
