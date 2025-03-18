package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: 4_1 <SECOND>")
			os.Exit(0)
    }
    i, err := strconv.Atoi(os.Args[1])
    if err != nil {
      panic(err)
    }

		limit := time.After(time.Second * time.Duration(i))
		// 擬似ジェネレータ
		for {
			select {
			case <- limit:
				fmt.Printf("After %s second\n", os.Args[1])
        os.Exit(0)
			default:
				fmt.Println("tick")
				time.Sleep(100 * time.Millisecond)
			}
		}
}
