package main
import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str := "abcdefg "
	fmt.Println(len(str))
	for i := 0; i < len(str); i++{
		fmt.Printf("%c ",str[i])
	}
	n, err := strconv.Atoi("1")//字符串转整数
	if err != nil{
		fmt.Println("转换错误", err)
	}else{
		fmt.Println(n)
	}
	str2 := strconv.Itoa(12345)//字符串转整数
	fmt.Printf("%v, %T \n", str2, str2)
	var bytes = []byte("hello, go")//字符串转byte
	fmt.Printf("%s\n %v\n", bytes, bytes)
	fmt.Println(string([]byte{97, 98, 99}))//byte转string
	fmt.Println(strconv.FormatInt(123, 16))//10进制转2， 8 ， 16进制
	fmt.Println(strings.Contains("hello", "ll"))//查找子串
	fmt.Println(strings.Count("afaswfeiaafawefwofgjawgjawiogaswiogagasfawega", "as"))//统计有多少个子串
	fmt.Println(strings.EqualFold("abc", "AbC"))//判断字符串是否相等，忽略大小写
	fmt.Println(strings.Index("bc_ABC_bc", "bc"))//返回子串在字符串中第一次出现的index
	fmt.Println(strings.LastIndex("bc_ABC_bc", "bc"))//返回子串在字符串中最后一次出现的位置
	fmt.Println(strings.Replace("go go hello", "go", "北京",-1))//将制指定的字符串途替换成另一个字符串，n可以指定希望替换几个，-1表示全部替换
	fmt.Println(strings.Split("hello,word,ok", ","))//按照指定的某个字符，将字符串拆分为字符串数组
	fmt.Println(strings.ToLower("Go"))//将字符串转换为小写
	fmt.Println(strings.ToUpper("Go"))
	fmt.Println(strings.TrimSpace(" Today is sunday "))//去掉两边空格
	fmt.Println(strings.Trim("!hello!", "!"))//将字符串两边指定的字符去掉
	fmt.Println(strings.TrimLeft("!hello!", "!"))
	fmt.Println(strings.TrimRight("!hello!", "!"))
	fmt.Println(strings.HasPrefix("ftp://192.168.1.1", "ftp"))//判断字符串是否以指定的字符串开头
	fmt.Println(strings.HasSuffix("abc.jpg", "png"))//判断字符串是否以指定的字符串结尾
	fmt.Println()
}