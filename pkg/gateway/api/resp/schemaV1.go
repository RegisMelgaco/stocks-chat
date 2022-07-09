package resp

type ErrorV1 struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
}
