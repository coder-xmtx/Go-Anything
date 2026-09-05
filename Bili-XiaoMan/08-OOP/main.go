package main

import "fmt"

// 定义函数
// 这里就是一个接口，意思就是：凡是能付钱的，必须有一个 Pay 功能，并且传入 money，返回 string 和 float64
// 刚好，Alipay 和 WeChatPay 这两个都有 Pay 方法，而且长得一样
type Payment interface {
	Pay(money float64) (string, float64)
}

// 支付宝支付的结构体
type Alipay struct {
	AppId string
}

// 结构体写方法
func (a Alipay) Pay(money float64) (string, float64) {
	fmt.Println("alipay id:", a.AppId)
	return "调用成功", money
}

// 微信支付的结构体
type WeChatPay struct {
	WxId string
}

// 结构体写方法
func (w WeChatPay) Pay(money float64) (string, float64) {
	fmt.Println("weChat id:", w.WxId)
	return "调用成功", money
}

// 业务层封装，多态
func CreateOrder(pay Payment, money float64) {
	str, num := pay.Pay(money)
	fmt.Println(str, num)
}

func main() {

	// 初始化
	ali := Alipay{AppId: "zfb_001"}
	// 调用
	str, money := ali.Pay(66.66)
	fmt.Println(str, money)

	// 业务层封装调用
	CreateOrder(Alipay{AppId: "zfb_002"}, 90)
}
