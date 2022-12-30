package main

import (
	"fmt"
	"net/rpc"
)

// 调用server中 的helloService方法
//func main() {
//	// 首先通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法
//	cli, err := rpc.Dial("tcp", ":1234")
//	if err != nil {
//		panic(cli)
//	}
//
//	var reply string
//	// 调用client.Call时
//	// 第一个参数用点号链接的RPC服务名称何方法名称
//	// 第二和第三个参数分别为我们定义RPC方法的两个参数
//	err = cli.Call("HelloService.Hello", "你好", &reply)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(reply)
//}

// watch方法调用
//func doClientWork(client *rpc.Client) {
//	go func() {
//		var keyChanged string
//		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println("watch:", keyChanged)
//	}()
//	err := client.Call(
//		"KVStoreService.Set", [2]string{"abc", "abc-value"},
//		new(struct{}),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	time.Sleep(time.Second * 3)
//}
//
//// 调用watch方法
//func main() {
//	cli, err := rpc.Dial("tcp", ":1234")
//	if err != nil {
//		panic(err)
//	}
//	doClientWork(cli)
//}

func main() {
	// 通过rpc.Dial拨号RPC服务,然后通过client.Call调用具体的RPC方法
	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply2 string

	// 因为server 添加了登录认证
	// panic: please login
	err = cli.Call("HelloService.Login", "user:password", nil)
	if err != nil {
		panic(err)
	}

	err = cli.Call("HelloService.Hello", "你好", &reply2)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply2)
}
