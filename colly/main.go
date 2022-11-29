package main

import (
	"fmt"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"github.com/henrylee2cn/goutil/coarsetime"
)

/**
    @date: 2022/11/27
**/
func main() {
	var (
		title, content string
	)
	Init()
	//	cooike := "_ga=GA1.2.497622810.1658154826; __gads=ID=054a197529a927b1:T=1658154828:S=ALNI_MYw98poBtgSobNbNZsCKwqvYdeFQw; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8NfDHj8mnYFAmPyhfXwJojec2zu4sR8k2BMjGmi9X86GWMc8ItbfOPrpqh3TRXjsoaZI_qYPHSLOQS4gVShVfzI_6zhhU-5OiFdLTeDOr-xaz6QYIRMZDTyTlM7mesdDNDb49Da5Yhbvs5JuQxuc7Qg; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1667638049,1667664830,1668863125; _gid=GA1.2.2065310825.1668863125; __gpi=UID=000007d6d7223f6d:T=1658154828:RT=1668863124:S=ALNI_MaNjWTdNaXApPSyZ4rqN5vzo6viaQ; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1668873196"
	// 创建连接器
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	// 当发生请求方法被调用
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url:", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("发生错误调用：OnError")
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("获得响应调用：OnResponse")
	})
	// 抓取html页面方法调用  goquery 选择器可以是 id name class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "index.html" {
			c.Visit(e.Request.AbsoluteURL(href))
		}
		//e.Request.Visit(e.Attr("href"))
	})
	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
		title := h.Text
		fmt.Println("Text:\n", title)
		//e.Request.Visit(e.Attr("href"))
	})
	c.OnHTML(".article", func(h *colly.HTMLElement) {
		content, _ := h.DOM.Html()
		fmt.Println("content:\n", content)
	})
	SaveToDB(title, content)
	// 完成
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("结束:", r.Request.URL)
	})
	// 放在回调函数之后
	c.Visit("https://gorm.io/zh_CN/docs/")
}

var engine *xorm.Engine
var err error

func Init() {
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		fmt.Printf("err-> %v", err)
	} else {
		err := engine.Ping()
		if err != nil {
			fmt.Printf("err-> %v", err)
		} else {
			fmt.Printf("连接成功")
		}
	}
}

type GormPage struct {
	Id      int64
	Title   string
	Content string `xorm:"text"`
	Created int64  `xorm:"created"`
	Updated int64  `xorm:"updated"`
}

// 保存到数据库
func SaveToDB(title, content string) {
	engine.Sync(new(GormPage))
	page := GormPage{
		Title:   title,
		Content: content,
		Created: coarsetime.FloorTimeNow().Unix(),
		Updated: coarsetime.FloorTimeNow().Unix(),
	}
	affected, err := engine.Insert(&page)
	if err != nil {
		fmt.Printf("err-> %v", err)
	}
	fmt.Println("save:" + string(affected))
}
