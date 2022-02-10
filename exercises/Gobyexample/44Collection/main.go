package main

//组合函数
/* 我们经常需要程序在数据集上执行操作，
比如选择满足给定条件的所有项，或者将所有的项通过一个自定义函数映射到一个新的集合上。

在某些语言中，会习惯使用泛型。 Go 不支持泛型，
在 Go 中，当你的程序或者数据类型需要 时，通常是通过组合的方式来提供操作函数。
*/

//返回目标字符串t的索引值,没有返回-1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}

	}
	return -1
}

//如果t在vs切片中返回true,和Index组合
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0 //非负代表有这个值

}

//如果切片中有一个字符串满足条件 f,则返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}

	return false
}

//如果切片中所有字符串都满足,则返回true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

// 返回一个包含所有切片中满足条件 `f` 的字符串的新切片。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个对原始切片中所有字符串执行函数 `f` 后的新切片。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
