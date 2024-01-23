package database

import (
	_ "go.uber.org/mock/mockgen/model"
)

//go:generate mockgen -destination=./mocks/db_mock.go -package=database . Db
