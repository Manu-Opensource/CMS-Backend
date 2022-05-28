package api

import (
    "encoding/json"
    "net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    res, _ := json.Marshal(map[string]string{"Message": "ping"})
    w.Write(res)
}
