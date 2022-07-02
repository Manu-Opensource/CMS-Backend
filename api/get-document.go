package api

import (
    "encoding/json"
    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func GetDocument(r middleware.MiddlewareRes) {
    if (r.Authorized) {
        q := r.Request.URL.Query()
        if len(q["documentid"]) != 1 || len(q["collectionname"]) != 1 {
            r.Writer.WriteHeader(400)
            return
        }
        collectionName := q["collectionname"][0]
        documentId := q["documentid"][0]

        document := controllers.GetDocument(collectionName, documentId)
        res, _ := json.Marshal(document)
        r.Writer.Write(res)
    } else {
        r.Writer.WriteHeader(403)
    }
}
