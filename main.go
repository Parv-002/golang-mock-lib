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
