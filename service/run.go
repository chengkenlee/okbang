package service

import (
	"github.com/robfig/cron/v3"
	"okbang/util"
	"strings"
	"sync"
)

func Crontab() {
	util.Logger.Info("定时挂载...")
	crontab := cron.New()
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	_, err := crontab.AddFunc(util.Config.GetString("okbang.crontab"), run)
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {}
	//run()
}

func test() {
	util.Logger.Info("running...")
}

func run() {
	util.Logger.Info("定时运行...")
	Chrome()
	html := Select(i.Count)
	if len(html) == 0 {
		util.Logger.Info("检测到所有物业费用已经缴清~")
		return
	}
	emails := util.Config.GetStringSlice("email.to")
	var wait sync.WaitGroup
	for _, email := range emails {
		wait.Add(1)
		go func(email string) {
			defer wait.Done()
			owner := strings.ToUpper(strings.Split(email, "@")[0])
			SendMail(strings.ReplaceAll(html, "Owner", owner), email)
		}(email)
	}
	wait.Wait()
}
