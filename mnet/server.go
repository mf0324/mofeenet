package mnet

import {
	"fmt"
	"net"
	"interface"
}

//Server 服务器实现
type Server struct {
	Name string
	IPVersion string
	IP string
	port int
}

func (s *Server)Start() {
	fmt.Printf("[START]	Server	listenner	at	IP:	%s,	Port	%d,	is	starting\n", s.IP, s.Port)
	go func {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if	err	!=	nil	{
			fmt.Println("resolve tcp addr err:", err)
			return
		}
	
	}
}