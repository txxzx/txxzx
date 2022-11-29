package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

/**
    @date: 2022/11/20
**/

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	// 新建一个文档
	d, _ := goquery.NewDocument(url)
	// 查找边栏，遍历这个边栏
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		// 获取连接的内容
		link, _ := s.Attr("href")
		fmt.Println(link)
		detail_url := url + link
		d, _ := goquery.NewDocument(detail_url)
		title := d.Find(".article-titlt").Text()
		fmt.Println(title)
		content, _ := d.Find(".article").Html()
		fmt.Println(content)
		// 取属性
		href, _ := s.Attr("href")
		fmt.Printf("href: %v", href)
	})
}
