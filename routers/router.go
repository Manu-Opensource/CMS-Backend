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
    addRoute("/api/addcollection", api.AddCollection)
    addRoute("/api/isauthenticated", api.IsAuthenticated)
    addRoute("/api/collections", api.Collections)
    addRoute("/api/getcollection", api.GetCollection)
    addRoute("/api/createdocument", api.CreateDocument)
    addRoute("/api/getdocument", api.GetDocument)
    addRoute("/api/updatedocument", api.UpdateDocument)
    addRoute("/api/deletedocument", api.DeleteDocument)
    http.ListenAndServe(":8081", nil)
}
