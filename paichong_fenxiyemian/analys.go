package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

/**
    @date: 2022/11/20
**/

// 利用正则分析页面

func Parse(html string) {
	// 解析页面
	// 替换空格
	html = strings.Replace(html, "\n", "", -1)
	// 边框栏内容块 任意匹配
	re_sider := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	// 找到边框内容块
	sidebar := re_sider.FindString(html)
	// 连接转正则
	re_link := regexp.MustCompile(`href="(.*?)"`)
	// 找到所有的连接
	links := re_link.FindAllString(sidebar, -1)
	base_url := "https://gorm.io/zh_CN/docs/"
	for _, v := range links {
		fmt.Println("url: %v\n", v)
		s := v[6 : len(v)-1]
		url := base_url + s
		fmt.Printf("url: %v\n", url)
		//body := Fech(url)
		//fmt.Println(body)
		// 启动另外一个线程处理
	}
}

func Fech(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 设置请求header和cookie
	// 解决防爬虫 1.用户代理  2.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "_ga=GA1.2.497622810.1658154826; __gads=ID=054a197529a927b1:T=1658154828:S=ALNI_MYw98poBtgSobNbNZsCKwqvYdeFQw; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8NfDHj8mnYFAmPyhfXwJojec2zu4sR8k2BMjGmi9X86GWMc8ItbfOPrpqh3TRXjsoaZI_qYPHSLOQS4gVShVfzI_6zhhU-5OiFdLTeDOr-xaz6QYIRMZDTyTlM7mesdDNDb49Da5Yhbvs5JuQxuc7Qg; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1667638049,1667664830,1668863125; _gid=GA1.2.2065310825.1668863125; __gpi=UID=000007d6d7223f6d:T=1658154828:RT=1668863124:S=ALNI_MaNjWTdNaXApPSyZ4rqN5vzo6viaQ; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1668873196")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get req err", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read err", err)
		return ""
	}
	return string(body)
}

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	s := Fech(url)
	//fmt.Printf("s:%v\n",s)
	Parse(s)
}
