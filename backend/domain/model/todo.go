// package models

package model

type Todo struct {
	ID   uint   `gorm:"primaryKey"`
	Text string `gorm:"not null"`
	Done bool   `gorm:"default:false"`
}
