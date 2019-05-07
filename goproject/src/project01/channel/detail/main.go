package main
import (
	"fmt"
	"time"
)

//管道可以声明为只读或者只写





func learn1(){
	// var chan1 chan int  //双向
	
	// var chan2 chan<- int //只写
	// chan2 = make(chan int, 3)
	// chan2<- 20
	// num := <-chan

	// var chan3 <-chan int//只读
	// num2 := <-chan3


	fmt.Println()
}


func learn2(){
	intChan := make(chan int, 10)
	for i := 0; i<10; i++{
		intChan<- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++{
		stringChan<- "hello" + fmt.Sprintf("%d", i)
	}

	//在实际开发中，有可能我们不好确定什么时候关闭管道，可以使用select方式解决
	for{
		select{
			//如果这里intchan一直没有关闭，不会一直阻塞而deadlock
			//会自动到下一个case匹配
			case v := <-intChan:
				fmt.Printf("从intChan中读取数据%d\n", v)
			case v := <-stringChan:
				fmt.Printf("从stringchan中取数据%s\n", v)
			default:
				fmt.Println("都取不到数据了")
				return
		}
	}
}

func sayHello(){
	for i := 0; i < 10; i++{
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}
func test(){
	defer func(){
		//捕获test抛出的panic,其它协程不会受到影响
		if err := recover(); err != nil{
			fmt.Println("发生异常")
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}
func learn3(){
	go sayHello()
	go test()
	for i := 0; i < 10; i++{
		time.Sleep(time.Second)
		fmt.Println("main() ok")
	}
}

func main(){
	learn3()
}