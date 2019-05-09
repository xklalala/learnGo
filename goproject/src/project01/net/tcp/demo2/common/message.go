package common

const (
	LoginMesType	= "LiginMes"
	LoginResMesType	= "LoginResMes"
)

type Message struct{
	Type string	`json:"type"`//消息类型
	Data string `json:"data"`//消息的类型
}

//定义两个消息，后面需要再增加

type LoginMes struct{
	UserId		int		`json:"userId"`//用户id
	UserPwd		int		`json:"userPwd"`//用户密码
	UserName	string	`json:"userName"`//用户名
}

type LoginResMes struct{
	Code 	int		`json:"code"`//返回状态码， 500表示未注册，200表示登陆成功
	Message	string	`json:"message"`//返回错误信息
}