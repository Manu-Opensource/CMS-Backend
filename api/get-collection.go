package api

import (
    "encoding/json"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func GetCollection(r middleware.MiddlewareRes) {
    r.Writer.Header().Set("Content-Type", "application/json")
    if (r.Authorized) {
        q := r.Request.URL.Query()
        if len(q["name"]) != 1 {
            r.Writer.WriteHeader(400)
            return
        }
        name := q["name"][0]

        documents := controllers.ReadCollection(name)
        res, _ := json.Marshal(documents)
        r.Writer.Write(res)
    } else {
        r.Writer.WriteHeader(403)
    }
}
