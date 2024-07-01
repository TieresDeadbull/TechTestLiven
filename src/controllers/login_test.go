// login_test.go

package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginController_Login(t *testing.T) {
	// Configuração dos mocks ou implementações simuladas
	mockDBConnector := &db.MSQL{}
	mockUserRepo := &repositories.UsersRepo{}
	db, err := mockDBConnector.Connect()
	if err != nil {
		t.Fatal(err)
	}
	mockUserRepo.SetDB(db)
	mockSecurity := &security.Encrypted{}
	mockAuth := &auth.JWTAuth{}

	// Criando instância do LoginController com os mocks
	lc := LoginController{
		DB:       mockDBConnector,
		UserRepo: mockUserRepo,
		Security: mockSecurity,
		Auth:     mockAuth,
	}

	// Dados do usuário para o teste
	user := models.User{
		Email:      "test@example.com",
		Passphrase: "password",
	}

	// Convertendo o usuário em JSON
	userJSON, _ := json.Marshal(user)

	// Criando uma requisição HTTP simulada com os dados do usuário
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Gravando o corpo da requisição
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(lc.LoginUser)
	handler.ServeHTTP(rr, req)

	// Verificando o código de status HTTP retornado
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verificando o corpo da resposta
	expected := "Você está logado."
	assert.Equal(t, expected, rr.Body.String(), "expected response body to be %s, got %s", expected, rr.Body.String())
}
