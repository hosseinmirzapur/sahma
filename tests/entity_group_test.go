package tests

import (
	"sahma/internal/database/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAttributes(t *testing.T) {
	var eg models.EntityGroup

	mapItems, err := eg.GetAttributes()
	assert.Nil(t, err)

	expectedType := make(map[string]interface{}, 0)
	assert.IsType(t, expectedType, mapItems)
}
