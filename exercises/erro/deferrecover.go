package main
import ("fmt")


func test(){
	//可以使用 defer recover来捕获及处理异常
	defer func(){//匿名函数的引用就在后面加（）
    err:=recover()//recover为内置函数，用于捕获异常，必须和defer配合使用
      if err != nil {
    fmt.Println("err=",err)
      }
	}()//表示匿名函数的调用
	//此个函数体内还可增加功能，如增加报错预警

num1:=10
num2:=0
res:=num1/num2//除以0一定会错误
fmt.Println(res)
}

func main(){

test()
fmt.Println("看见这个说明成功recover了那个panic")

}