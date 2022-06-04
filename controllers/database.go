package controllers

import (
    "fmt"
    "sync"
    "context"
    "log"
    
    "golang.org/x/crypto/bcrypt"
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

func dbCon() (*mongo.Database) {
    var conOnce sync.Once
    var database *mongo.Database
    conOnce.Do(func() {
        database = connect()
    })
    return database
}

func LsCollections() ([]string) {
    d := dbCon()
    ret, err := d.ListCollectionNames(context.Background(), bson.M{}) 
    if err != nil {
        log.Print(err)
    }
    return ret
}

func GetCollection(name string) (*mongo.Collection) {
    d := dbCon()
    ret := d.Collection(name)
    return ret
}

func AddCollection(name string) (*mongo.Collection) {
    d := dbCon()
    err := d.CreateCollection(context.Background(), name)
    if err != nil {
        log.Print(err)
    }
    return GetCollection(name)
}

func DeleteCollection(name string) {
    err := GetCollection(name).Drop(context.Background())
    if err != nil {
        log.Print(err)
    }
}

func DoesUserExist(user string, pass string) (bool) {
    var fRes bson.M //ToDo: Check for hashed
    GetCollection("Users").FindOne(context.Background(), bson.M{"user": user}).Decode(&fRes)
    if fRes == nil {
        return false
    }
    return bcrypt.CompareHashAndPassword([]byte(fRes["pass"].(string)), []byte(pass)) == nil
}

func AddUser(user string, pass string) {
    passHashed,_ := bcrypt.GenerateFromPassword([]byte(pass), 5)
    GetCollection("Users").InsertOne(context.Background(), bson.M{"user": user, "pass": string(passHashed[:])})
}
