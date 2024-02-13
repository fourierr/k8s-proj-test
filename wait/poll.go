package wait

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

func PrintTimeCondTrue() (done bool, err error) {
	fmt.Println("hello: " + time.Now().Format("2006-01-02 15:04:05.99"))
	return true, nil
}

func PrintTimeCondFalse() (done bool, err error) {
	fmt.Println("hello: " + time.Now().Format("2006-01-02 15:04:05.99"))
	return false, nil
}

func TestPollCondTrue(stopCh <-chan struct{}) {
	err := wait.Poll(3*time.Second, 5*time.Second, PrintTimeCondTrue)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestPollCondFalse(stopCh <-chan struct{}) {
	err := wait.Poll(3*time.Second, 20*time.Second, PrintTimeCondFalse)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestPollImmediateCondFalse(stopCh <-chan struct{}) {
	err := wait.PollImmediate(3*time.Second, 20*time.Second, PrintTimeCondFalse)
	if err != nil {
		fmt.Println(err)
		return
	}
}
