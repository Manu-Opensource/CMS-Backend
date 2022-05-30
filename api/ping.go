package api

import (
    "encoding/json"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
)

func Ping(r middleware.MiddlewareRes) {
    r.Writer.Header().Set("Content-Type", "application/json")
    if (r.Authorized) {
        res, _ := json.Marshal(map[string]string{"Message": "Authorized"})
        r.Writer.Write(res)
    } else {
        res, _ := json.Marshal(map[string]string{"Message": "Unauthorized"})
        r.Writer.Write(res)
    }
}
