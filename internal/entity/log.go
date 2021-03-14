package entity

import "encoding/json"

// LogActionRequest is entity request to do log action
type LogActionRequest struct {
	Action       string `db:"action"`
	Method       string `db:"method"`
	Request      interface{}
	RequestJSON  json.RawMessage `db:"request"`
	Response     interface{}
	ResponseJSON json.RawMessage `db:"response"`
}
