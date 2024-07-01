package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Auth interface {
	CreateToken(userID uint64) (string, error)
	ValidateToken(r *http.Request) error
	ExtractUserID(r *http.Request) (uint64, error)
}

type JWTAuth struct{}

// Funçao que retorna um token assinado com as permissoes do usuário
func (a *JWTAuth) CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["expire"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// Faz o split no header de modo a remover o Bearer
// caso esteja em um formato diferente de "Bearer xyzaasahsa..."
// retorna string vazia, dessa forma invalidando o token na função ValidateToken
func (a *JWTAuth) extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	//token format =>  Bearer anything...
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func (a *JWTAuth) verificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// Verifica se o token na request é valido
func (a *JWTAuth) ValidateToken(r *http.Request) error {
	tokenString := a.extractToken(r)
	token, err := jwt.Parse(tokenString, a.verificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

//UserFunctions

// Extrai o ID salvo no token
func (a *JWTAuth) ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := a.extractToken(r)
	token, err := jwt.Parse(tokenString, a.verificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}

	return 0, errors.New("token inválido")

}
