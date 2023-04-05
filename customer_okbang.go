package main

import (
    "encoding/json"
    "fmt"
    "okbang/util"
    "strings"
    "time"
)

type result struct {
    JFYFstr   string `json:"JFYFstr"`
    FH        string `json:"FH"`
    HTID      string `json:"HTID"`
    URLString string `json:"urlString"`
    WYMC      string `json:"WYMC"`
    Valid     bool   `json:"valid"`
    Ads       []struct {
        Createtime int64  `json:"createtime"`
        Descr      string `json:"descr"`
        Href       string `json:"href"`
        ID         string `json:"id"`
        ImgSrc     string `json:"imgSrc"`
        Source     string `json:"source"`
        Status     string `json:"status"`
        Updatetime int64  `json:"updatetime"`
        WYID       string `json:"wYID"`
    } `json:"ads"`
    Feelist []struct {
        A1   string `json:"项目名称"`
        A2   string `json:"欠收应收款"`
        A3   string `json:"欠收滞纳金"`
        A4   string `json:"计费编号"`
        A5   string `json:"月费用"`
        A6   string `json:"本次读数"`
        A7   string `json:"欠收往月"`
        HTID string `json:"HTID"`
        A8   string `json:"本月应收款"`
        A9   string `json:"净用量"`
        A10  string `json:"计费用量"`
        A11  string `json:"标准价"`
        A12  string `json:"计费用量1"`
        A21  string `json:"项目编号"`
        BZID string `json:"BZID"`
        A13  string `json:"月初滞纳金"`
        A14  string `json:"时段单位"`
        A15  string `json:"用量单位"`
        A16  string `json:"上次读数"`
        A17  string `json:"走表项目"`
        A18  string `json:"欠收合计"`
        A19  string `json:"计费月份"`
        YSID string `json:"YSID"`
        A20  string `json:"是否付清"`
    } `json:"feelist"`
    Phone   string `json:"phone"`
    JFYF2   string `json:"JFYF2"`
    Addr    string `json:"addr"`
    Account struct {
        AppSecret      string `json:"appSecret"`
        Appid          string `json:"appid"`
        Appuserid      string `json:"appuserid"`
        AppvKey        string `json:"appv_key"`
        BeforehandType int    `json:"beforehandType"`
        ID             string `json:"id"`
        LTID           string `json:"lTID"`
        Offline        bool   `json:"offline"`
        PlatpKey       string `json:"platp_key"`
        Platsystem     string `json:"platsystem"`
        Property       string `json:"property"`
        Source         string `json:"source"`
        Status         string `json:"status"`
        SubAppid       string `json:"subAppid"`
        Sumflag        int    `json:"sumflag"`
        WXID           string `json:"wXID"`
        WYID           string `json:"wYID"`
        Wxh5           int    `json:"wxh5"`
    } `json:"account"`
    DLMC     string `json:"DLMC"`
    TSZT     string `json:"TSZT"`
    Username string `json:"username"`
    Jzmj     string `json:"jzmj"`
}

func okbang() string {
    var arr, tt []string
    var u string
    url := fmt.Sprintf(`http://okbang.com.cn/wy-server/htmlAction/doAction.htm?feetype=0&WYID=95&HTID=32870&JFYF=%s&SOURCE=rongshanghui&SIGN=%s`, time.Now().Format("200601"), util.Config.GetString("okbang.sign"))
    r := Post("POST", url, nil)
    var b result
    err := json.Unmarshal(r, &b)
    if err != nil {
        fmt.Println(err.Error())
        return ""
    }

    for _, s := range b.Feelist {
        if s.A20 == "true" {
            u = "付清"
        } else {
            url := fmt.Sprintf(`<p><a href="http://okbang.com.cn/wy-client/customer.html?source=rongshanghui&fromtype=5&feetype=0" target="_blank"><strong>点击缴费</strong><strong></strong></a></p>`)
            u = "未付清<br>" + url
        }
        tt = append(tt, fmt.Sprintf(`<strong>%s</strong>：<strong><span style="color:#006600;">%s</span></strong>`, s.A1, u))
        arr = append(arr, fmt.Sprintf(`
  <table style="width:100%%;" cellpadding="1" cellspacing="0" border="1" bordercolor="#000000"> 
   	<tbody>
<br>
<tr><td>项目名称	</td><td><strong><span style='color:#4C33E5;'>%s</span></strong></td></tr>
<tr><td>欠收应收款	</td><td><strong>%s</strong></td></tr>
<tr><td>欠收滞纳金	</td><td><strong>%s</strong></td></tr>
<tr><td>计费编号	</td><td><strong>%s</strong></td></tr>
<tr><td>月费用		</td><td><strong>%s</strong></td></tr>
<tr><td>本次读数	</td><td><strong>%s</strong></td></tr>
<tr><td>欠收往月	</td><td><strong>%s</strong></td></tr>
<tr><td>HTID		</td><td><strong>%s</strong></td></tr>
<tr><td>本月应收款	</td><td><strong>%s</strong></td></tr>
<tr><td>净用量		</td><td><strong>%s</strong></td></tr>
<tr><td>计费用量	</td><td><strong>%s</strong></td></tr>
<tr><td>标准价		</td><td><strong>%s</strong></td></tr>
<tr><td>计费用量1	</td><td><strong>%s</strong></td></tr>
<tr><td>项目编号	</td><td><strong>%s</strong></td></tr>
<tr><td>BZID		</td><td><strong>%s</strong></td></tr>
<tr><td>月初滞纳金	</td><td><strong>%s</strong></td></tr>
<tr><td>时段单位	</td><td><strong>%s</strong></td></tr>
<tr><td>用量单位	</td><td><strong>%s</strong></td></tr>
<tr><td>上次读数	</td><td><strong>%s</strong></td></tr>
<tr><td>走表项目	</td><td><strong>%s</strong></td></tr>
<tr><td>欠收合计	</td><td><strong>%s</strong></td></tr>
<tr><td>计费月份	</td><td><strong>%s</strong></td></tr>
<tr><td>YSID		</td><td><strong>%s</strong></td></tr>
<tr><td>是否付清	</td><td><strong>%s</strong></td></tr>
   	</tbody> 
  </table> 
`, s.A1, s.A2, s.A3, s.A4, s.A5, s.A6, s.A7, s.HTID, s.A8, s.A9, s.A10, s.A11, s.A12, s.A21, s.BZID, s.A13, s.A14, s.A15, s.A16, s.A17, s.A18, s.A19, s.YSID, s.A20))
    }
    htm := fmt.Sprintf(`<html>
 <head></head>
 <body>
		%s
 		%s
  <br/> 
  <span id="__kindeditor_bookmark_start_15__"></span>
 </body>
</html>`, strings.Join(tt, "<br>"), strings.Join(arr, " "))
    return htm
}
