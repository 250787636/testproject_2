package main

import "fmt"

type RoundRobin struct {
	servers []string
	current int
}

func (R *RoundRobin) next() string {
	R.current++
	R.current = R.current % len(R.servers) // 访问到最后时 从0重新访问 3%3=0
	return R.servers[R.current]
}

func main() {
	r := &RoundRobin{
		servers: []string{"192.168.1", "192.168.2", "192.168.3"},
		current: -1,
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("|%d|%s|\n", i+1, r.next())
	}
}
