package main
import (
	"fmt"
	"time"
)
// 实例1
// 完成goroutine和channel协同工作的案例，具体要求
// 	1.开启一个writeData协程，向管道intChan中写入50个整数
// 	2.开启一个readData协程，从管道intChan中读取writeData写入的数据
// 	2.注意writeData和readData操作的是同一个管道
// 	4.主线程需要等待writeData和readData协程都完成工作才能退出
func writeData(intChan chan int){
	for i := 1; i<= 500; i++{
		intChan<- i
		fmt.Println("写入数据", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool){
	for{
		v, ok := <-intChan
		if !ok{
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//读取完数据后，任务完成
	exitChan<- true
	close(exitChan)
}

func learn1(){
	//创建两个管道
	intChan := make(chan int, 500)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for{
		_, ok := <-exitChan
		if !ok{
			break
		}
	}

}

//统计1-200000的数字中哪些是素数


func putNum(intChan chan int){
	for i := 1; i <= 200000; i++{
		intChan<- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool){
	//使用for循环
	var flag bool
	for{
		num, ok := <-intChan
		if !ok{
			break
		}
		flag = true //假定这个数字是素数
		for i := 2; i < num; i++{
			if num % i == 0{
				flag = false
				break
			}
		}
		if flag{
			primeChan<- num
		}
	}
	fmt.Println("有一个协程因为取不到数据，退出")
	//这里我们还不能关闭管道
	exitChan<- true
}

func learn2(){
	intChan := make(chan int, 100000)
	primeChan := make(chan int, 100000)//放结果
	exitChan := make(chan bool, 4)//结束的管道

	//开启一个协程，放入1-20万个数
	go putNum(intChan)

	//开启四个协程，从intChan取出数据，并判断是否是素数，如果是就放到primeChan中
	for i := 0; i<4; i++{
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里主线程要进行处理
	go func(){
		for i := 0; i < 4; i++{
			<-exitChan
		}
		//当我们从exitChan中取出了4个数据，可以关闭管道
		close(primeChan)
	}()

	//遍历primeNum，把结果取出
	for{
		res, ok := <-primeChan
		if !ok{
			break
		}
		fmt.Println(res)
	}
	fmt.Println("主线程退出")
	
}
func main(){
	start := time.Now().UnixNano()
	learn2()
	end := time.Now().UnixNano()
	fmt.Printf("执行test耗费的时间为%f秒\n",float64( (end - start))/1000000000)
	
}