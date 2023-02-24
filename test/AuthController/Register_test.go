package authcontroller_test

import (
	"net/http"
	authcontroller "setup/controllers/AuthController"
	"setup/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	testinitializer.Starter()

	json := `{"email":"akifkadioglu@yaani.com","name":"akif","password":"deneme1","password_confirmation":"deneme1","username":"akif"}`

	c, rec := testinitializer.Method(http.MethodPost, json)

	if assert.NoError(t, authcontroller.Register(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
