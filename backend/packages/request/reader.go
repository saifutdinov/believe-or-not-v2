package request

import (
	"encoding/json"
	"net/http"
)

func Bind(r *http.Request, i any) error {
	dec := json.NewDecoder(r.Body)
	// dec.DisallowUnknownFields()
	return dec.Decode(i)
}
