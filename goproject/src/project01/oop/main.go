package main
import(
	"fmt"
	"encoding/json"
)
//go不是纯粹的面向对象编程，面向对象是借助结构体实现的
//go仍然有面向对象的继承、封装、多态的特性

type Cat struct{
	Name  string
	Age   int
	Color string
}

type Person struct{
	Name string
	Age int
}
func learn2(){
	p2 := Person{"lalala", 20}
	fmt.Println(p2)

	var p3 *Person = new(Person)
	(*p3).Name = "xk"
	(*p3).Age = 18
	//上面的写法等价于
	p3.Name = "xk"
	p3.Age = 20
	fmt.Printf("%v, %p \n", *p3, p3)

	var person *Person = &Person{}
	(*person).Name = "stop"
	(*person).Age = 11
	//等价于
	person.Name = "abc"
	person.Age = 2

	var p1 *Person = &Person{"marry", 60}
	fmt.Println(*p1)
}
func learn1(){
	var cat Cat
	cat.Name = "小白"
	cat.Age = 2
	cat.Color = "黑色"

	cat1 := Cat{"小花", 3, "白色"}
	fmt.Println(cat1)
	fmt.Println(cat)
}

type Monster struct{
	Name  string `json:"name"`
	Age   int	 `json:"age"`
	Skill string `json:"skill"`
}

func learn3(){
	monster := Monster{"李白", 34, "青莲剑歌"}
	jsonStr, err := json.Marshal(monster)
	if err != nil{
		fmt.Println("josn 处理错误", err)
	}else{
		fmt.Println(string(jsonStr))
	}

}

type Animal struct{
	Name string
}
//给person绑定一个方法
func (a Animal) test(){
	a.Name = "猫"
	fmt.Println("test（）name = ", a.Name)
}

func (a Animal) cac(n int){
	res := 0
	for i := 0; i <= n; i++{
		res += i
	}
	fmt.Println(res)
}

func learn4(){
	var a Animal
	a.Name = "狗"
	
	a.test()
	fmt.Println(a)
	a.cac(100)
}

func main(){
	learn4()
}