package types

import (
	"fmt"
	"strings"
	"time"
)

type JsonDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonDate) ToTime() time.Time {
	t := time.Time(j)
	return t
}
func (j JsonDate) ToTimeString() string {
	return time.Time(j).Format("02 Jan 2006")
}
func (j JsonDate) ToTimeStringDefault() string {
	return time.Time(j).Format("2006-01-02")
}
