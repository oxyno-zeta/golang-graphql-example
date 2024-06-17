package deltaplugin

import (
	"encoding/json"
	"time"

	"emperror.dev/errors"
)

type ActionType string

const (
	UPDATE ActionType = "UPDATE"
	PATCH  ActionType = "PATCH"
	DELETE ActionType = "DELETE"
	CREATE ActionType = "CREATE"
)

// Create a custom time type for JSON Marshal/Unmarshal in nano seconds.
type NanoDateTime time.Time

type Delta struct {
	EventDate NanoDateTime           `json:"eventDate"`
	Patch     map[string]interface{} `json:"patch"`
	Result    interface{}            `json:"result"`
	Table     string                 `json:"table"`
	Action    ActionType             `json:"action"`
}

func (d NanoDateTime) MarshalJSON() ([]byte, error) {
	// Cast time
	t := time.Time(d)
	// Format in RFC3339 Nano
	jsonStr := t.Format(time.RFC3339Nano)

	return []byte("\"" + jsonStr + "\""), nil
}

func (d *NanoDateTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	// Init time
	var t time.Time
	// Parse
	err := json.Unmarshal(data, &t)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Save
	*d = NanoDateTime(t)

	return nil
}
