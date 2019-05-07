package main
import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	// bson "gopkg.in/mgo.v2/bson"
	// "log"
)

// func ConnectToDB() *mgo.Collection{

// }

func main(){
	url := "mongodb://39.106.148.118:27017/shop"
	session, err := mgo.Dial(url)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(session)
}