package model

//定义一个结构体
type student struct{
	Name  string
	Score float64
}
//工厂模式
func NewStudent(name string, score float64) *student{
	return &student{
		Name : name,
		Score : score,
	}
}