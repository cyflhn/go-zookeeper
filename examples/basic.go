package main

import (
	"github.com/test/go-zookeeper/zk"
	"time"
)

func main() {

	servers := []string{"192.168.28.192:2181"}

	conf := zk.ConnConf{
		RecvTimeout:    5 * time.Second,
		ConnTimeout:    5 * time.Second,
		SessionTimeout: 1,
	}

	c, _, err := zk.ConnectWithConf(servers, conf)
	if err != nil {
		panic(err)
	} else {
		_, err := c.Create("/gozk-test/1", []byte("test"), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			panic(err)
		}
		_, _, err = c.Get("/zk/codis")
		if err != nil {
			panic(err)
		}
		_, _, _, err = c.ChildrenW("/gozk-test")
		if err != nil {
			panic(err)
		}
		c.Close()
	}
}
