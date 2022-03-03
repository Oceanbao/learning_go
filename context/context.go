// context.Background() empty impl Context interface

// no vlaues
// never cancelled
// no deadline

// context.ToDo()
// - returns empty context, when surrounding func not been passed a context
// and need to use context as placeholder in current func and plans to add acutal
// context in near future

// Context Tree
// - Background is root of all context
// - a new context is created by wrapping existing immutable context
// and adding data

// 2-level Tree
// rootCtx := context.Background()
// childCtx := context.WithValue(rootCtx, "msgId", "someMsgId")
// - childCtx derived having data storing request-scoped values

// 3-level Tree
// rootCtx := context.Background()
// childCtx := context.WithValue(rootCtx, "msgId", "someMsgId")
// childOfChildCtx, cancelFunc := context.WithCancel(childCtx)
// - second value is trigger of signal to cancel

// Deriving From Context
// - passing request-scoped values - WithValue()
// - cancle signal - WithCancel()
// - deadline - WithDeadline()
// - timeout - WithTimeout()

// BP and Caveat
// - Do not store a context within a struct type!!!
// - Should flow through program - HTTP request, a new context can be
// created for each incoming request to hold a request_id or some info
// - Always pass context as first arg to function
// - Whenever not sure to use, better to use Context.ToDo() as placeholder
// - Only the parent goroutine or func should cancel context - do not
// pass cancelFunc to downstram goroutines or functions!!!

// CookieJar

package main

import (
	"context"
	"fmt"
	"time"
)

func taskCancel(ctx context.Context) {
	i := 1

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefuly exit (cancelled)")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func taskTimeout(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefuly exit (Timeout)")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func taskDeadline(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefuly exit (Deadline)")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

// RunContextDemo runs demo for context
func RunContextDemo() {

	// withCancel
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	go taskCancel(cancelCtx)

	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)

	// withTimeout
	cancelCtxTimeout, cancelFuncTimeout := context.WithTimeout(ctx, time.Second*3)
	defer cancelFuncTimeout()
	go taskTimeout(cancelCtxTimeout)
	time.Sleep(time.Second * 4)

	// withDeadline
	cancelCtxDeadline, cancelFuncDeadline := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
	defer cancelFuncDeadline()
	go taskDeadline(cancelCtxDeadline)
	time.Sleep(time.Second * 6)
}

/*
v := Values{map[string]string{
    "1": "one",
    "2": "two",
}}

c := context.Background()
c2 := context.WithValue(c, "myvalues", v)

fmt.Println(c2.Value("myvalues").(Values).Get("2"))
*/
