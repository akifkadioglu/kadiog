package authcontroller_test

import (
	"net/http"
	"reflect"
	authcontroller "setup/controllers/AuthController"
	"setup/database"
	"setup/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	
	//Cant find the .env file now, I will deal with it later
	database.Init()
	c, rec := testinitializer.Method(http.MethodPost)
	c.SetPath("/login")

	// Assertions
	if assert.NoError(t, authcontroller.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, reflect.TypeOf("token"), reflect.TypeOf(rec.Body.String()))
	}
}
