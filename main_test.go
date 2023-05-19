package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/uraankhayayaal/2childapp/initializers"
	"github.com/uraankhayayaal/2childapp/seed"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// func TestMe(t *testing.T) {
// 	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
// 	r := SetUpRouter()
// 	r.GET("/api/user/me", UserController.GetMe)
// 	req, _ := http.NewRequest("GET", "/api/user/me", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	responseData, _ := ioutil.ReadAll(w.Body)
// 	assert.Equal(t, mockResponse, string(responseData))
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

func TestSignUpUser(t *testing.T) {
	seed.Load(initializers.DB)

	mockResponse := `{"message":"We sent an email with a verification code to new_user@example.ru","status":"success"}`

	r := SetUpRouter()
	r.POST("/api/auth/register", AuthController.SignUpUser)

	body := []byte(`{
		"name": "Test",
		"email": "new_user@example.ru",
		"password": "00000000",
		"passwordConfirm": "00000000",
		"photo": "https://i.pravatar.cc/300?img=10"
	}`)

	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestSignInNotVerifiedUser(t *testing.T) {
	seed.Load(initializers.DB)
	r := SetUpRouter()
	r.POST("/api/auth/login", AuthController.SignInUser)

	body := []byte(`{
		"email": "not_verified_user@example.ru",
		"password": "00000000"
	}`)

	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestSignInNotExistUser(t *testing.T) {
	seed.Load(initializers.DB)
	r := SetUpRouter()
	r.POST("/api/auth/login", AuthController.SignInUser)

	body := []byte(`{
		"email": "not_exists_user@example.ru",
		"password": "00000000"
	}`)

	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignInVerifiedUser(t *testing.T) {
	seed.Load(initializers.DB)
	mockResponse := `"status":"success"`

	r := SetUpRouter()
	r.POST("/api/auth/login", AuthController.SignInUser)

	body := []byte(`{
		"email": "uraankhayayaal@yandex.ru",
		"password": "00000000"
	}`)

	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(responseData), mockResponse)
	assert.Equal(t, http.StatusOK, w.Code)
}
