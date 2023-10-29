package types

import (
	"fmt"
	"strings"
	"time"
)

type JsonDateTime time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02  15:04:05", s)
	if err != nil {
		return err
	}
	*j = JsonDateTime(t)
	return nil
}

func (j JsonDateTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Maybe a Format function for printing your date
func (j JsonDateTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonDateTime) ToTime() time.Time {
	t := time.Time(j)
	return t
}
func (j JsonDateTime) ToTimeString() string {
	return time.Time(j).Format("2006-01-02 15:04:05")
}
