package main
import(
	"fmt"
)

func learn1(){
	fmt.Println()

	var a map[string]string
	//在使用map前要给先make， make是给map分配数据空间
	a = make(map[string]string, 10)
	a["name"] = "松江"
	a["age"] = "56"
	//go里面的map是无序的

	fmt.Println(a)
}

func learn2(){
	//map使用的三种方式

	//1.
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["k1"] = "v1"
	map1["k2"] = "v2"
	map1["k3"] = "v3"
	//2.
	cities := make(map[string]string)
	cities["c1"] = "北京"
	cities["c2"] = "tianjing"
	cities["c3"] = "shanghai"
	cities["c4"] = "guangzhou"

	//3.
	river := map[string]string{
		"r1": "长江",
		"r2": "黄河",
		"r3": "松花江",
		"r4": "扬子江",
	}
	river["r5"] = "澜沧江"
	fmt.Println(river)

	//删除
	delete(river, "r1")
	river = make(map[string]string)

	//查找
	v, status := map1["k1"]
	if status{
		fmt.Println("值为：",v)
	}else{
		fmt.Println("没有呀")
	}
}

func learn3(){
	studentMap := make(map[string]map[string]string);

	studentMap["stu1"] = make(map[string]string, 3)
	studentMap["stu1"]["name"] = "tom"
	studentMap["stu1"]["sex"] = "男"
	studentMap["stu1"]["adress"] = "上海"

	studentMap["stu2"] = make(map[string]string, 3)
	studentMap["stu2"]["name"] = "tim"
	studentMap["stu2"]["sex"] = "女"
	studentMap["stu2"]["adress"] = "北京"

	fmt.Println(studentMap)
}

func learn4(){
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil{
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil{
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "yutu"
		monsters[1]["age"] = "500"
	}

	newMonster := map[string]string{
		"name":"新妖怪",
		"age":"200",
	}
	monsters = append(monsters, newMonster)
}

type stu struct{
	Name string
	Age  int
	Adress string
}

func learn5(){
	students := make(map[string]stu, 10)

	stu1 := stu{"tom", 17, "北京"}
	stu2 := stu{"tim", 15, "上海"}
	stu3 := stu{"jerrym", 18, "南昌"}

	students["sut1"] = stu1
	students["sut2"] = stu2
	students["sut3"] = stu3
	
	fmt.Println(students)

}


func learn6_a(users map[string]map[string]string, name string){
	if users[name] != nil{
		users[name]["password"] = "98888"
	}else{
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] ="456789"
		users[name]["nickname"] = "昵称·~"
	}
}

func learn6(){
	users := make(map[string]map[string]string, 10)
	learn6_a(users, "tom")
	learn6_a(users, "tom")
	learn6_a(users, "marry")
	fmt.Println(users)
}



func main(){
	learn6()
}