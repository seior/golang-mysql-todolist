package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConnection(t *testing.T) {
	db := GetDatabase()

	assert.NotNil(t, db)
}