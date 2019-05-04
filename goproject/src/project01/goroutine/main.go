package main
import (
	"fmt"
	"runtime"
	"sync"
)

func learn1(){
	cpuNum := runtime.NumCPU()//获取当前机器有多少个cpu
	fmt.Println("cpu的数量为：", cpuNum)
	runtime.GOMAXPROCS(cpuNum-1)//设置程序最大运行的cpu
	//go1.8以后默让程序运行在多核上，1.8以前要设置一下
}

//需求：现在计算1-200的各个数的阶乘，并且把各个数的阶乘放到map中
//最后显示出来


var (
	myMap = make(map[int]int)

	lock sync.Mutex	//全局互斥锁
)

func learn2_a(n int){
	res := 1
	for i := 1; i<=n; i++{
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func learn2(){
	for i := 1; i <= 20; i++{
		go learn2_a(i)
	}
}

func main(){
	learn2()
	
	lock.Lock()
	for k, v := range myMap{
		fmt.Printf("map[%v]=%v\n", k, v)
	}
	lock.Unlock()
}