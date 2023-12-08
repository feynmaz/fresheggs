package v1

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, a any) {
	reponseBody, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reponseBody)
}
