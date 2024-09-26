package time_helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestCreateTimestampFromTimezone(t *testing.T) {
	th := TimeHelper{Timezone: "America/Sao_Paulo"}
	timestamp, err := th.CreateTimestampFromTimezone()

	assert.NotNil(t, timestamp, "Timestamp should not be nil")
	assert.Nil(t, err, "Expected no error when creating timestamp with a valid timezone")
	assert.IsType(t, time.Time{}, *timestamp, "Expected the returned type to be time.Time")

	// Test with an invalid timezone
	wrongTimezone := TimeHelper{Timezone: "AnyPlace/In_The_World"}
	wrongTimestamp, err := wrongTimezone.CreateTimestampFromTimezone()

	assert.Nil(t, wrongTimestamp, "Timestamp should be nil for an invalid timezone")
	assert.NotNil(t, err, "An error should be returned for an invalid timezone")
}

func TestCreateAndFormatTimestampFromTimezone(t *testing.T) {
	th := TimeHelper{Timezone: "America/Sao_Paulo"}
	timestamp, err := th.CreateAndFormatTimestampFromTimezone()

	assert.NotNil(t, timestamp, "Formatted timestamp should not be nil")
	assert.Nil(t, err, "Expected no error when creating and formatting timestamp")
	assert.IsType(t, "string", *timestamp, "Expected the returned type to be string")

	// Test with an invalid timezone
	wrongTimezone := TimeHelper{Timezone: "AnyPlace/In_The_World"}
	wrongTimestamp, err := wrongTimezone.CreateAndFormatTimestampFromTimezone()

	assert.Nil(t, wrongTimestamp, "Formatted timestamp should be nil for an invalid timezone")
	assert.NotNil(t, err, "An error should be returned for an invalid timezone")
}

func TestFormatTimestamp(t *testing.T) {
	th := TimeHelper{Timezone: "America/Sao_Paulo"}
	timestampInTime := time.Now()

	formattedTimestamp := th.FormatTimestamp(timestampInTime)

	assert.NotNil(t, formattedTimestamp, "Formatted timestamp should not be nil")
	assert.IsType(t, "string", formattedTimestamp, "Expected the returned type to be string")
}

func TestParseTimestampToTime(t *testing.T) {
	th := TimeHelper{Timezone: "America/Sao_Paulo"}
	timestampTime := time.Now()
	formattedTimestamp := th.FormatTimestamp(timestampTime)

	timestampParsedToTime, err := th.ParseTimestampToTime(formattedTimestamp)

	assert.Nil(t, err, "Expected no error when parsing the timestamp")
	assert.NotNil(t, timestampParsedToTime, "Parsed timestamp should not be nil")
	assert.IsType(t, time.Time{}, timestampParsedToTime, "Expected the returned type to be time.Time")
}

func TestGetDateAndTimeFromTime(t *testing.T) {
	fixedTime := time.Date(2024, time.September, 26, 15, 30, 0, 0, time.UTC)
	th := TimeHelper{Timezone: "UTC"}

	date, time := th.GetDateAndTimeFromTime(fixedTime)

	expectedDate := "2024-09-26"
	expectedTime := "15:30:00"

	assert.Equal(t, expectedDate, date, "The returned date is not as expected")
	assert.Equal(t, expectedTime, time, "The returned time is not as expected")
}

func TestGetDateAndTimeFromString(t *testing.T) {
	formattedTimestamp := "2024-09-26 15:30:00 +00:00"
	th := TimeHelper{Timezone: "UTC"}

	date, time, err := th.GetDateAndTimeFromString(formattedTimestamp)

	expectedDate := "2024-09-26"
	expectedTime := "15:30:00"

	assert.Nil(t, err, "Expected no error when getting date and time from a formatted string")
	assert.Equal(t, expectedDate, date, "The returned date is not as expected")
	assert.Equal(t, expectedTime, time, "The returned time is not as expected")
}

