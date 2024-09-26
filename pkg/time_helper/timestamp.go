package time_helper

import (
	"time"
)

const timestampFormatLayout = "2006-01-02 15:04:05 -07:00"
const timestamptDateLayout = "2006-01-02"
const timestampTimeLayout = "15:04:05"

type TimeHelper struct{
	Timezone string
}

func (th *TimeHelper) CreateTimestampFromTimezone () (*time.Time, error) {
	location, err := time.LoadLocation(th.Timezone)
	if err != nil {
		return nil, err
	}
	timestamp := time.Now().In(location)
	return &timestamp, nil
}

func (th *TimeHelper) CreateAndFormatTimestampFromTimezone () (*string, error) {
	timestamp, err := th.CreateTimestampFromTimezone()
    if err != nil {
        return nil, err
    }
    formattedTimestamp := th.FormatTimestamp(*timestamp)
    return &formattedTimestamp, nil
}

func (th *TimeHelper) FormatTimestamp(timestamp time.Time) string { 
	return timestamp.Format(timestampFormatLayout)
}

func (th *TimeHelper) ParseTimestampToTime(formatedTimestamp string) (time.Time,  error) {
	layout := timestampFormatLayout
	t, err := time.Parse(layout, formatedTimestamp)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func (th *TimeHelper)GetDateAndTimeFromTime(timestamp time.Time) (string, string) {
	date := timestamp.Format(timestamptDateLayout) 
	time := timestamp.Format(timestampTimeLayout)   
	return date, time
}

func (th *TimeHelper)GetDateAndTimeFromString(timestampStr string) (string, string, error) {
	timestamp, err := th.ParseTimestampToTime(timestampStr)
	if err != nil {
		return "", "", err
	}
	date := timestamp.Format(timestamptDateLayout)
	time := timestamp.Format(timestampTimeLayout)
	return date, time, nil
}
