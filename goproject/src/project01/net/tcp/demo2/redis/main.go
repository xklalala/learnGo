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


//连接池
var pool *redis.Pool

func init(){
	pool = &redis.Pool{
		MaxIdle:8,//最大空闲数
		MaxActive:0,//表示数据库的最大连接数，0表示没有限制
		IdleTimeout:100,//最大空闲时间
		Dial:func()(redis.Conn, err){//初始化连接的代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func learn2(){
	conn := pool.Get()
	defer conn.Close()
}

func main(){
	pool.Close()//关闭连接池 
	learn1()
}