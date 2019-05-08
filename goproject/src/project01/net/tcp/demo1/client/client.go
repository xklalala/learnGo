package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("error", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)//os.Stdin代表标准输入即终端
	for{
		//从终端读取一行用户的输入，并准备发送给服务器
		str, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
		}
		str = strings.Trim(str, "\r\n")
		if str == "exit"{
			fmt.Printf("客户端退出")
			return
		}

		//将str发送给服务器
		n , err := conn.Write([]byte(str+"\n"))
		if err != nil{
			fmt.Println(err)
		}
		if n > 0{
			
		}

	}
}