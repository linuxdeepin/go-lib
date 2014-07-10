package utils

import (
	"fmt"
	. "launchpad.net/gocheck"
	"math/rand"
	"time"
)

func doTask(taskName, childTaskName string) {
	fmt.Printf("%s: child task %s begin\n", taskName, childTaskName)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("%s: child task %s end\n", taskName, childTaskName)
}

func (*UtilsTest) TestOverrideRunner(c *C) {
	rand.Seed(time.Now().UnixNano())
	r := NewOverrideRunner()
	go r.Loop()
	for i := 1; i <= 5; i++ {
		taskName := fmt.Sprintf("task group %d", i)
		r.AddTaskGroup(func() {
			doTask(taskName, "one")
		}, func() {
			doTask(taskName, "two")
		})
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
	time.Sleep(2000 * time.Millisecond)
}
