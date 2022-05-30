package api

import (
    "encoding/json"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
)

func Ping(r middleware.MiddlewareRes) {
    r.Authorized = true
    r.Writer.Header().Set("Content-Type", "application/json")
    res, _ := json.Marshal(map[string]string{"Message": "ping"})
    r.Writer.Write(res)
}
