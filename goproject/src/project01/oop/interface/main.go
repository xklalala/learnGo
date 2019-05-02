package main
import "fmt"

type Usb interface{
	Start()
	Stop()
}

type Phone struct{

}
func (phone Phone) Start(){
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop(){
	fmt.Println("手机停止工作")
}

type Camera struct{

}
func (camera Camera) Start(){
	fmt.Println("相机开始工作")
}
func (camera Camera) Stop(){
	fmt.Println("相机停止工作")
}

type Computer struct{

}

func (computer Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}

func learn1(){
	computer := Computer{}
	phone    := Phone{}
	camera   := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}

func main(){
	learn1()
	fmt.Println()
}