package main

import (
	"k8s-proj-test/wait"
	"time"
)

func main() {
	// TestJitterUntil
	//stopCh := make(chan struct{})
	//go wait.TestJitterUntil(stopCh)
	//time.Sleep(10 * time.Second)
	//close(stopCh)
	//fmt.Println("执行结束")

	// TestJitterUntilWithContext
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancelFunc()
	//go wait.TestUntilWithContext(ctx)

	// TestForever
	go wait.TestForever()

	time.Sleep(300 * time.Second)
}
