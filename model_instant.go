package apono

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var nilTime = (time.Time{}).UnixNano()

type Instant struct {
	time.Time
}

func (i *Instant) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		i.Time = time.Time{}
		return
	}

	var (
		seconds int64
		nano    int64
	)
	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return errors.New("illegal time format")
	}

	seconds, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return err
	}

	if len(parts) == 2 {
		nano, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return err
		}
	}

	i.Time = time.Unix(seconds, nano)
	return
}

func (i *Instant) MarshalJSON() ([]byte, error) {
	if i.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%d.%d\"", i.Time.Unix(), i.Time.Nanosecond())), nil
}

func (i *Instant) IsSet() bool {
	return i.UnixNano() != nilTime
}

type NullableInstant struct {
	value *Instant
	isSet bool
}

func (v NullableInstant) Get() *Instant {
	return v.value
}

func (v *NullableInstant) Set(val *Instant) {
	v.value = val
	v.isSet = true
}

func (v NullableInstant) IsSet() bool {
	return v.isSet
}

func (v *NullableInstant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstant(val *Instant) *NullableInstant {
	return &NullableInstant{value: val, isSet: true}
}

func (v NullableInstant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
