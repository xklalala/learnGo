package lib
import (
	"fmt"
	"../common"
	"net"
	"encoding/json"
	"encoding/binary"
)

//写一个函数，完成登陆
func Login(userId int, userPwd int) (err error){
	// fmt.Printf("userId=%d, pwd=%d", userId, pwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil{
		fmt.Println("网络错误", err)
		return
	}
	//延时关闭
	defer conn.Close()
	fmt.Println(conn, err)
	//准备通过conn发送消息给服务器
	var mes common.Message
	mes.Type = common.LoginMesType
	//创建一个loginmes结构体
	loginMes :=  common.LoginMes{
		UserId : userId,
		UserPwd : userPwd,
	}
	
	//将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	//把data赋给message的data字段
	mes.Data = string(data)

	//将mes进行序列化
	send, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//send就是我们要发送的信息

	//先把send的长度发送给服务器
	//先获取到send的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(send))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)

	//发送长度
	n, err := conn.Write(bytes)
	if n==0 || err != nil{
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("客户端消息的长度发送成功")

	return nil
}