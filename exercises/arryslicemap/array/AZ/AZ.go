package main
import"fmt"
func main(){

var mychars [26]byte//声明数组

for i:=0;i<26;i++{
mychars[i]='A'+byte(i)//i是int类型，必须使用类型转换才能做运算

}
for k,v:=range mychars{//注意for range遍历的用法，
fmt.Printf("第%d个元素是%c\n",k+1,v)

}
}

