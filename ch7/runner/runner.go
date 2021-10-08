package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeOut   <-chan time.Time
	tasks     []func(int)
}

var ErrInterrupt = errors.New("received interrupt")
var ErrTimeout = errors.New("received timeout")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		//After返回Time的管道,会在经过d时间之后向这个通道发送一个time.Time的值
		timeOut: time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	//注册程序接收的系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()
	//select可以监听多个channel的读写事件,如果没有任何管道接受到数据会阻塞,除非有default分支
	select {
	case err := <-r.complete:
		return err
	case <-r.timeOut:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已注册的任务
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
