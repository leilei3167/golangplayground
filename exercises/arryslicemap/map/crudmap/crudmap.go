package main

import "fmt"

func main() {
	/* map的增删改查 */
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println(cities)
	//增改：如果没有key存在 就是增加，有就是更新 如
	cities["no3"] = "shanghai~~~~"
	cities["no4"] = "chongqing~~~~"
	fmt.Println(cities)
	//删
	delete(cities, "no1")
	fmt.Println(cities)
	delete(cities, "no5") //删除不存在的key，不会报错，go中没有专门的方法一次删除完
	fmt.Println(cities)
	/* 如果要一次性删除所有key：
	   1，遍历删除
	   2，直接make一个新空间
	   如：
	   cities=make(map[string]string)
	   原来的内容就会被GC回收
	*/

	//查
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有这个key，值为%v", val)
	} else {
		fmt.Printf("没有这个key")

	}

}
