package routers


import (
    "net/http"
    "github.com/Manu-Opensource/CMS-Backend/routers/api"
)

func addRoute(path string, f func(http.ResponseWriter, *http.Request)) {
    rHandler := http.HandlerFunc(f)
    http.Handle(path, rHandler)
}

func RouterInit() {
    addRoute("/api/ping", api.Ping)
    http.ListenAndServe(":8080", nil)
}
