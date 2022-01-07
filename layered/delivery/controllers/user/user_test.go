package user

import (
	"Project/playground/be5/rest/layered/delivery/controllers/auth"
	custmiddlewares "Project/playground/be5/rest/layered/delivery/custMiddlewares"
	"Project/playground/be5/rest/layered/entities"
	"bytes"
	"errors"
	"fmt"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("UserGet", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(mockUserRepository{})
		userController.Get()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Data[0].Nama, "jerry")
		//
	})
	t.Run("Error User", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users")

		falseUserController := New(mockFalseUserRepository{})
		falseUserController.Get()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "Something wrong")
	})
}

func TestInsert(t *testing.T) {
	jwtToken := ""
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"name": "jerry",
			"hp":   "08123456789",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("Test Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "arif",
			"hp":   "08123456555",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		// fmt.Println(jwtToken)
		// tmp := ""
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		// req.Header.Set("Auhtorization", fmt.Sprintf("Bearer %v", tmp))
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(mockUserRepository{})
		if err := custmiddlewares.JWTMiddleware()(userController.Insert())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := entities.User{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Nama, response.Nama)
		assert.Equal(t, response.HP, response.HP)
	})
}

// MOCK OBJECT //

type mockAuthRepository struct{}

func (ma mockAuthRepository) Login(name, hp string) (entities.User, error) {
	return entities.User{Nama: "jerry", HP: "08123456789"}, nil
}

type mockUserRepository struct{}

func (m mockUserRepository) Get() ([]entities.User, error) {
	return []entities.User{
		{Nama: "jerry", HP: "08123456789"},
	}, nil
}

func (m mockUserRepository) Insert(entities.User) (entities.User, error) {
	return entities.User{Nama: "jerry", HP: "08123456789"}, nil
}

type mockFalseUserRepository struct{}

func (m mockFalseUserRepository) Get() ([]entities.User, error) {
	return nil, errors.New("False User Object")
}

func (m mockFalseUserRepository) Insert(entities.User) (entities.User, error) {
	return entities.User{Nama: "jerry", HP: "08123456789"}, errors.New("False Login Object")
}
