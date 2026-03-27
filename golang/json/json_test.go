package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

var j = []byte(`{"user":"Johny Bravo","items":[{"id":"4983264583302173928","qty": 5}]}`)
var createRequest = CreateOrderRequest{
	User: "Johny Bravo",
	Items: []OrderItem{
		{ID: "4983264583302173928", Qty: 5},
	},
}
var err error
var body []byte

// OrderItem представляет элемент заказа.
type OrderItem struct {
	ID  string `json:"id"` // Идентификатор элемента
	Qty int    `json:"qty"` // Количество
}

// CreateOrderRequest описывает запрос на создание заказа.
type CreateOrderRequest struct {
	User  string      `json:"user"` // Пользователь, совершающий заказ
	Items []OrderItem `json:"items"` // Список элементов заказа
}

// BenchmarkJsonUnmarshal измеряет производительность функции json.Unmarshal.
func BenchmarkJsonUnmarshal(b *testing.B) {
	b.ReportAllocs() // Отчет о выделениях памяти
	req := CreateOrderRequest{}
	b.ResetTimer() // Сброс таймера для чистого измерения

	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(j, &req) // Десериализация JSON в структуру
	}
}

// BenchmarkJsonDecoder измеряет производительность использования json.Decoder.
func BenchmarkJsonDecoder(b *testing.B) {
	b.ReportAllocs()
	req := CreateOrderRequest{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer() // Остановка таймера на время подготовки
		buff := bytes.NewBuffer(j) // Создание буфера для чтения
		b.StartTimer() // Возобновление измерения времени

		decoder := json.NewDecoder(buff) // Создание декодера
		err = decoder.Decode(&req) // Декодирование JSON в структуру
	}
}

// BenchmarkJsonMarshal измеряет производительность функции json.Marshal.
func BenchmarkJsonMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		body, err = json.Marshal(createRequest) // Сериализация структуры в JSON
	}
}

// BenchmarkJsonEncoder измеряет производительность использования json.Encoder.
func BenchmarkJsonEncoder(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(io.Discard) // Создание энкодера, вывод в /dev/null
		err = encoder.Encode(createRequest) // Кодирование структуры в JSON
	}
}

