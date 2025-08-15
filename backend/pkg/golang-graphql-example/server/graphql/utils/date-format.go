package utils

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"emperror.dev/errors"
)

// Date Format enumeration.
type DateFormat string

var (
	DateFormatRFC3339     DateFormat = "RFC3339"
	DateFormatRFC3339Nano DateFormat = "RFC3339Nano"
)

var AllDateFormat = []DateFormat{
	DateFormatRFC3339,
	DateFormatRFC3339Nano,
}

func (e DateFormat) IsValid() bool {
	switch e {
	case DateFormatRFC3339, DateFormatRFC3339Nano:
		return true
	}

	return false
}

func (e DateFormat) String() string {
	return string(e)
}

func (e *DateFormat) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return errors.Errorf("enums must be strings")
	}

	*e = DateFormat(str)
	if !e.IsValid() {
		return errors.Errorf("%s is not a valid DateFormat", str)
	}

	return nil
}

func (e DateFormat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func FormatTime(format *DateFormat, ti time.Time) string {
	// Check if it is null
	if format == nil {
		// Override
		format = &DateFormatRFC3339
	}

	switch *format {
	case DateFormatRFC3339Nano:
		return ti.UTC().Format(time.RFC3339Nano)
	case DateFormatRFC3339:
		fallthrough
	default:
		return ti.UTC().Format(time.RFC3339)
	}
}
