package httputils

import (
	"encoding/json"
	"net/http"
)

const (
	CodeOK       = 20000
	CodeInterval = 50000
)

func ResponseJson(w http.ResponseWriter, code int, message string, data interface{}) (int, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	d := map[string]interface{}{
		"code":    200,
		"message": message,
		"data":    data,
	}
	jsonData, err := json.Marshal(d)
	if err != nil {
		return w.Write([]byte("failed"))
	}
	return w.Write(jsonData)
}
