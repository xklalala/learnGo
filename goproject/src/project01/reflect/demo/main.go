package main
import (
	"fmt"
	"reflect"
)

type Moster struct{
	Name	string	`json:"name"`
	Age		int	`json:"age"`
	Score	float32
	Gender	string
}
func (s Moster) Print(){
	fmt.Println("-----------start-----------")
	fmt.Println(s)
	fmt.Println("-----------end-----------")
}
func (s Moster) GetSum(n1 int, n2 int) int{
	return n1 + n2
}

func (s Moster) Set(name string, age int, score float32, gender string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Gender = gender
}

func TestStruct(a interface{}){
	typ := reflect.TypeOf(a)
	val  := reflect.ValueOf(a)
	kd	 := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("except struct")
		return
	}
	num := val.NumField()
	fmt.Println("struct has", num, "field")
	for i := 0; i < num; i++{
		fmt.Printf("Field %d 的值为：%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != ""{
			fmt.Printf("field %d: tag为=%v\n", i, tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("这个类的方法有：%d个\n", numOfMethod)
	val.Method(1).Call(nil)

	//调用结构体 的第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))

	res := val.Method(0).Call(params) //传入的参数是[]reflect.Value
	fmt.Println("res = ", res[0].Int())
}

func main(){
	var a Moster = Moster{
		Name:	"牛魔王",
		Age :	500,
		Score:	98.5,
		Gender:	"男",
	}
	TestStruct(a)
}