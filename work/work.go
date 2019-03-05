package work

import "sync"

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个goroutine池，这个池可以完成任何已提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(macGorontines int) *Pool {
	p := Pool{
		work: make(chan Worker), // 无缓冲队列
	}

	p.wg.Add(macGorontines)
	for i := 0; i < macGorontines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有goruntine停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
