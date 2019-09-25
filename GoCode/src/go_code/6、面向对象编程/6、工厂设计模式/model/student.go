package model

type student struct {
	Name string
	Age int
	score float64
}

func (stu *student) SetScore(score float64) {
	stu.score = score
}
func (stu *student) GetScore() float64 {
	return stu.score
}

// 因为student结构体名首字母小写，私有，仅能在本包中使用
// 可设计一个函数，通过传入的参数，创建对象返回
// 返回指向结构体指针，通过指针可以操作新创建的结构体对象
func CreateStudent(name string, age int, score float64) *student{

	return &student{ name, age, score}
}