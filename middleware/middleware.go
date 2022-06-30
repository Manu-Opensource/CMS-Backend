package middleware

import (
    "net/http"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

type MiddlewareRes struct {
    Authorized bool
    IsGetRequest bool
    IsPostRequest bool
    IsOptionsRequest bool
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

    if r.Method == "GET" {
        res.IsGetRequest = true
    }
    if r.Method == "POST" {
        res.IsPostRequest = true
    }
    if r.Method == "OPTIONS" {
        res.IsOptionsRequest = true
    }

    w.Header().Set("Access-Control-Allow-Origin", controllers.Getenv("FRONTEND_LINK"))
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "content-type")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    return res;
}
