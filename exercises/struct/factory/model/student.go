package model

type student struct { //大写代表能够被其他包引用
	Name  string
	score float64
}

//如果改成小写的话该如何实现跨包引用？？？
//改成小写说明是私有的，只能在model包内使用

func Newstudent(n string, s float64) *student { //返回指针
	return &student{ //创建了一个student的实例，返回指针，在其他地方使用的时候就可以直接访问其本身
		Name:  n,
		score: s,
	}
}
func (s *student) Getscore() float64 {
	return s.score

}
