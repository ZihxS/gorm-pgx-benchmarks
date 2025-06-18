package bench

import (
	"context"
)

var ctx = context.Background()

type Model struct {
	Id      int `orm:"auto" gorm:"primary_key" db:"id" bun:",pk,autoincrement"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func (m *Model) TableName() string {
	return "models"
}

func (m *Model) Table() string {
	return "models"
}

func NewModel() *Model {
	m := new(Model)
	m.Name = "GORM vs PGX Benchmarks (Column Name)"
	m.Title = "GORM vs PGX Benchmarks (Column Title)"
	m.Fax = "12345678"
	m.Web = "https://alwaysngoding.com"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}
