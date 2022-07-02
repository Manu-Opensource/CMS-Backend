package api

import (
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func DeleteDocument(r middleware.MiddlewareRes) {
    if (r.Authorized) {
        q := r.Request.URL.Query()
        if len(q["documentid"]) != 1 || len(q["collectionname"]) != 1 {
            r.Writer.WriteHeader(400)
            return
        }

        collectionName := q["collectionname"][0]
        documentId := q["documentid"][0]

        controllers.DeleteDocument(collectionName, documentId)
        r.Writer.WriteHeader(200)
    } else {
        r.Writer.WriteHeader(403)
    }
}
