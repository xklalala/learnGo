package main
import(
	"fmt"
)

func learn1(){
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	fmt.Println(intArr)

	//定义一个切片
	slice := intArr[1:3]   //表示slice引用intArr这个数组下标为1到2的元素    左闭右开 [)
	fmt.Println(slice)
	fmt.Println("slice的元素个数是：", len(slice))
	fmt.Println("slice的元素容量是：", cap(slice))
}

func learn2(){
	arr := [...]int{11, 22, 33, 44, 55}
	slice := arr[0:3]
	fmt.Println("arr的首地址是：", &arr[0])
	fmt.Println("slice的首地址是：",&slice[0])
	slice[2] = 88
	fmt.Println(arr)
}

func learn3(){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()
	//切片使用的三种方式
	//1.定义一个切片，然后让切片去引用一个已经创建好的数组，比如上面的
	//2.通过make来创建切片   语法：var slice []type = make([]type, len, [cap])
	//3.定义一个切片，直接指定具体数组，原理make类似
	var slice1 []float64 = make([]float64, 5, 10)
	slice1[1] = 10.2
	slice1[0] = 20
	fmt.Println(slice1)
	//通过make方式创建的切片，可以指定大小和容量。如果没有给切片各个元素赋值，那么就会使用默认值 

	var slice2 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice2)
	slice2[3] = "24234"//不能越界
	fmt.Println("size : ", len(slice2))
	fmt.Println("cap : ", cap(slice2))
}

func learn4(){
	//切片遍历

	var slice []string = []string{"tom", "jack", "jerry", "hook"}


	//01
	for i := 0; i < len(slice); i++{
		fmt.Printf("%v=>%v ", i, slice[i])
	}
	fmt.Println()
	slice2 := append(slice, "lalala", "xk")
	fmt.Println(slice2, slice)

	slice3 := append(slice, slice2...)
	fmt.Println(slice3)
	for k, v := range slice{
		fmt.Printf("%v=>%v ", k, v)
	}
}

//字符串与切片
func learn5(){
	str := "abcdefg"

	slice1 := str[2:]
	fmt.Println(slice1)

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	arr2 := []rune(str) //按照字符来处理
	arr2[0] = '被'
	str = string(arr2)
	fmt.Println(str)
}

func main(){
	learn5()
}