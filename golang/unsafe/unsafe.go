package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func raw_ptr_transformation(){
	var x int = 42
	ptr := unsafe.Pointer(&x)
	fmt.Println(ptr)
	address := uintptr(ptr)+4
	// скорее всего промажем
	ptr2 := unsafe.Pointer(address)
	fmt.Println(ptr2)
}

func size_of(){
	var x int64
	fmt.Println(unsafe.Sizeof(x)) // 8 байт
}

type MyStruct struct {
	A int8
	B int64
}

func allign_of(){
	var s MyStruct
	fmt.Println(unsafe.Sizeof(s)) // 16 байт = (1+9)+7(паддинг) байт
	fmt.Printf("Выравнивание поля A: %d\n", unsafe.Alignof(s.A))
	fmt.Printf("Выравнивание поля B: %d\n", unsafe.Alignof(s.B))
}

func int_to_float(){
	var x int = 42

	// Преобразуем указатель на x в unsafe.Pointer
	ptr := unsafe.Pointer(&x)

	// А затем в указатель на float32
	floatPtr := (*float32)(ptr)

	// Внимание: здесь вас ждёт undefined behavior!
	fmt.Println(*floatPtr) // 5.9e-44 - представление не соответствует int
}

type SecretSrtuct struct{
	private_attribute string
}

func private_access(){
	secret := SecretSrtuct{
		private_attribute: "it's a secret",
	}

	// получение доступа через reflect
	field := reflect.ValueOf(&secret).Elem().FieldByName("private_attribute")
	ptr := unsafe.Pointer(field.UnsafeAddr())
	realptr := (*string)(ptr)

	*realptr = "it's not a secret"
	fmt.Println(secret.private_attribute)
}

// при этом изменения в b заторнут s
func trasform_copy_optimization(){
	b := []byte("Hello")

	// without copying
	s := *(*string)(unsafe.Pointer(&b))
	fmt.Println(s)
}

type Unaligned struct {
	A int32 // Поле A занимает 4 байта
	B int16 // Поле B занимает 2 байта, но для выравнивания Go добавляет "пустое место"
}

func alighning(){
	u := Unaligned{A:10, B:20}
	offset := unsafe.Offsetof(u.B) // 4 байта
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&u))+offset)
	*(*int16)(ptr) = 42
	fmt.Println(u.B)
}

func main() {
	// raw_ptr_transformation()
	// size_of()
	// allign_of()
	// int_to_float()
	// private_access()
	// trasform_copy_optimization()
	// alighning()

	typeCastingExample()
	pointerArithmeticExample()
	structFieldAccessExample()
	unionExample()
	sliceHeaderExample()
	stringConversionExample()
}