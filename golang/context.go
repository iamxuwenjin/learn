package golang

import (
	"context"
	"fmt"
	"time"
)

type orderID int

func contextDemo() {
	var x = context.TODO()
	x = context.WithValue(x, orderID(1), "1234")
	x = context.WithValue(x, orderID(2), "2345")

	y := context.WithValue(x, orderID(3), "4567")
	x = context.WithValue(x, orderID(3), "3456")

	fmt.Println(x.Value(orderID(3)))
	fmt.Println(y.Value(orderID(3)))
}

func valueCtx() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, 1, "123")
	ctx = context.WithValue(ctx, 2, "234")
	ctx = context.WithValue(ctx, 3, "345")
	ctx_ := context.WithValue(ctx, 3, "567")
	fmt.Println(ctx.Value(3))
	fmt.Println(ctx_.Value(3))
}

func cancelCtx() {
	jobChan := make(chan struct{})
	ctx, cancelFn := context.WithCancel(context.TODO())
	worker := func() {
	jobLoop:
		for {
			select {
			case <-jobChan:
				fmt.Println("do my job")
			case <-ctx.Done():
				fmt.Println("parent call me to quit")
				break jobLoop
			}
		}
	}

	// start worker
	go worker()
	jobChan <- struct{}{}
	// stop all worker
	cancelFn()
	time.Sleep(time.Second)
}
