package api

import (
    "github.com/Manu-Opensource/CMS-Backend/middleware"
)

func IsAuthenticated(r middleware.MiddlewareRes) {
    if (r.Authorized) {
        r.Writer.WriteHeader(200)
    } else {
        r.Writer.WriteHeader(403)
    }
}
