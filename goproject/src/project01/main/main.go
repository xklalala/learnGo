package main
import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}
func AddUper() func(x int) int{
	var n int = 10
	return func(x int) int{
		n = n + x
		return n
	}
}
func main(){

	f1 := AddUper();
	fmt.Println(f1(5))
	fmt.Println(f1(5))
	
	f := makeSuffix(".jp")
	fmt.Println("文件名处理后缀=", f("winter"))
	fmt.Println("文件名处理后缀=", f("winter.jpg"))
}