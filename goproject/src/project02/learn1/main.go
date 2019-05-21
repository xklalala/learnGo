package main
import (
	"fmt"
)

func main(){
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\t%T\n", arr1, arr2)
	b := &arr1
	fmt.Println(b[0], arr1[0])
}