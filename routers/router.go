package routers


import (
    "net/http"
    "github.com/Manu-Opensource/CMS-Backend/api"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
)


func addRoute(path string, f func(middleware.MiddlewareRes)) {
    rHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        res := middleware.MiddlewareManager(w, r)
        f(res)
    })
    http.Handle(path, rHandler)
}

func RouterInit() {
    addRoute("/api/ping", api.Ping)
    addRoute("/api/login", api.Login)
    addRoute("/api/adduser", api.AddUser)
    http.ListenAndServe(":8080", nil)
}
