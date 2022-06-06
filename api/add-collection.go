package api

import (
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func AddCollection(r middleware.MiddlewareRes) {
    if (r.Authorized) {
        q := r.Request.URL.Query()
        if len(q["name"]) != 1 {
            r.Writer.WriteHeader(400)
            return
        }
        name := q["name"][0]

        controllers.AddCollection(name)
        r.Writer.WriteHeader(200)
    } else {
        r.Writer.WriteHeader(403)
    }
}
