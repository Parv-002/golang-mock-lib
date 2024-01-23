package database

//go:generate mockgen -destination=./mocks/db_mock.go -package=database . Db
type Db interface {
	Get() []Student
	GetById(int) Student
	Set(Student) Class
	IsPass(Student) bool
}

type Class struct {
	ClassName string
	Students  []Student
}

type Student struct {
	Name  string
	Id    int
	Grade int
}

func (c Class) Get() []Student {
	return c.Students
}
func (c Class) GetById(id int) Student {
	for _, i := range c.Students {
		if i.Id == id {
			return i
		}
	}
	return Student{}
}

func (c Class) Set(s Student) Class {
	c.Students = append(c.Students, s)
	return c
}

func (c Class) IsPass(s Student) bool {
	if s.Grade > 40 {
		return true
	}
	return false
}
