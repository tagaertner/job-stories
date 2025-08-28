package models

import (
	"fmt"
	"time"
	"io"
)

type Time time.Time

func (t Time) MarshalGQL(w io.Writer) {
	str := time.Time(t).Format(time.RFC3339)
	_, _ = w.Write([]byte(fmt.Sprintf("%q", str)))
}

func (t *Time) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("time should be a string")
	}
	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*t = Time(parsed)
	return nil
}