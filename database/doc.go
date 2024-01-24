package database

import (
	_ "github.com/golang/mock/mockgen/model"
	_ "go.uber.org/mock/mockgen/model"
)

//go:generate mockgen -destination=./mocks/db_mock.go -package=database . Db
