package database

import (
	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -destination=./mocks/db_mock.go -package=databaseMock . Db
