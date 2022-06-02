package controllers

import (
    "fmt"
    "sync"
    "context"
    "log"
    
    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
    Link string 
    Client *mongo.Client
}

var dbCon *DatabaseConnection
var mutLock = &sync.Mutex{}

func getConnection() (*DatabaseConnection) {
    if dbCon == nil {
        mutLock.Lock()
        defer mutLock.Unlock()
        if dbCon == nil {
            dbCon = &DatabaseConnection{}
            dbCon.Init()
        }
    }
    return dbCon
}

func (c DatabaseConnection) Init() {
    user := Getenv("DB_USER")
    pass := Getenv("DB_PASS")
    link := Getenv("DB_LINK")
    c.Link = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, pass, link)

    var err error
    c.Client, err = mongo.NewClient(options.Client().ApplyURI(c.Link))

    if err != nil {
        log.Fatal(err)
    }

    c.Client.Connect(context.Background())
}

func (c DatabaseConnection) Disconnect() {
    c.Client.Disconnect(context.Background())
}

