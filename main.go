/*
 *@author ChengKen
 *@date   16/03/2023 14:04
 */
package main

import (
	"okbang/util"
	"sync"
)

var wait sync.WaitGroup

func main() {
	util.Parm()
	util.Loggers()
	Run()

	return
	html := okbang()
	emails := util.Config.GetStringSlice("email.to")
	for _, email := range emails {
		wait.Add(1)
		go SendMail(html, email)
	}
	wait.Wait()
}
