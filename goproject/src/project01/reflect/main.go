package main
import (
	"fmt"
	"reflect"
)
// 反射可以在运行的时候动态获取变量的各种信息，比如变量的类型，类别
// 如果是结构体变量， 还可以获取到结构体本身的信息（包括结构体的字段、方法）
// 通过反射可以修改变量的值，可以调用关联的方法
// 使用反射需要import "reflect"



//对基本数据类型、interface{}、reflect.Value进行反射的基本操作

func reflectTest01(b interface{}){
	//通过反射获取传入变量的type kind值
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	rValue := reflect.ValueOf(b)
	fmt.Printf("%T\n", rValue)

	iv := rValue.Interface()
	num := iv.(int)
	fmt.Printf("%v, %T", num, num)

}



func learn1(){
	reflectTest01(123)
}

//演示反射结构体
func reflectTest02(b interface{}){
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	rValue := reflect.ValueOf(b)

	iv := rValue.Interface()
	fmt.Printf("%T, %v\n", iv, iv)

	switch iv.(type) {
		case dog:
			d := iv.(dog)
			fmt.Printf("%T, %v\n", d, d)
		case student: 
			d := iv.(student)
			fmt.Printf("%T, %v\n", d, d)
		default:return
	}
	
}
type dog struct{
	Name string
	Age	 int
}
type student struct{
	Name string
	Age	 int
}
func learn2(){
	dog := dog{"小黄", 5}
	reflectTest02(dog)
	student := student{"啦啦啦", 12}
	reflectTest02(student)
}

 
func main(){
	// learn1()
	learn2()
}