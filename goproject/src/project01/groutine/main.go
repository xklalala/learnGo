package main
import (
	"fmt"
	"time"
)
//go里面叫协程
//特点：有独立的栈空间、共享程序堆空间、调度由用户控制、协程是轻量级的线程

func learn1(){
	for i := 1; i <= 10; i++{
		fmt.Println("hello go", i)
		time.Sleep(time.Second)
	}
}



func main(){
	go learn1()  //开启一个协程
	for j := 'A'; j <= 'A'+10 ; j++{
		fmt.Printf("%c\n", j)
		time.Sleep(time.Second)
	}
}