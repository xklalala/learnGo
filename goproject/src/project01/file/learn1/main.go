package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
	"flag"//获取命令行参数的包
)

func learn1(){
	file, err := os.Open("main.go")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(file)

	err = file.Close()
	if err != nil{
		fmt.Println("文件关闭错误")
	}
}

func learn2(){
	file ,err := os.Open("main.go")
	if err != nil{
		fmt.Println(err)
	}

	//默认缓冲区为4096个字节
	reader := bufio.NewReader(file)

	for{
		str, err := reader.ReadString('\n')

		if err == io.EOF{
			break
		}
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束")


	defer file.Close()
}

func learn3(){
	//一次性读取文件，适用于较小的文件
	file := "main.go"
	content, err := ioutil.ReadFile(file)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%v", string(content))

}

//命令行参数
func learn4(){
	var user 	string
	var pwd		string
	var host	string
	var port	int

	flag.StringVar(&user, "u", 	  "", 			"用户名， 默认为空")
	flag.StringVar(&pwd,  "p", 	  "", 			"密码， 默认为空")
	flag.StringVar(&host, "h", 	  "localhost",  "主机名， 默认localhost")
	flag.IntVar(&port, "port", 3306, 		"端口， 默认3306")
	
	//必须调用下面的方法
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}

func main(){
	learn4()
}