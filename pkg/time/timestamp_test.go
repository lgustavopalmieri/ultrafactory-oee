package location_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTimestampFromTimezone(t *testing.T) {
	timezone := "America/Sao_Paulo"
	timestamp, err := CreateTimestampFromTimezone(timezone)

	assert.NotNil(t, timestamp)
	assert.Nil(t, err)
	assert.IsType(t, "string", *timestamp)

	wrongTimezone := "AnyPlace/In_The_World"
	wrongTimestamp, err := CreateTimestampFromTimezone(wrongTimezone)

	assert.Nil(t, wrongTimestamp)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unknown time zone AnyPlace/In_The_World")
}

// go test -v -run ^TestCreateTimestampFromTimezone$
