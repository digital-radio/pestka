package tests_test

import (
	. "github.com/digital-radio/pestka/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRoot(t *testing.T) {
	// given
	app := TestApp{}

	// when
	w := app.MakeRequest("GET", "/", nil)

	// then
	a := assert.New(t)
	a.Equal(404, w.Code, "Invalid status code")
	a.Equal(`{"message":"not found"}`, w.Body.String())
}
