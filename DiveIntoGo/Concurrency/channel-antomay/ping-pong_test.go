package channel_antomay

import (
	"fmt"
	"testing"
	"time"
)

type Ball struct{ hits int }

func TestPingPong(t *testing.T) {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // game on; toss the ball // 注释代码这行代码会怎么样？
	time.Sleep(2 * time.Second)
	<-table // game over; grab the ball  // 注释这行代码呢？会怎么样？

	panic("show me the stacks") // 还可以用 panic 来简单演示堆栈情况，学到了
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

