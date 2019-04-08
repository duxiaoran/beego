package models

import "time"

type Role struct {
	Id    int
	Name  string
	Desc  string
	Rule string
	Created time.Time
}

func (m *Role) TableName() string {
	return TableName("role")
}