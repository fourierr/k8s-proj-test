package main

import (
	"fmt"
	"k8s-proj-test/wait"
	"time"
)

func main() {
	fmt.Println("start: " + time.Now().Format("2006-01-02 15:04:05.99"))
	// TestJitterUntil
	stopCh := make(chan struct{})
	go wait.TestJitterUntil(stopCh)
	//time.Sleep(10 * time.Second)
	//close(stopCh)
	//fmt.Println("执行结束")

	// TestJitterUntilWithContext
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancelFunc()
	//go wait.TestUntilWithContext(ctx)

	// TestForever
	//go wait.TestForever()

	// TestPollCondTrue
	//stopCh := make(chan struct{})
	//go wait.TestPollCondTrue(stopCh)
	//go wait.TestPollCondFalse(stopCh)

	// TestPollCondTrue
	//go wait.TestPollImmediateCondFalse(stopCh)

	time.Sleep(300 * time.Second)
}
