package main

import(
	"fmt"
)

type Foo struct{
	Val int
}

type Bar struct{
	Val int
}

func fooToBar(foo Foo) Bar{
	return Bar{
		Val: foo.Val,
	}
}

// проблема в постоянном пересоздании массива
// так как заканчивается ёмкость -> доп нагрузка на GC
func convert(foos []Foo) []Bar{
	bars := make([]Bar, 0)
	for _, bar := range foos{
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

// создание среза с заданной ёмкостью
func convert1(foos []Foo) []Bar{
	bars := make([]Bar, 0, len(foos))
	for _, bar := range foos{
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

// создание среза с заданной длиной, проблема 
// в интерфейсе присваиваний
func convert2(foos []Foo) []Bar {
    n := len(foos)
    bars := make([]Bar, n) // Создаем срез с заданной длиной
    for i, foo := range foos {
        bars[i] = fooToBar(foo) // Присваиваем значение по индексу
    }
    return bars
}

func something(foo Foo) bool{
	if foo.Val%2==0{
		return true
	}
	return false
}

// при неизвестной длине среза - создаём nil срез
func convertConditionally(foos []Foo) []Bar{
	var bars []Bar // nil срез
	for _, foo := range foos {
        if something(foo) { // Добавляем только при выполнении условия
            bars = append(bars, fooToBar(foo))
        }
    }
    return bars
}

func main(){
	// init slice(len=3, cap=6)
	s := make([]int , 3, 6)
	fmt.Println(s) // [0 0 0]
	// обновление элементов
	s[1] = 1
	fmt.Println(s) // [0,1,0]
	s = append(s, 2)
	fmt.Println(s) // [0 1 0 2]
	s = append(s, 3, 4, 5)
	fmt.Println(s) // [0 1 0 2 3 4 5]

	s1 := make([]int, 3, 6) // [0 0 0]
	s2 := s1[1:3]           // Новый срез с длиной 2 и емкостью 5
	s1[1] = 1
	fmt.Println(s1, s2) // [0 1 0] [1 0]

	s2 = append(s2, 2)
	fmt.Println(s1, s2) // [0 1 0] [1 0 2]

	s2 = append(s2, 3, 4, 5)
	fmt.Println(s1, s2) // [0 1 0] [1 0 2 3 4 5]

	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	// длина копирования - минимум среди len(dst), len(src)
	fmt.Println(dst) // Выведет []

}

