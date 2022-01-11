package main
/* map中的数据是无序的,go中也没有提供直接对map排序的方法,每次遍历出来的结果都可能
不一样 */
import (
	"fmt"
	"sort"
)
func main()  {
	map1:=make(map[int]int)
	map1[10]=100
	map1[1]=13
	map1[4]=56
	map1[8]=90
fmt.Println(map1)

var keys []int

for k := range map1 {//map并取出key放入切片
keys=append(keys,k)//通过append放入keys值

}
sort.Ints(keys)//对取出的keys进行递增排序
fmt.Println(keys)

for _,k:=range keys{//遍历这个切片
fmt.Printf("map1[%v]=%v\n",k,map1[k])//打印值key对应的map1中的值
}
}