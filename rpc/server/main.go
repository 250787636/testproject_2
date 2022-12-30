package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

/*
//第一版
type HelloServer struct{}

func (h *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
//func main() {
//	// 将helloService类型的对象注册为一个RPC服务
//	// rpc.Register函数调用将对象类型所满足的RPC规则的对象方法注册为RPC函数
//	// 所有注册方法在“HelloService”服务空间之下
//	err := rpc.RegisterName("HelloService", new(HelloServer))
//	if err != nil {
//		panic(err)
//	}
//	// 然后建立一个唯一的TCP链接，通过rpc.ServeConn函数在该TCP上为对方提供RPC服务
//	lis, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		panic(err)
//	}
//	con, err := lis.Accept()
//	if err != nil {
//		panic(err)
//	}
//	rpc.ServeConn(con)
//}
// 第二版
func main() {
	// 实例化一个server
	_ = rpc.RegisterName("HelloService", &HelloServer{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
*/

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

// 基于上下文信息 可以为RPC增加 登录状态的验证
// 这样可以要求在客户端链接RPC服务时,首先执行登陆操作,登录成功后才能正常执行其他的服务。
func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

// hello方法中可根据conn成员时报不同链接的RPC调用
func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ",from" + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// 为每个链接启动独立的RPC服务:
		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			p.Register(&HelloService{conn: conn})
			p.ServeConn(conn)
		}()
	}
}
