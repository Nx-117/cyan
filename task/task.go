package task

import (
	"github.com/robfig/cron/v3"
	"time"
)

/**
spec，传入 cron 时间设置
task，对应执行的任务
*/
func StartJob(spec string, task func()) {
	c := cron.New()
	c.AddFunc(spec, task)
	c.Start()
}

/**
通用定时
t 间隔时间 单位秒
f 执行的方法
*/
func Timing(t time.Duration, f func()) {
	t = t * time.Second
	go func() {
		tick := time.Tick(t)
		for {
			f()
			<-tick // 等待下一个时钟周期到来
		}
	}()
}
