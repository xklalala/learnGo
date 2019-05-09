package main
//稀疏数组,三元组
import (
	"fmt"
)

//用三元组保存二维数组

type spareArray struct{
	i		int
	j 		int
	value	int
}

func main(){
	var chessMap [11][11]int
	var SpareArray []spareArray
	chessMap[1][2] = 1	//表示白子
	chessMap[2][3] = 2	//表示黑子
	chessMap[2][4] = 1
	chessMap[3][3] = 1
	chessMap[5][4] = 2
	chessMap[4][5] = 2
	// for _, v := range chessMap{
	// 	fmt.Println(v)
	// }
	val := spareArray{
		i:len(chessMap),
		j:len(chessMap[0]),
		value:0,
	}
	SpareArray = append(SpareArray, val)
	for key, value := range chessMap{
		for k, v := range value{
			if v != 0{
				//创建一个节点
				val := spareArray{
					i:key,
					j:k,
					value:v,
				}
				SpareArray = append(SpareArray, val)
			}
		}
	}
	fmt.Println(SpareArray)

	//恢复原始数组
	//先创建一个原始数组
	fmt.Printf("%T", SpareArray[0].i)



}
