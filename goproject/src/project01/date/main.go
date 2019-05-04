package main
import(
	"fmt"
	"time"
	"strconv"
)

func test(){
	str := ""
	for i := 0; i< 10000; i++{
		str += "hello" + strconv.Itoa(i)
	}
}

func main(){
	now := time.Now()//获得当前时间
	fmt.Printf("%v, %T \n", now, now)
	
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//格式化输出
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf(now.Format("2006/01/02 15:04:05")+"\n")//输出当前时间
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	
	//时间常量
	//time.Second
	//time.Millisecond 毫秒
	//每隔10ms打印一个数字
	i := 0
	for{
		i++
		fmt.Printf("%d ", i)
		time.Sleep(10 * time.Millisecond)
		if i == 10{
			break
		}
	}

	//时间戳
	//Unix		秒数
	//UnixNano  纳秒数
	fmt.Printf("\n%v, %v \n",now.Unix(), now.UnixNano())

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("执行test耗费的时间为%f秒\n",float64( (end - start))/1000000000)

}