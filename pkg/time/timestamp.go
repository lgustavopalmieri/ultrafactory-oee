package location_entity

import (
	"time"
)

const timestampFormatLayout = "2006-01-02 15:04:05 -07:00"
const timestamptDateLayout = "2006-01-02"
const timestampTimeLayout = "15:04:05"

func CreateTimestampFromTimezone (timezone string) (*string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	now := time.Now().In(location)
	timestamp := FormatTimestamp(now)
	return &timestamp, nil
}

func FormatTimestamp(timestamp time.Time) string{
	return timestamp.Format(timestampFormatLayout)
}

func ParseTimestampToTime(formatedTimestamp string) (time.Time,  error) {
	layout := timestampFormatLayout
	t, err := time.Parse(layout, formatedTimestamp)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func GetDateAndTimeFromTime(timestamp time.Time) (string, string) {
	date := timestamp.Format(timestamptDateLayout) 
	time := timestamp.Format(timestampTimeLayout)   
	return date, time
}

func GetDateAndTimeFromString(timestampStr string) (string, string, error) {
	timestamp, err := ParseTimestampToTime(timestampStr)
	if err != nil {
		return "", "", err
	}
	date := timestamp.Format(timestamptDateLayout)
	time := timestamp.Format(timestampTimeLayout)
	return date, time, nil
}