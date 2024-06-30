package security

import "golang.org/x/crypto/bcrypt"

//Aplicando Hash a senha recebida como string a fim de aumentar
//a segurança ao armazenar esse dado sensível no banco
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func VerifyPass(hashedPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
