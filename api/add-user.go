package api

import (
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func AddUser(r middleware.MiddlewareRes) {
    r.Writer.Header().Set("Content-Type", "application/json")
    if (r.Authorized) {
        q := r.Request.URL.Query()
        if len(q["username"]) != 1 || len(q["password"]) != 1 {
            r.Writer.WriteHeader(400)
            return
        }
        username := q["username"][0]
        password := q["password"][0]
        if controllers.DoesUserExist(username, password) {
            r.Writer.WriteHeader(406)
            return
        }

        controllers.AddUser(username, password)

    } else {
        r.Writer.WriteHeader(401)
    }
}
