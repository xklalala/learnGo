package main
import(
	"fmt"
)

func learn1(){
	fmt.Println()
	var arr [2]int;
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])
}

//二维数组
func learn2(){
	var arr [4][4]int
	arr[0][0] = 1
	arr[1][1] = 2
	arr[2][2] = 3
	arr[3][3] = 4
	for _,v := range arr{
		fmt.Println(v)
	}
}

func main(){
	learn2()

}