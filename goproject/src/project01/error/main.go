package main
import(
	"fmt"
	"errors"
)

func test(){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

//自定义错误
func readConf(name string) (err error){
	if name == "config.ini"{
		return nil
	}else{
		return errors.New("读取文件错误")
	}
}

func test2(){
	err := readConf("configini")
	if err != nil{
		panic(err)
	}
	fmt.Println("继续执行")
}

func main(){
	
	// test()
	fmt.Println(2345)
	test2();
}