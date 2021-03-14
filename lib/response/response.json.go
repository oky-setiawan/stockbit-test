package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

//JSONResponse is entity of JSON Response writer
type JSONResponse struct {
	Data         interface{}       `json:"data,omitempty"`
	Message      string            `json:"message,omitempty"`
	ErrorMessage string            `json:"error,omitempty"`
	Status       bool              `json:"ok"`
	HTTPCode     int               `json:"-"`
	Latency      string            `json:"latency"`
	ServerTime   string            `json:"server_time"`
	header       map[string]string `json:"-"`
	CustWriter
}

type CustWriter struct {
	Use            bool                      `json:"-"`
	CustWriterBody bool                      `json:"-"`
	CustWriterFn   func(http.ResponseWriter) `json:"-"`
}

//NewJSONResponse wiil create new JSONResponse
func NewJSONResponse() *JSONResponse {
	return &JSONResponse{
		HTTPCode: http.StatusOK,
		header:   map[string]string{},
		Status:   true,
	}
}

//SetData will set data into JSONResponse
func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

//SetError will set error and message, you can use custom message by passing variadic msg variable
func (r *JSONResponse) SetError(err error, msg ...string) *JSONResponse {
	if err != nil {
		r.HTTPCode = GetHTTPCode(err)
		r.ErrorMessage = err.Error()
		r.Status = false
		r.Message = GetErrorMessage(err).String()
		if len(msg) > 0 {
			r.Message = msg[0]
		}
	}
	return r
}

//SetMessage will set response message
func (r *JSONResponse) SetMessage(message string) *JSONResponse {
	r.Message = message
	return r
}

//SetHeader will set header value based on key value header
func (r *JSONResponse) SetHeader(key, value string) *JSONResponse {
	r.header[key] = value
	return r
}

//SetLatency will set latency of api
func (r *JSONResponse) SetLatency(latency float64) *JSONResponse {
	r.Latency = fmt.Sprintf("%.2f ms", latency)
	return r
}

func (r *JSONResponse) CustomWriter(withBody ...bool) *JSONResponse {
	r.CustWriter.Use = true
	if len(withBody) > 0 {
		r.CustWriterBody = withBody[0]
	}
	return r
}

//Send will send the data to writer and will be consume to client
func (r *JSONResponse) Send(w http.ResponseWriter) {
	if r.CustWriter.Use {
		r.CustWriterFn(w)
		r.send(w, r.CustWriterBody)
	} else {
		r.send(w, true)
	}
}

func (r *JSONResponse) send(w http.ResponseWriter, withBody bool) {
	w.Header().Set("Content-Type", "application/json")
	for k, v := range r.header {
		w.Header().Set(k, v)
	}
	w.WriteHeader(r.HTTPCode)

	r.ServerTime = time.Now().Format(time.RFC3339)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Errorf("[JSONResponse] error encoding response, err: %v", err.Error())
	}
}
