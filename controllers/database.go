package controllers

import (
    "fmt"
    "sync"
    "context"
    "log"
    
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connect() (*mongo.Database) {
    user := Getenv("DB_USER")
    pass := Getenv("DB_PASS")
    link := Getenv("DB_LINK")
    abslink := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, pass, link)

    client, err := mongo.NewClient(options.Client().ApplyURI(abslink))

    if err != nil {
        log.Fatal(err)
    }

    client.Connect(context.Background())
    database := client.Database("Main")

    if err != nil {
        log.Fatal(err)
    }

    return database
}

func DbCon() (*mongo.Database) {
    var conOnce sync.Once
    var database *mongo.Database
    conOnce.Do(func() {
        database = connect()
    })
    return database
}

func LsCollections(d *mongo.Database) ([]string) {
    ret, err := d.ListCollectionNames(context.Background(), bson.M{}) 
    if err != nil {
        log.Print(err)
    }
    return ret
}

func GetCollection(d *mongo.Database, name string) (*mongo.Collection) {
    ret := d.Collection(name)
    return ret
}

func AddCollection(d *mongo.Database, name string) (*mongo.Collection) {
    err := d.CreateCollection(context.Background(), name)
    if err != nil {
        log.Print(err)
    }
    return GetCollection(d, name)
}
