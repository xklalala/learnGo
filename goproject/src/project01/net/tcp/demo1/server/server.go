package main
import (
	"fmt"
	"net"
)


func process(conn net.Conn){
	//这里我们循环接收客户端发送的数据
	defer conn.Close()//这里记得关闭连接

	for{
		//创建一个新的切片
		buf := make([]byte, 1024)
		//如果对方没有发送数据，这个协程就一直阻塞在这里
		// fmt.Printf("服务器在等待客户端%v发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)	//从conn读取
		if err != nil{
			fmt.Println("服务端的read错误",err)
			return
		}
		//显示客户端发送的数据到终端
		fmt.Printf(string(buf[:n]))
	}
}

func learn1(){
	fmt.Println("服务器开始监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")//在本地监听8888端口
	if err != nil{
		fmt.Println("listen error", err)
		return
	}
	defer listen.Close()	//延时关闭listen

	//循环等待
	for{
		fmt.Println("等待连接。。。")
		conn, err := listen.Accept()//等待客户端连接
		if err != nil{
			fmt.Println("Accept error", err)
		}else{
			fmt.Printf("连接成功， 客户端ip为：%v\n", conn.RemoteAddr().String())
		}
		//这里写一个协程为客户端服务
		go process(conn)
	}

}


func main(){
	learn1()
}