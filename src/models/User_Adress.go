package models

type UserAddresses struct {
	User      User      `json:"user"`
	Addresses []Address `json:"addresses"`
}
