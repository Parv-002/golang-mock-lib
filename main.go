package main

import (
	"example.com/m/database"
	"fmt"
)

func main() {
	var c database.Class
	c = c.Set(database.Student{
		Name:  "Parv",
		Id:    1,
		Grade: 55,
	})
	c = c.Set(database.Student{
		Name:  "Atharva",
		Id:    2,
		Grade: 75,
	})
	c = c.Set(database.Student{
		Name:  "Varad",
		Id:    3,
		Grade: 85,
	})

	fmt.Printf("Students : \n %+v \n", c.Get())
	fmt.Printf("Student at id - %d : \n %+v \n", 1, c.GetById(1))
	if c.IsPass(c.GetById(1)) {
		fmt.Printf("Student at id - %d is pass", 1)
	} else {
		fmt.Printf("Student at id - %d is fail", 1)
	}
}

func GetAllStudents(db database.Db) []database.Student {
	return db.Get()
}

func GetStudentDetails(db database.Db, id int) database.Student {
	return db.GetById(id)
}

func CheckStudentResult(db database.Db, s database.Student) bool {
	return db.IsPass(s)
}

func GetStudentDetailsByIds(db database.Db, ids ...int) []database.Student {
	var students []database.Student
	for _, i := range ids {
		students = append(students, db.GetById(i))
	}
	return students
}
