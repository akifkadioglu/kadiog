package indexcontroller_test

import (
	"net/http"
	indexcontroller "setup/controllers/IndexController"
	"setup/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	testinitializer.Starter()

	c, rec := testinitializer.Method(http.MethodGet, "")

	if assert.NoError(t, indexcontroller.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
