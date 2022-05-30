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
    res := MiddlewareRes {
        controllers.IsAuthorized(r.Header.Get("Authorization")),
        w,
        r,
    }
    return res
}
