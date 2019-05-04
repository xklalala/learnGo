package main
import "fmt"
import "sort"
import "math/rand"

type Usb interface{
	Start()
	Stop()
}

type Phone struct{
	Name string
}
func (phone Phone) Start(){  
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop(){
	fmt.Println("手机停止工作")
}
func (phone Phone) Call(){
	fmt.Println("手机在打电话")
}

type Camera struct{
	Name string
}
func (camera Camera) Start(){
	fmt.Println("相机开始工作")
}
func (camera Camera) Stop(){
	fmt.Println("相机停止工作")
}

type Computer struct{

}

func (computer Computer) Working(usb Usb){
	usb.Start()
	//类型断言
	// if phone, ok := usb.(Phone); ok == true{
	// 	phone.Call()
	// }
	phone, ok := usb.(Phone)
	if ok{
		phone.Call()
	}
	usb.Stop()
}

func learn1(){
	computer := Computer{}
	// phone    := Phone{"huawei"}
	// camera   := Camera{"soni"}
	// computer.Working(phone)
	// computer.Working(camera)

	var usbArr [3]Usb
	usbArr[0] = Phone{"华为"}
	usbArr[1] = Phone{"iPhone"}
	usbArr[2] = Camera{"索尼"}
	for _, v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}
}

func learn(){
	var intSlice = []int{0, -1, 10, 50, 100}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}

type Hero struct{
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int{
	return len(hs)
}
//按年龄排序
func (hs HeroSlice) Less(i, j int) bool{
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int){
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}

func learn2(){
	var heroes HeroSlice
	for i := 0; i < 10; i++{
		hero := Hero{
			Name : fmt.Sprintf("英雄-%d", rand.Intn(100)),
			Age  : rand.Intn(500),
		}
		heroes = append(heroes ,hero)
	}
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
}


//判断参数的类型
func TypeJudge(item ...interface{}){
	for i, x := range item{
		switch x.(type){
			case bool:
				fmt.Printf("参数  %d 的类型是bool 值是%v\n", i, x)
			case float64:
				fmt.Printf("参数  %d 的类型是float 值是%v\n", i, x)
			case int, int64:
				fmt.Printf("参数  %d 的类型是int 值是%v\n", i, x)
			case nil:
				fmt.Printf("参数  %d 的类型是nil 值是%v\n", i, x)
			case string:
				fmt.Printf("参数  %d 的类型是string 值是%v\n", i, x)
			case Hero:
				fmt.Printf("参数  %d 的类型是Hero 值是%v\n", i, x)
			case *Hero:
				fmt.Printf("参数  %d 的类型是Hero 值是%v\n", i, x)
			default:
				fmt.Printf("参数  %d 的类型未知 值是%v\n", i,x)
		}
	}
}


func main(){
	hero := Hero{"xk", 25}
	TypeJudge(1, 2.5, "lalala", hero, &hero)

}