package wait

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

func PrintTime() {
	fmt.Println("hello: " + time.Now().Format("2006-01-02 15:04:05.99"))
}
func PrintTimeWithCyx(ctx context.Context) {
	fmt.Println("hello: " + time.Now().Format("2006-01-02 15:04:05.99"))
}

/*
hello: 2024-01-28 22:13:50.78
hello: 2024-01-28 22:13:53.98
hello: 2024-01-28 22:13:57.28
hello: 2024-01-28 22:14:00.48
执行结束
*/
func TestJitterUntil(stopCh <-chan struct{}) {
	wait.JitterUntil(PrintTime, 3*time.Second, 0.1, true, stopCh)
}

/*
hello: 2024-01-28 22:27:21.93
hello: 2024-01-28 22:27:24.95
hello: 2024-01-28 22:27:27.96
hello: 2024-01-28 22:27:30.96

*/
func TestUntilWithContext(ctx context.Context) {
	wait.UntilWithContext(ctx, PrintTimeWithCyx, 3*time.Second)
}

func TestForever() {
	wait.Forever(PrintTime, 3*time.Second)
}
