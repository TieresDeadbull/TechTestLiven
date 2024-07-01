package security

import "golang.org/x/crypto/bcrypt"

type Security interface {
	Hash(pass string) ([]byte, error)
	VerifyPass(hashedPass string, pass string) error
}

type Encrypted struct{}

//Aplicando Hash a senha recebida como string a fim de aumentar
//a segurança ao armazenar esse dado sensível no banco
func (e *Encrypted) Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func (e *Encrypted) VerifyPass(hashedPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
