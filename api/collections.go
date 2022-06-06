package api

import (
    "encoding/json"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func Collections(r middleware.MiddlewareRes) {
    r.Writer.Header().Set("Content-Type", "application/json")
    if (r.Authorized) {
        collections := controllers.LsCollections()
        res, _ := json.Marshal(collections)
        r.Writer.Write(res)
    } else {
        r.Writer.WriteHeader(403)
    }
}
