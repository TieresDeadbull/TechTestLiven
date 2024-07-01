package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type SafeUser interface {
	PreparePrepare(option string, security security.Security)
}

// Esta estrutura define um usuario com apenas alguns campos mais genericos
type User struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json: "name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Passphrase  string    `json:"pass, omitempty"`
	PhoneNumber string    `json:"phone,omitempty"`
	CreatedAt   time.Time `json:"created_at, omitempty"`
}

// Função que por meio das funções validate e prepare verifica
// o preenchimento correto dos campos do usuário recebido
func (user *User) Prepare(option string, security security.Security) error {
	if err := user.validate(option); err != nil {
		return err
	}
	if err := user.formatting(option, security); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(option string) error {
	if user.Name == "" {
		return errors.New("obrigatório preenchimento do campo nome")
	}
	if user.Email == "" {
		return errors.New("obrigatório preenchimento do campo email")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email inserido em formato inválido")
	}

	if option == "register" && user.Passphrase == "" {
		return errors.New("obrigatório preenchimento do campo senha")
	}
	if user.PhoneNumber == "" {
		return errors.New("obrigatório preenchimento do campo telefone")
	}

	return nil
}

func (user *User) formatting(option string, security security.Security) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.PhoneNumber = strings.TrimSpace(user.PhoneNumber)

	if option == "register" {
		hashedPass, err := security.Hash(user.Passphrase)
		if err != nil {
			return err
		}
		user.Passphrase = string(hashedPass)
	}
	return nil
}
