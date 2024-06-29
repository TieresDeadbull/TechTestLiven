package models

import "gorm.io/gorm"

//Esta estrutura define um endereço com apenas alguns campos mais genericos
type Address struct {
	gorm.Model
	Street  string `json:"street"`
	ZipCode string `json:"zipcode"`
	City    string `json:"city"`
}
