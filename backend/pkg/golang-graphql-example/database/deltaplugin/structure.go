package deltaplugin

type ActionType string

const (
	UPDATE ActionType = "UPDATE"
	PATCH  ActionType = "PATCH"
	DELETE ActionType = "DELETE"
	CREATE ActionType = "CREATE"
)

type Delta struct {
	Patch  map[string]interface{} `json:"patch"`
	Result interface{}            `json:"result"`
	Table  string                 `json:"table"`
	Action ActionType             `json:"action"`
}
