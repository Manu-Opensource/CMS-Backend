package main

import ( 
    "fmt"
    "github.com/Manu-Opensource/CMS-Backend/routers"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func main() {
    controllers.LoadEnv()

    fmt.Println(controllers.LsCollections(controllers.DbCon()))

    routers.RouterInit()
}

