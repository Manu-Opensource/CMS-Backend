package middleware

import (
    "net/http"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

type MiddlewareRes struct {
    Authorized bool
    Writer http.ResponseWriter
    Request *http.Request
}

func MiddlewareManager(w http.ResponseWriter, r *http.Request) (MiddlewareRes) {
    cookie, err := r.Cookie("Authorization")
    res := MiddlewareRes {
        Writer: w,
        Request: r,
    }
    if err == nil {
        res.Authorized = controllers.IsAuthorized(cookie.Value)
    } else {
        res.Authorized = false
    }
    return res;
}
