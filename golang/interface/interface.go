package interface

import(
	"os"
	"strconv"
)

// определение интерфейса
type ReadCloser interface{
	Read([] byte) (n int, err os.Error)
	Close()
}

// принимает любой объект структуры, реализующей интерфейс
func ReadAndClose(r ReadCloser, buf []byte) (n int, err os.Error) {
	for len(buf) > 0 && err == nil{
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}

type Stringer interface{
	String() string
}

// Проверка наличия дополнительного метода у значения интерфейса
// у объектов типа interface{} нет гарантий на наличие методов
func toString(any interface{}) string{
	if v, ok := any.(Stringer); ok{
		return v.String()
	}
	switch v := any.(type){
	case int:
		return strconv.Itoa(v)
	case float:
		return strconv.Ftoa(v, 'g', -1)
	}
	return "no correct method"
}

type Binary uint64

func (i Binary) String() string{
	return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64{
	return uint64(i)
}

func main(){
	b := Binary(200) 
	s := Stringer(b) 
 	fmt.Println(s.String())
}