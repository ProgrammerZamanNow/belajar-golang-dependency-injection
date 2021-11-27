package test

import (
	"github.com/stretchr/testify/assert"
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databaes")
	assert.NotNil(t, connection)

	cleanup()
}
