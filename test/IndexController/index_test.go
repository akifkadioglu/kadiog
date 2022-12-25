package indexcontroller_test

import (
	"net/http"
	"reflect"
	indexcontroller "setup/controllers/IndexController"
	"setup/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	c, rec := testinitializer.Method(http.MethodGet)

	// Assertions
	if assert.NoError(t, indexcontroller.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, reflect.TypeOf("welcome to index"), reflect.TypeOf(rec.Body.String()))
	}
}
