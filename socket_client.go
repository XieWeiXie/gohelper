package gohelper

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type ClientTcp struct {
	IP *net.TCPAddr
	Port int
	con *net.TCPConn
	socket string
}

func NewClientTcp(host string, port int) *ClientTcp{
	socket := fmt.Sprintf("%s:%d", host, port)
	ip, _ := net.ResolveTCPAddr("tcp4", socket)
	con, _ := net.DialTCP("tcp", nil,ip)
	return &ClientTcp{
		IP:   ip,
		Port: port,
		con: con,
		socket: socket,
	}
}

func (c *ClientTcp) Dial() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(c.socket + "> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		c.con.Write([]byte(text))

		var a =make([]byte,512)
		c.con.Read(a)

		//ioutil.ReadAll(c.con)
		fmt.Println(c.socket + "> " +string(a))

	}
}
