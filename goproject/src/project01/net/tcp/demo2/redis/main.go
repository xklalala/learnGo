package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func learn1(){
	fmt.Println()
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")//连接
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()//关闭
	// fmt.Println(conn)
	_, err = conn.Do("set", "name", "lalala")//写入数据
	if err != nil{
		fmt.Println(err)
		return
	}
	res, err := redis.String(conn.Do("get", "name"))//conn.Do("get", "name")返回一个interface{}
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)

	//操作hash
	_, err = conn.Do("HSet", "student1", "name", "啦啦啦")//写入数据
	if err != nil{
		fmt.Println(err)
		return
	}
	_, err = conn.Do("HSet", "student1", "age", "15")//写入数据
	if err != nil{
		fmt.Println(err)
		return
	}
	res2, err := redis.String(conn.Do("HGet", "student1", "name"))//conn.Do("get", "name")返回一个interface{}
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res2)
	res3, err := redis.String(conn.Do("HGet", "student1", "age"))//conn.Do("get", "name")返回一个interface{}
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res3)
	

	//批量set/get
}


func main(){
	learn1()
}