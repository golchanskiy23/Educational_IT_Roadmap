# JSON

## FAQ по JSON
### Не знаем схему json
Используем map

```go
req := map[string]interface{}{}
if err := json.Decoder(r.Body).Decode(&req); err != nil{
    ...
}

// после итерации по нему можно обработать логику запроса используя рефлексию
for k, v := range req {
	refVal := reflect.TypeOf(v)
	fmt.Printf("ключ '%s' содержит значение типа %s\n", k, refVal)
}

```

### Не вижу свои поля в JSON после сериализации
- Поле не публичное
- Использован тэг `json:"-"`

### Пропускать ошибку в Marshal(), Unmarshal()
Нет
