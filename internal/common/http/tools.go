package http

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, a any) {
	reponseBody, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reponseBody)
}

func WriteErr(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}
