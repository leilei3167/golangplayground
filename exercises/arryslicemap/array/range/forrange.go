package main
import "fmt"
func main(){
/*用range来遍历，语法
for index value:=range array01{
...
}
index value分别代表下标和对应的值，可自定义名称
*/
arr1:=[...]int{1,45,23,12,56,66}
for i,v:=range arr1{//i,v分别对应下标和值，如果不需要可以用_代替
fmt.Printf("i=%d v=%d\n",i,v)

}


}