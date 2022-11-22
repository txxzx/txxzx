package main

/**
    @date: 2022/11/20
**/

// selection 可以提取 添加删除元素，属性内容
// 内置函数
// eq(index int) *selection  根据索引获取某个元素
// First() *selection 获取第一个节点集
// Last() *selection 获取最后一个子节点集
// Next（）*selection  获取下一个兄弟节点
// NextAll()  *selection  获取最后所有的兄弟节点集
// Prev（）*selection // 前一个兄弟节点集
// Get(index int) *html.Node  根据索引获取一个节点
// Index() int 返回选择对象的第一个元素的位置
// Slice(start ,end int) *selection 根据起始位置获取子节点集

// 遍历选择的节点
// Each(f func(i int,s *selection)) *selection // 遍历
// EachWithBreak(f func(int,*selection) bool) *selection 可中断遍历
// Map(f func(int,*selection) string) [result []string] 返回字符串数组

// 检测获取节点的属性
//Attr() RemoveAttr(),SetAttr() 获取,移出 设置属性
// AddClass(),HashClass() remove
