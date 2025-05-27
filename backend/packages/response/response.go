package response

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Content any    `json:"content,omitempty"`
}

type Error Message

func Json(w http.ResponseWriter, data ...any) {
	if len(data) == 0 {
		Set(w, Message{Message: `Ok!`})
		return
	}

	if m, ok := data[0].(string); ok {
		Set(w, Message{Message: m})
		return
	}

	Set(w, Message{Content: data[0]})
	return
}

func ErrJson(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	Set(w, Message{
		Status:  1,
		Message: err.Error(),
	})
}

func Set(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
