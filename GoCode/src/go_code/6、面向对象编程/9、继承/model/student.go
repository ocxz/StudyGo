package model
import (
	"fmt"
)

type Pupil struct {
	Name string
	Age int
	Score int
}

func (pupil *Pupil) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", pupil.Name, pupil.Age, pupil.Score)
}

func (p *Pupil) SetScore(score int) {
	p.Score = score
}

func (p *Pupil) Testing() {
	fmt.Println("小学生正在考试")
}