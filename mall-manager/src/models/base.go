package models

import "time"

type Base struct {

	Id int64

	CreateUser string

	CreateTime time.Time

	UpdateUser string

	UpdateTime time.Time

}
