package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
	"xiamei.guo/blog-api/models"
)

func corn() {
	log.Println("Starting...")

	c := cron.New()                   //会根据本地时间创建一个新（空白）的 Cron job runner。会首先解析时间表，如果填写有问题会直接 err，无误则将 func 添加到 Schedule 队列中等待执行
	c.AddFunc("* * * * * *", func() { //AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start() //在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制哦
	//（1）time.NewTimer: 会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	//（2）for + select: 阻塞 select 等待 channel
	//（3）t1.Reset: 会重置定时器，让它重新开始计时

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
