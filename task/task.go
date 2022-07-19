package task

import (
	"fmt"
	"time"

	"github.com/tsxylhs/go-starter/task"
)

func DoSomeThing() interface{} {
	return 1 + 34 + 3
}
func TaskSimpleConcurency() {
	result := make(chan interface{})

	task := task.Task{
		CountPool:   2,
		TimeOut:     time.Now().Add(10 * time.Microsecond),
		Results:     result,
		Description: []string{"简单的并发模型"},
		Do:          DoSomeThing,
	}
	task.SimpleConcurrency()

	for {
		select {
		case val, _ := <-task.Results:
			fmt.Print(val)
			return
		}
	}
}
