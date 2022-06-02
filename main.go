package main

import ( 
    "github.com/Manu-Opensource/CMS-Backend/routers"
    "github.com/Manu-Opensource/CMS-Backend/controllers"
)

func main() {
    controllers.LoadEnv()

    routers.RouterInit()
}

