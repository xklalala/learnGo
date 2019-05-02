package main
import(
	"fmt"
)

type Student struct{
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (student *Student) Say() string{
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]", student.Name, student.Gender, student.Age, student.Id, student.Score)
	return infoStr
}

func learn1(){
	var stu = Student{
		Name:   "zhangsan",
		Gender: "man",
		Age:    15,
		Id:     1000,
		Score:  99.8,
	}
	fmt.Println(stu.Say())

}

type Box struct{
	Len    float64
	Width  float64
	Height float64
}
func (box *Box) getVolumn() float64{
	return box.Len * box.Width * box.Height
}

func learn2(){
	box := Box{3.0, 4.0, 5.0}
	fmt.Println(box.getVolumn())
}

type Visitor struct{
	Name string
	Age  int
}

func (visitor *Visitor) showPrice(){
	if visitor.Age >= 90 || visitor.Age<=8{
		fmt.Println("考虑到安全，就不要玩了")
	}
	if visitor.Age > 18{
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元\n", visitor.Name, visitor.Age)
	}else{
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费\n", visitor.Name, visitor.Age)
	}
}

func learn3(){
	var v Visitor
	for{
		fmt.Printf("请输入你的名字：")
		fmt.Scanln(&v.Name)
		if v.Name == "n"{
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}

func main(){
	learn3()
}