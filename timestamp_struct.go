package main

import (
	"database/sql/driver"
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(buf []byte) error {

	i, err := strconv.Atoi(strings.Trim(string(buf), "\""))
	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(i), 0)
	return nil

}

func (t *Timestamp) Scan(value interface{}) error {

	t.Time = value.(time.Time)

	return nil

}

func (t Timestamp) Value() (driver.Value, error) {

	return t.Time, nil

}
