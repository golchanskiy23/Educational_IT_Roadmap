package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "os"
	"strconv"
)

// структурне тэги
type Person struct{
	Age int			`json:"-"` // игнор при десериализации
	Name string 	`json:"name,omitempty"` // пропускается, если значение пусто
	Height float32  `json:"height"` // переименовали
	Weight float32	`json:",omitempty"` // оставляем название, если пусто - пропускаем
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"user"`
	Age      int    `json:"-"` // Полностью игнорируется при сериализации
}

func print_err(err error){
	fmt.Println(fmt.Errorf("Error is: %v", err))
}

func serialization_deserialization(){
	max_person := &Person{
		Age: 22,
		Name : "Max",
		Height: 173,
		Weight: 65,
	}

	// сериализация(Marshal) -> []byte, десериализация(Unmarshal) -> struct
	data, err := json.Marshal(max_person)
	if err != nil{
		print_err(err)
	}
	fmt.Println(data)

	/*
	ch := make(chan struct{})
	_, err := json.Marshal(ch) // returns error *json.UnsupportedTypeError

	compl := complex(10, 11)
	_, err = json.Marshal(compl) // returns error *json.UnsupportedTypeError

	fn := func() {}
	_, err = json.Marshal(fn) // returns error *json.UnsupportedTypeError
	*/

	// json сериализует только публичные поля
	myVal := Person{}
	byte_arr := `{"name":"Max", "height":173.0}`
	err = json.Unmarshal([]byte(byte_arr), &myVal)
	if err != nil{
		print_err(err)
	}
	fmt.Print(myVal)

	user := User{}
	arr := []byte(`{"id":"some-id","user":"admin"}`)
	err = json.Unmarshal(arr, &user)
	if err != nil{
		print_err(err)
	}
	fmt.Println(user)
}

func encoder_decoder(){
	// Encode и Decoder работают с потоками данных и обрабатывают их на лету
	// также с большими данными
	type TmpRequest struct{
		UserId int		`json:"userId"`
		Id int			`json:"id"`
		Title string	`json:"title"`
		Completed bool	`json:"completed"`
	}
	// возвращает json
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil{
		print_err(err)
	}
	
	var req TmpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		print_err(err)
	}
	fmt.Printf("%v\n", req)

	req1 := TmpRequest{}
	
	r, err = http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil{
		print_err(err)
	}
	
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err = json.Unmarshal(body, &req1); err != nil{
		print_err(err)
	}
	fmt.Println(req1)
}

func format_json(){
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	b, err := json.MarshalIndent(data, "<префикс>", "<отступ>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

/*
type Marshaler interface {
	MarshalJSON([]byte) error
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
*/

const(
	_ = 1 << (10*iota)
	KiB
	MiB
	GiB
	TiB
)

var memorySizes = []int{TiB, GiB, MiB, KiB}
var sizeSuffixes = []string{"Tb", "Gb", "Mb", "Kb"}

type Memory string

type PC struct{
	CPU string
	OperatingSystem string
	Memory Memory
	Storage Memory
}

func (m *Memory) UnmarshalJSON(b []byte) error{
	size, err := strconv.Atoi(string(b))
	if err != nil{
		print_err(err)
		return err
	}

	for i, d := range memorySizes{
		if size > d{
			*m = Memory(fmt.Sprintf("%d %s\n", size / d, sizeSuffixes[i]))
			return nil
		}
	}
	*m = Memory(fmt.Sprintf("%d %s", size, "b"))
	return nil
}

// кастомизация процесса преобразования в структуры в json
func serialization_deserialization_rules(){
	data := []byte(`{
		"cpu": "Intel Core i5",
		"operatingSystem": "Windows 11",
		"memory": 17179869184,
		"storage": 274877906944
	}`)
	pc := PC{}
	if err := json.Unmarshal(data, &pc); err != nil{
		log.Fatalf("Error is: %v",err)
	}
	fmt.Printf("%+v", pc)
}

func main(){
	// serialization_deserialization()
	// encoder_decoder()
	// format_json()
	// serialization_deserialization_rules()
}
