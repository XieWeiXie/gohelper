package gohelper

import (
	"fmt"
	"log"
	"net"
	"time"
)


/*
TCP: 流式套接字:一种可靠的、面向连接的双向数据传输服务
UDP: 数据报套接字:一种无连接、不可靠的双向数据传输服务
IP: 原始套接字:允许对较低层协议（如 IP 或 ICMP）进行直接访问

*/

type ServerTcp struct {
	IP   *net.TCPAddr
	Port int
}

func NewServerTcp(host string, port int) *ServerTcp {
	ip, e := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", host, port))
	if e != nil {
		panic(e)
	}
	return &ServerTcp{
		IP:   ip,
		Port: port,
	}
}

func (s *ServerTcp) Listen() {
	listener, _ := net.ListenTCP("tcp", s.IP)
	log.Println(s.IP.String())
	for {
		con, _ := listener.Accept()
		// 处理
		go s.handle(con)
	}
}

func (s *ServerTcp) handle(con net.Conn) {
	defer con.Close()
	for {
		var b []byte
		b = make([]byte,1024)
		con.Read(b)
		fmt.Println(string(b))
		con.Write([]byte(fmt.Sprintf("%s Success!", time.Now().String())))
	}
}
