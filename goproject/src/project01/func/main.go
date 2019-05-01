package main
import(
	"fmt"
)

func main(){
	num1 := 100
	fmt.Printf("%T, %v, %v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("%T, %v, %v", num2, num2, &num2)

}