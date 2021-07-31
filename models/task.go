package models

type Todo struct {
	// gorm.Model
	ID        uint
	Title     string
	Body      string
	Completed bool
}
