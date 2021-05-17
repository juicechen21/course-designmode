package simplefactory

import "fmt"

// 在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。
// simplefactory 简单工厂模式

type API interface {
	Say(name string)string
}

func NewAPI(t int) API {
	if t == 1{
		return &hiAPI{}
	}else if t==2{
		return &helloAPI{}
	}
	return nil
}
type hiAPI struct {}

func (*hiAPI)Say(name string) string {
	return fmt.Sprintf("Hi,%s",name)
}

type helloAPI struct {}

func (*helloAPI)Say(name string)string  {
	return fmt.Sprintf("Hello,%s",name)
}