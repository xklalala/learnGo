package main
import (
	"fmt"
)

type Monkey struct{
	Name string
}

func (this *Monkey) climbing(){
	fmt.Println(this.Name, "会爬树")
}

type LittleMonkey struct{
	Monkey
}

type BirdAble interface{
	Flying()
}

type FishAble interface{
	Swiming()
}
func (this *LittleMonkey) Flying(){
	fmt.Println(this.Name, "通过学习会飞翔")
}
func (this *LittleMonkey) Swiming(){
	fmt.Println(this.Name, "通过学习会游泳")
}

func main(){
	monkey := LittleMonkey{
		Monkey{
			Name : "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swiming()
}