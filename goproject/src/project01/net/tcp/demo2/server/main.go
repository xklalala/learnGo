package main
import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	//读客户发送的信息
	defer conn.Close()
	for{

		buf := make([]byte, 8096)
		_, err := conn.Read(buf)
		if err != nil{
			fmt.Println("读取客户端发送的数据")
		}
		fmt.Println("读到的buff=", string(buf))
	}
}

func main(){
	//提示信息
	fmt.Println("服务器在8889端口监听。。。")
	list, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil{
		fmt.Println("net.Listen err=", err)
		return
	}
	defer list.Close()
	//一旦监听成功，就等待客户端来连接服务器
	for{
		fmt.Println("等待客户端来连接服务器。。。。")
		conn, err := list.Accept()
		if err != nil{
			
		}
		go process(conn)
	}
}