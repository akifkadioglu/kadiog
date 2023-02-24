package authcontroller_test

import (
	"net/http"
	authcontroller "setup/controllers/AuthController"
	"setup/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	testinitializer.Starter()
	
	json := `{"email":"akifkadioglu@yaani.com","password":"deneme1"}`

	c, rec := testinitializer.Method(http.MethodPost, json)

	if assert.NoError(t, authcontroller.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
