package main
import(
	"fmt"
	"../model"
)

func learn1(){
	// var stu = model.Student{
	// 	Name : "tom",
	// 	Score : 78.5,
	// }
	var stu = model.NewStudent("tom", 99.5)
	fmt.Println(*stu)
}

func main(){
	learn1()
}