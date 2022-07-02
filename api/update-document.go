package api

import (
    "fmt"
    "io/ioutil"
    "encoding/json"

    "github.com/Manu-Opensource/CMS-Backend/middleware"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

type updateDocumentRequestStruct struct {
    CollectionName string `json:"collectionName"`
    DocumentId string `json:"documentId"`
    Doc []map[string]interface{} `json:"doc"`
}

func UpdateDocument(r middleware.MiddlewareRes) {
    if r.Authorized {
        if r.Request.Header.Get("Content-Type") != "application/json" {
            r.Writer.WriteHeader(400)
            return
        }

        var decoded updateDocumentRequestStruct

        body, _ := ioutil.ReadAll(r.Request.Body)
        err := json.Unmarshal(body, &decoded)

        if err != nil {
            fmt.Println(err)
        }

        controllers.UpdateDocument(decoded.CollectionName, decoded.DocumentId, decoded.Doc)

        r.Writer.WriteHeader(200)
    } else {
        if r.IsOptionsRequest {
            r.Writer.WriteHeader(200)
        } else {
            r.Writer.WriteHeader(403)
        }
    }
}
