package main
import (
	"fmt"
)

//channel线程安全，多个协程操作一个管道时不会发生资源竞争问题

func learn1(){
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan)

	//向管道中存
	intChan<- 10
	intChan<- 20
	intChan<- 30

	//取
	num := <-intChan
	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock 死锁

	fmt.Println(num)

}


func main(){
	learn1()
}