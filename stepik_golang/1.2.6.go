// echo "1 1 4 5" | go run 1.2.6.go

package main

import (
	"fmt"
	"math"
)

func main() {
	// объявите переменные x1, y1, x2, y2 типа float64
	var x1, y1, x2, y2 float64

	// считываем числа из os.Stdin
	// гарантируется, что значения корректные
	// не меняйте этот блок
	fmt.Scan(&x1, &y1, &x2, &y2)

	// рассчитайте d по формуле эвклидова расстояния
	// используйте math.Pow(x, 2) для возведения в квардрат
	// используйте math.Sqrt(x) для извлечения корня
	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	// выводим результат в os.Stdout
	fmt.Println(d)
}
