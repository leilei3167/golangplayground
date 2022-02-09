package main

import (
	"fmt"
	"sort"
)

/* 有时候我们想使用和集合的自然排序不同的方法对集合进行排序。
例如，我们想按照字母的长度而不是首字母顺序对字符串排序。
这里是一个 Go 自定义排序的例子。 */

//创建一个自定义类型
type ByLength []string //因为内置类型不能指定方法

//只要实现sort.Interface里的 `Len`，`Less` 和 `Swap` 方法，
//即可调用通用的sort方法

/* type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
} */

func (s ByLength) Len() int {
	return len(s)

}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]

}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])//自定义按照字符串长度排序

}

func main() {

	slice := []string{"ds21312321312a", "12321", "21321dsa"}
	sort.Sort(ByLength(slice))
	fmt.Println(slice)
}

/* 类似的，参照这个创建一个自定义类型的方法，实现这个类型的 这三个接口方法，
然后在一个这个自定义类型的集合上调用 sort.Sort 方法，
我们就可以使用任意的函数来排序 Go 切片了。 */
