package main
import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name		string	`json:"name"`
	Birthday	string	`json:"birthday"`
	Sal			float64
	Skill		string
}
//序列号结构体
func learn1(){
	moster := Monster{
		Name 	 : "牛魔王",
		Birthday : "2019-5-4",
		Sal		 : 8000.0,
		Skill	 : "山崩地裂",
	}
	jsonStr, err := json.Marshal(&moster)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
}


//将map进行序列化
func testMap(){
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"]	 = "红孩儿"
	a["age"] 	 = 30
	a["adress"]  = "洪崖洞"
	jsonStr, err := json.Marshal(a)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))

}
//对切片进行序列化
func testSlice(){
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"]	 = "红孩儿"
	m1["age"] 	 = 300
	m1["adress"]  = "洪崖洞"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"]	 = "jack"
	m2["age"] 	 = 25
	m2["adress"]  = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	var m3 map[string]interface{}
	m3 = make(map[string]interface{})
	m3["name"]	 = "barry"
	m3["age"] 	 = 18
	m3["adress"]  = "USA"
	slice = append(slice, m3)

	jsonStr, err := json.Marshal(slice)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))

}

//序列化普通的数据类型, 意义不大
func testFloat64(){
	var num1 float64 = 23456.7
	jsonStr, err := json.Marshal(num1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
}

//反序列化成结构体
func learn2(){
	str := "{\"name\":\"牛魔王\",\"birthday\":\"2019-5-4\",\"Sal\":8000,\"Skill\":\"山崩地裂\"}"
	var moster Monster

	err := json.Unmarshal([]byte(str), &moster)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(moster)
}

//反序列化成map
func learn3(){
	str := "{\"adress\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(a)
}

//反序列化为切片，注意每个反序列化都需要格式一致
func learn4(){
	str := "[{\"adress\":\"洪崖洞\",\"age\":300,\"name\":\"红孩儿\"},{\"adress\":[\"墨西哥\",\"夏威夷\"],\"age\":25,\"name\":\"jack\"},{\"adress\":\"USA\",\"age\":18,\"name\":\"barry\"}]"
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(a)
}




func main(){
	// learn2()
	// testFloat64()
	// testMap()
	// learn3()
	testSlice()
	learn4()
	}
}