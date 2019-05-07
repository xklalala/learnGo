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


type Cat struct{
	Name string
	Age	 int
}

func learn2(){
	//定义一个可以存放任意数据类型的管道  3个数据
	allChan := make(chan interface{}, 3)
	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat

	//要获得第三个数据，要把前面两个推掉。fifo
	<-allChan
	<-allChan
	newCat := <-allChan
	// fmt.Println(newCat.Name)//这里的newCat是接口类型的
	a := newCat.(Cat)//类型断言
	fmt.Println(a.Name)
}

func learn3(){
	intchan := make(chan int, 3)
	intchan<- 100
	intchan<- 200
	close(intchan)//关闭管道
	//这时不能够再写入数据到channel
	// intchan<- 300
	n1 := <-intchan//但是可以读数据
	fmt.Println(n1)
}

//遍历取数据
//channel支持for-range的方式进行遍历
//在遍历时，如果channel没有关闭，则会出现deadlock的错误
//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成之后，就会退出
func learn4(){
	//遍历需要先关闭管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++{
		intChan2<- i * 2
	}
	close(intChan2)

	for v := range intChan2{
		fmt.Println(v)
	}
}
 
func main(){
	// learn1()
	// learn2()
	// learn3()
	learn4()
}