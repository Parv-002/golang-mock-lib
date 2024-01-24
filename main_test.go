package main

import (
	"example.com/m/database"
	databaseMock "example.com/m/database/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This will check if dbMock receives Get() call
func Test_Database_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbMock := databaseMock.NewMockDb(ctrl)
	// This will mock the IsPass function's return
	dbMock.EXPECT().Get().Return([]database.Student{
		{Name: "Parv", Id: 1, Grade: 70},
		{Name: "Atharva", Id: 2, Grade: 50},
	})
	t.Run("database get all students", func(t *testing.T) {
		students := GetAllStudents(dbMock)
		assert.NotNil(t, students)
	})
}

// This will check dbMock receives IsPass() call
func Test_Database_IsPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	// ctrl.Finish() is optional in newer version of go.
	defer ctrl.Finish()
	dbMock := databaseMock.NewMockDb(ctrl)
	// This will mock the IsPass function's definition as well as return
	dbMock.EXPECT().IsPass(gomock.Any()).DoAndReturn(func(s database.Student) bool {
		if s.Grade > 55 {
			return true
		}
		return false
	})
	t.Run("database check result", func(t *testing.T) {
		res := CheckStudentResult(dbMock, database.Student{Name: "Parv", Id: 1, Grade: 70})
		assert.Equal(t, res, true)
	})
}

// This will check dbMock receives GetById() call
func Test_Database_Get_Student(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := databaseMock.NewMockDb(ctrl)

	// This will mock function GetById and will also check if this function gets called only once
	dbMock.EXPECT().GetById(gomock.Any()).DoAndReturn(func(id int) database.Student {
		return database.Student{
			Name:  "Parv",
			Id:    1,
			Grade: 65,
		}
	}).Times(1)
	t.Run("database get student by id", func(t *testing.T) {
		res := GetStudentDetails(dbMock, 1)
		assert.NotNil(t, res)
	})
}

// This will check dbMock receives GetById() call
func Test_Database_Get_Student_Ids(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := databaseMock.NewMockDb(ctrl)

	// This will mock function GetById and will also check if this function gets called 3 times
	dbMock.EXPECT().GetById(gomock.Any()).DoAndReturn(func(id int) database.Student {
		return database.Student{
			Name:  "Parv",
			Id:    1,
			Grade: 65,
		}
	}).Times(3)
	t.Run("database get student by id", func(t *testing.T) {
		res := GetStudentDetailsByIds(dbMock, 1, 2, 3)
		assert.NotNil(t, res)
	})
}
