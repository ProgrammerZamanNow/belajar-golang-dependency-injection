package test

import (
	"github.com/stretchr/testify/assert"
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
