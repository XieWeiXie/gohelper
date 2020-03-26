package main

import "github.com/wuxiaoxiaoshen/gohelper"

func main(){
	s := gohelper.NewServerTcp("127.0.0.1", 4444)
	s.Listen()
}
