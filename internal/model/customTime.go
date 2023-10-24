package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime time.Time

var _ json.Marshaler = &CustomTime{}
var _ json.Unmarshaler = &CustomTime{}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Format(time.RFC3339) + `"`), nil
}

func (ct *CustomTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	if ct.String() == fmt.Sprintf("%q", "0001-01-01 00:00:00") {
		return nil, nil
	}
	return []byte(time.Time(ct).Format(time.RFC3339)), nil
}

func (ct *CustomTime) Scan(s string) error {
	tTime, _ := time.Parse(time.RFC3339, s)
	*ct = CustomTime(tTime)
	return nil
}

func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(time.RFC3339))
}
