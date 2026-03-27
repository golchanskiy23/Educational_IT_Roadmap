package main

import (
	"fmt"
	"unsafe"
)

func typeCastingExample() {
	fmt.Println("=== Приведение типов ===")

	var x int = 10
	// Интерпретируем байты int как float64 — результат неопределён,
	// но демонстрирует возможность переинтерпретации памяти.
	y := *(*float64)(unsafe.Pointer(&x))
	fmt.Printf("int значение: %d\n", x)
	fmt.Printf("те же байты как float64: %v\n", y)
}

func pointerArithmeticExample() {
	fmt.Println("\n=== Арифметика указателей ===")

	arr := [5]int{10, 20, 30, 40, 50}
	ptr := unsafe.Pointer(&arr[0])

	for i := 0; i < len(arr); i++ {
		// Смещаемся на i элементов от начала массива
		elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(arr[0]))
		val := *(*int)(elemPtr)
		fmt.Printf("arr[%d] = %d\n", i, val)
	}
}

type MyStructABC struct {
	A int
	B float64
	C int32
}

func structFieldAccessExample() {
	fmt.Println("\n=== Доступ к полям структуры ===")

	s := MyStructABC{A: 5, B: 3.14, C: 42}

	// Доступ к полю B через смещение
	bPtr := (*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.B)))
	fmt.Printf("s.B через unsafe: %v\n", *bPtr)

	// Доступ к полю C через смещение
	cPtr := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.C)))
	fmt.Printf("s.C через unsafe: %v\n", *cPtr)

	// Смещения полей
	fmt.Printf("Размер структуры: %d байт\n", unsafe.Sizeof(s))
	fmt.Printf("Смещение A: %d, B: %d, C: %d\n",
		unsafe.Offsetof(s.A), unsafe.Offsetof(s.B), unsafe.Offsetof(s.C))
}

// Union эмулирует C-union: несколько типов разделяют одну область памяти.
type Union struct {
	data [8]byte
}

func (u *Union) SetInt64(val int64) {
	*(*int64)(unsafe.Pointer(&u.data[0])) = val
}

func (u *Union) GetInt64() int64 {
	return *(*int64)(unsafe.Pointer(&u.data[0]))
}

func (u *Union) SetFloat64(val float64) {
	*(*float64)(unsafe.Pointer(&u.data[0])) = val
}

func (u *Union) GetFloat64() float64 {
	return *(*float64)(unsafe.Pointer(&u.data[0]))
}

func (u *Union) GetBytes() [8]byte {
	return u.data
}

func unionExample() {
	fmt.Println("\n=== Union-подобное поведение ===")

	var u Union

	u.SetInt64(12345678)
	fmt.Printf("Записали int64: %d\n", u.GetInt64())
	fmt.Printf("Байты: %v\n", u.GetBytes())

	u.SetFloat64(3.14159)
	fmt.Printf("Записали float64: %f\n", u.GetFloat64())
	fmt.Printf("Байты: %v\n", u.GetBytes())

	// Важно: после SetFloat64 GetInt64 вернёт «мусор» с точки зрения int
	fmt.Printf("Те же байты как int64: %d\n", u.GetInt64())
}

// SliceHeader — ручная версия reflect.SliceHeader для демонстрации.
type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func sliceHeaderExample() {
	fmt.Println("\n=== Манипуляция заголовком слайса ===")

	original := []int{1, 2, 3, 4, 5}

	// Получаем указатель на заголовок слайса
	header := (*SliceHeader)(unsafe.Pointer(&original))
	fmt.Printf("Len: %d, Cap: %d\n", header.Len, header.Cap)

	// Создаём новый слайс, указывающий на тот же массив, но с другой длиной
	// (без новой аллокации)
	newHeader := SliceHeader{
		Data: header.Data,
		Len:  3,
		Cap:  header.Cap,
	}
	newSlice := *(*[]int)(unsafe.Pointer(&newHeader))
	fmt.Printf("Новый слайс (первые 3 элемента): %v\n", newSlice)
	fmt.Printf("Оригинальный слайс не изменился: %v\n", original)
}

// StringHeader — внутреннее представление строки в Go.
type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// stringToBytes конвертирует строку в []byte без копирования.
// ВНИМАНИЕ: результирующий слайс нельзя модифицировать!
func stringToBytes(s string) []byte {
	sh := (*StringHeader)(unsafe.Pointer(&s))
	bh := SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// bytesToString конвертирует []byte в string без копирования.
func bytesToString(b []byte) string {
	bh := (*SliceHeader)(unsafe.Pointer(&b))
	sh := StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func stringConversionExample() {
	fmt.Println("\n=== Конвертация string ↔ []byte (без аллокации) ===")

	s := "Hello, unsafe world!"
	b := stringToBytes(s)
	fmt.Printf("Строка: %s\n", s)
	fmt.Printf("Байты: %v\n", b[:5])

	original := []byte("Go is great")
	str := bytesToString(original)
	fmt.Printf("[]byte -> string: %s\n", str)
}
