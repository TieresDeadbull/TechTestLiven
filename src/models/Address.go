package models

import (
	"errors"
	"strings"
	"time"
)

// Esta estrutura define um endereço com apenas alguns campos mais genericos
type Address struct {
	ID        uint64    `json:"id,omitempty"`
	Street    string    `json:"street"`
	ZipCode   string    `json:"zipcode"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at, omitempty"`
}

// Função que por meio das funções validate e formatting verifica
// o preenchimento correto dos campos do endereço recebido
func (address *Address) Prepare() error {
	if err := address.validate(); err != nil {
		return err
	}
	if err := address.formatting(); err != nil {
		return err
	}

	return nil
}

func (address *Address) validate() error {

	if address.Street == "" {
		return errors.New("obrigatório preenchimento do campo street")
	}
	if address.ZipCode == "" {
		return errors.New("obrigatório preenchimento do campo zipcode")
	}
	if address.City == "" {
		return errors.New("obrigatório preenchimento do campo city")
	}

	return nil
}

func (address *Address) formatting() error {
	address.Street = strings.TrimSpace(address.Street)
	address.ZipCode = strings.TrimSpace(address.ZipCode)
	address.City = strings.TrimSpace(address.City)

	return nil
}
