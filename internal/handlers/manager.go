package handlers

import "gorm.io/gorm"

type manager struct{}

func NewManager(db *gorm.DB) *manager {
	m := &manager{}
	return m
}
