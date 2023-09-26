package service

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"okbang/util"
	"strconv"
	"time"
)

var i *IHome

type IHome struct {
	Home  string
	Area  float64
	Year  string
	Water float64
	Manr  float64
	Count float64
}

func Chrome() {
	// 加快渲染速度
	imagCaps := map[string]interface{}{
		//禁止加载图片
		"profile.managed_default_content_settings.images": 2,
		//禁止加载cs
		"profile.managed_default_content_settings.stylesheet": 2,
		//禁止加载flash
		"profile.managed_default_content_settings.flash": 2,
		//禁止加载js
		//"profile.managed_default_content_settings.javascript": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) " +
				"AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}

	var ops []selenium.ServiceOption
	//selenium.SetDebug(true)
	service, _ := selenium.NewChromeDriverService(util.Config.GetString("selenium.chrome"), 9515, ops...)
	defer func(service *selenium.Service) {
		err := service.Stop()
		if err != nil {
			util.Logger.Error(err.Error())
		}
	}(service)

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	caps.AddChrome(chromeCaps)

	wd, _ := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	/*启用等待方式, 超出还未加载元素则超时*/
	err := wd.SetImplicitWaitTimeout(time.Second * 10)
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	/*end*/
	defer func(wd selenium.WebDriver) {
		err := wd.Quit()
		if err != nil {
			util.Logger.Error(err.Error())
			return
		}
	}(wd)
	err = wd.Get(util.Config.GetString("selenium.url"))
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}

	we, _ := wd.FindElement(selenium.ByID, "cliName") //找到姓名输入框的id
	err = we.SendKeys(util.Config.GetString("okbang.cliName"))
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	we, _ = wd.FindElement(selenium.ByID, "mobile") //找到手机号输入框的id
	err = we.SendKeys(util.Config.GetString("okbang.mobile"))
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	we, _ = wd.FindElement(selenium.ByID, "search") //找到点击的id
	// 点击
	err = we.Click()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	util.Logger.Info("自动登陆成功")

	tmp, _ := wd.FindElement(selenium.ByXPATH, "//*[@id=\"addr\"]") //地址
	address, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	tmp, _ = wd.FindElement(selenium.ByXPATH, "//*[@id=\"jzmj\"]") //面积
	jzmj, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	tmp, _ = wd.FindElement(selenium.ByXPATH, "//*[@id=\"year\"]/option[1]")
	year, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	tmp, _ = wd.FindElement(selenium.ByXPATH, "//*[@id=\"month\"]/option[4]")
	month, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	tmp, _ = wd.FindElement(selenium.ByXPATH, "//*[@id=\"jmglfeehj0\"]")
	waterRate, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	tmp, _ = wd.FindElement(selenium.ByXPATH, "//*[@id=\"jmglfeehj1\"]")
	manRate, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	tmp, _ = wd.FindElement(selenium.ByXPATH, "//*[@id=\"price2\"]")
	price2, err := tmp.Text()
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}

	j, _ := strconv.ParseFloat(jzmj, 64)
	w, _ := strconv.ParseFloat(waterRate, 64)
	m, _ := strconv.ParseFloat(manRate, 64)
	p, _ := strconv.ParseFloat(price2, 64)

	i = &IHome{
		Home:  address,
		Area:  j,
		Year:  fmt.Sprintf("%s.%s", year, month),
		Water: w,
		Manr:  m,
		Count: p,
	}
	util.Logger.Info(fmt.Sprintf("家:%s,面积:%s,计费月份:%s.%s,水费:%s,管理费:%s,未缴费合计:%s", address, jzmj, year, month, waterRate, manRate, price2))
}
