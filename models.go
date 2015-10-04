package alioss
import (
    "time"
    "errors"
)

const ISO8601 string = "2006-01-02T15:04:05.000Z"

type ISO8601Time struct {
    time.Time
}

// to UTC location
func (t *ISO8601Time) MarshalText() ([]byte, error) {
    if y := t.Year(); y < 0 || y >= 10000 {
        return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
    }
    return []byte(t.UTC().Format(ISO8601)), nil
}