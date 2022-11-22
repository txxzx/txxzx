package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

/**
    @date: 2022/11/20
**/

func main() {
	// 爬虫的url
	url := "https://gorm.io/zh_CN/docs/"

	GetDoc2(url)

}

func GetDoc1(url string) {
	// 创建一个文档对象
	d, _ := goquery.NewDocument(url)
	// 返回一个选择器，查找某一个元素的名称或者类，id,然后进行遍历
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		// 获取连接的内容
		s2 := s.Text()
		fmt.Println(s2)
		// 取属性
		href, _ := s.Attr("href")
		fmt.Printf("href: %v", href)
	})
}

func GetDoc2(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("client do err-> %v", err)
		return
	}
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Errorf("goquery err-> %v", err)
		return
	}
	dom.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		// 获取连接的内容
		s2 := s.Text()
		fmt.Println(s2)
		// 取属性
		href, _ := s.Attr("href")
		fmt.Printf("href: %v", href)
	})
}

func GetDoc3() {
	html := `<body>
             <div id="div1">DIV1</div>
			<div class= "name">DIV2</div>
			<span>SPAN</span>
			</body>
            `
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Errorf("err-> %v", err)
	}
	// 元素的名称
	dom.Find("#div1").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "select text", s.Text())
	})
	dom.Find(".name").Each(func(i int, s *goquery.Selection) {
		fmt.Println("i", i, "select text", s.Text())
	})
}
