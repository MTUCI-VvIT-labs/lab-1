package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
//	lab()
	dz()
}

//func lab () {
//	sides := []int{3, 2, 4, 7, 5, 12, 11, 13, 15, 16, 14, 14}
//	sort.Sort(sort.Reverse(sort.IntSlice(sides)))
//
//	smax := 0
//	for i := range sides {
//		for j := i + 1; j < len(sides); j++ {
//			for k := j + 1; k < len(sides); k++ {
//				a := sides[i]
//				b := sides[j]
//				c := sides[k]
//				if a + b > c && a + c > b && b + c > a {
//					p := (a + b + c) / 2
//					s := math.Sqrt(float64(p * (p - a) * (p - b) * (p - c)))
//					if s > float64(smax) {
//						smax = int(s)
//					}
//				}
//			}
//		}
//	}
//
//	fmt.Println("Максимальная площадь треугольника: ", smax)
//
//}


func dz {
	a, b, c := getRatios()     // получаем коэффициенты
	solveDiscriminant(a, b, c) // решаем дискриминант
}

func getRatios() (float64, float64, float64) {
	var a, b, c string // создание переменных для коэффицентов
	for {
		// запросы ввода коэффицентов, в функцию fmt.Scanln передаются указатели на переменные
		fmt.Print("Введите первый коэффициент: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второй коэффициент: ")
		fmt.Scanln(&b)
		fmt.Print("Введите третий коэффициент: ")
		fmt.Scanln(&c)

		// каждый введеный коэффицеинт преобразовывается string -> int и проверяется на ошибку ввода не числового значения
		aInt, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Первый коэффициент введен неверно. Попробуйте ввести коэффициенты еще раз.")
			continue
		}
		bInt, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println("Второй коэффициент введен неверно. Попробуйте ввести коэффициенты еще раз.")
			continue
		}
		cInt, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Третий коэффициент введен неверно. Попробуйте ввести коэффициенты еще раз.")
			continue
		}

		// возрщаем коэффиценты с преобразованием int -> float64 для дальнейших вычислений
		return float64(aInt), float64(bInt), float64(cInt)
	}
}

func solveDiscriminant(a float64, b float64, c float64) {
	d := b*b - 4*a*c
	switch { // проходимся по всем возможным вариантам значений дискриминанта
	case d < 0:
		fmt.Println("Уравнение не имеет корней.")
		return

	case d == 0:
		x := -b / (2 * a)
		fmt.Printf("Уравнение имеет один корень: %f\n", x)

	default:
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("Уравнение имеет два корня: %f и %f\n", x1, x2)
	}
}
