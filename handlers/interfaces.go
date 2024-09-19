package handlers

import "gorm.io/gorm"

type DBInterface interface {
	Create(value interface{}) *gorm.DB
}

type ErrorResponse struct {
	Error string `json:"error"`
}
