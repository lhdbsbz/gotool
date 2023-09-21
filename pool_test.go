package gotool

import (
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/atomic"
)

var n atomic.Int64

func getItems() []int {
	time.Sleep(1 * time.Second)
	if n.Load() >= 20 {
		return []int{}
	}
	rand.Seed(time.Now().Unix())
	res := []int{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)}
	return res
}

func f(i int) func() {
	return func() {
		fmt.Println("num: ", i, "start work")
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		n.Add(1)
	}
}
