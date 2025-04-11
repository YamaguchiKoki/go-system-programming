package main

import (
	"fmt"
	"time"
)

type Conn struct {
	ID int
}

type Result struct {
	ConnID int
	Result string
}

func (c *Conn) DoQuery() Result {
	time.Sleep(time.Second * 1)
	return Result{
		ConnID: c.ID,
		Result: "done",
	}
}

func Query(cons []Conn) Result {
	ch := make(chan Result, len(cons))
	for _, conn := range cons {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery():
			default:
			}
		}(conn)
	}
	return <-ch
}

func main() {
	conns := []Conn{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	// 一番早く結果が返ってくるものを取得する
	result := Query(conns)
	fmt.Println(result)
}
