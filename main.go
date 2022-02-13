package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x1, y1 := change(getCoordinates())
	fmt.Println("Перавя точка:", x1, y1)

	x2, y2 := change(getCoordinates())
	fmt.Println("Вторая точка:", x2, y2)

	x3, y3 := change(getCoordinates())
	fmt.Println("Третья точка:", x3, y3)
}

func getCoordinates() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100), rand.Intn(100)
}

func change(x, y int) (int, int) {
	x = 2*x + 10
	y = -3*y - 5
	return x, y
}
