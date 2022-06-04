package api

import (
    "net/http"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func Login(r middleware.MiddlewareRes) {
    q := r.Request.URL.Query()
    if len(q["username"]) != 1 || len(q["password"]) != 1 {
        r.Writer.WriteHeader(400)
        return
    }

    t, err := controllers.AssignAuthToken(q["username"][0], q["password"][0])

    if err != nil {
        r.Writer.WriteHeader(400)
        return
    } else {
        if t != "Invalid" {
            cookie := &http.Cookie {
                Name: "Authorization",
                Value: t,
                MaxAge: 60 * 60 * 24 * 2,
            }
            http.SetCookie(r.Writer, cookie)
            r.Writer.WriteHeader(200)
        } else {
            r.Writer.WriteHeader(403)
        }
    }
}
