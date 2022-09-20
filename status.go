package httprdb

import "net/http"

// type struct implementation of the "apistatus" type
//
// is used to declare the current status of an api and to
// respond to "healthz"-like requests
type status struct {
	Code    int    `json:"code"`
	Ready   bool   `json:"ready"`
	Healthy bool   `json:"healthy"`
	Message string `json:"msg"`
}

// sets the current status of the httpserver
//
// Parameters:
// 	- `code` : int > http statuscode (200, 400, 500, ...)
// 	- `ready` : bool > readyness value
// 	- `healthy` : bool > healthyness value
func (aStatus *status) set(code int, ready, healthy bool) {
	aStatus.Code = code
	aStatus.Ready = ready
	aStatus.Healthy = healthy
	aStatus.Message = http.StatusText(code)
}

// get the current status, unstructed
//
// output is (http.StatusCode, Ready, Healthy, http.StatusMessage)
func (aStatus status) Get() (int, bool, bool, string) {

	return aStatus.Code, aStatus.Ready, aStatus.Healthy, aStatus.Message
}
