package service

import (
	"fmt"
	"time"
)

func Select(c float64) string {
	var str string
	if c > 0 {
		str = fmt.Sprintf(`<p>
	Dear <strong>Owner</strong>，%s:
</p>
<p>
	<br />
</p>
<p>
	&nbsp;&nbsp;&nbsp;&nbsp;Property fees not settled...
</p>
<p>
	&nbsp;&nbsp;&nbsp;&nbsp;面积:<strong>%.2f</strong>,计费月份:<strong>%s</strong>,水费:<strong>%.2f</strong>,管理费:<strong>%.2f</strong>,未缴费合计:<strong><span style="color:#E53333;">%.2f</span></strong>
</p>
<p>
	&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://okbang.com.cn/wy-client/customer.html?source=rongshanghui&fromtype=5&feetype=0" target="_blank">点击缴费</a>
<p>
	<strong><span style="color:#E53333;"><br />
</span></strong>
</p>
<p>
	<span>Thanks.</span>
</p>
<p>
	<span>%s</span>
</p>
<p>
	<strong><span style="color:#E53333;"><br />
</span></strong>
</p>`, i.Home, i.Area, i.Year, i.Water, i.Manr, i.Count, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		str = ""
	}
	return str
}
