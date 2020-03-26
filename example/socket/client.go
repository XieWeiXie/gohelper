package main

import "github.com/wuxiaoxiaoshen/gohelper"

func main(){
	client := gohelper.NewClientTcp("127.0.0.1",4444)
	client.Dial()
}
