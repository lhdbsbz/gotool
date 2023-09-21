package gotool

import (
	"fmt"

	"sync"
)

type WorkerPool struct {
	workerWaitGroup *sync.WaitGroup
	jobChannel      chan func()
	WorkerNum       int
	JobBufferNum    int
}

func NewWorkerPool(workerNum int, jobBufferNum int) *WorkerPool {
	workerPool := new(WorkerPool)
	workerPool.jobChannel = make(chan func(), jobBufferNum)
	workerPool.WorkerNum = workerNum
	workerPool.workerWaitGroup = new(sync.WaitGroup)
	workerPool.workerWaitGroup.Add(workerNum)
	for i := 0; i < workerPool.WorkerNum; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("123")
				}
				workerPool.workerWaitGroup.Done()
			}()
			for {
				fn := <-workerPool.jobChannel
				if fn == nil {
					break
				}
				fn()
			}
		}()
	}
	return workerPool
}

func (wp *WorkerPool) Add(fn func()) {
	wp.jobChannel <- fn
}

func (wp *WorkerPool) Run() {
	close(wp.jobChannel)
	wp.workerWaitGroup.Wait()
}
