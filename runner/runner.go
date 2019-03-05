package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner Runner
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time // 只可以用来接收time.Time类型的数据
	tasks     []func(int)
}

// ErrTimeout ErrTimeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt ErrInterrupt
var ErrInterrupt = errors.New("received interrupt")

// New new
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d), // 等待参数duration时间后，向返回的chan里面写入当前时间。
	}
}

// Add add
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
