package main

/* 通过反射来修改值 */
import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflect1(b interface{}) {
	//通过反射获取到reflect.Value
	rVal := reflect.ValueOf(b)
	//看看此时的类别是什么
	fmt.Printf("rVal kind=%v\n", rVal.Kind()) //rVal kind=ptr，是一个指针类型

	//Value提供了一个方法,但Set方法是需要值类型才能调用，但是此刻却是一个指针类型
	//rVal.SetInt(5) //panic: reflect: reflect.Value.SetInt using unaddressable value
	//因此要用到Elmen方法将rVal转换为一个指向原数值地址的指针
	rVal.Elem().SetInt(5)

}

func main() {
	num := 10
	reflect1(&num)
	fmt.Println(num)
}
