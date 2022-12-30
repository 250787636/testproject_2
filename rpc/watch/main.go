package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// m 为map 用于存储KV数据
// filter 为拦截器 用于过滤数据
// mu 为互斥锁 用于Grountine
type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      map[string]string{},
		filter: map[string]func(key string){},
	}
}

// GET SET
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	if p.m == nil {
		p.m = make(map[string]string)
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.m[key] = value
	return nil
}

// 在Set方法中，输入参数是key和value组成的数组,用一个匿名的空结构表示忽略了输出参数
// 当修改某个key对应的值时会调用每一个过滤器 而过滤列表在watch方法中提供:
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	if p.filter == nil {
		p.filter = make(map[string]func(key string))
	}
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
	return nil
}

func main() {
	err := rpc.RegisterName("KVStoreService", new(KVStoreService))
	if err != nil {
		panic(err)
	}
	// 然后建立TCP 并通过rpc.ServeConn函数
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	con, err := lis.Accept()
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(con)
}
