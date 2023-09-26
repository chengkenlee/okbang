/*
 *@author ChengKen
 *@date   16/03/2023 14:04
 */
package main

import (
	"okbang/service"
	"okbang/util"
)

func main() {
	util.Parm()
	util.Loggers()
	service.Crontab()
}
