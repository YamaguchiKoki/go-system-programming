package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(3000)) * time.Millisecond)

	num := r.Intn(10)
	c <- num
}

/*
	Memo
	- メインのゴルーチンが終わったら、他ゴルーチンの終了を待たずにプログラム全体が終了

	- waitGroup構造体は内部でカウンターを保持
	- Doneはwg.Add(-1)をラップしているだけ
	- Waitはカウンターが0になるまでブロックする
*/
func main() {
	fmt.Println("what is today's lucky number?")

	c := make(chan int)
	go getLuckyNum(c)

	// メインのごルーチンはチャネルから値を受け取るまでブロックされる
	num := <-c
	fmt.Printf("Today's your lucky number is %d!\n", num)

	close(c)

}
