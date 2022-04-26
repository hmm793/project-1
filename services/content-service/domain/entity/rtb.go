package entity

import "time"

type RTBEntity struct {
	ID            int
	Title         string
	Description   string
	FileName      string
	UpdatedAt     time.Time
	UpdatedById   int
	UpdatedByName string
}