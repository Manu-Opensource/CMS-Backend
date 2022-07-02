package controllers

import (
    "fmt"
    "context"
    "log"
    
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

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
    if database == nil {
        database = connect()
    }
    return database
}

func contentChange(change string) {
    changes := GetCollection("CMSContentChanges")
    changes.InsertOne(context.Background(), bson.M{"change": change})
}

func LsCollections() ([]string) {
    d := dbCon()
    ret, err := d.ListCollectionNames(context.Background(), bson.M{}) 
    if err != nil {
        log.Print(err)
    }
    return ret
}

func ReadCollection(name string) ([]bson.M) {
    cursor, err := GetCollection(name).Find(context.Background(), bson.M{})
    if err != nil {
        log.Print(err)
    }
    var ret []bson.M
    err = cursor.All(context.Background(), &ret)
    if err != nil {
        log.Print(err)
    }
    return ret
}

func ReadCollectionWithFilter(name string, filter bson.M) ([]bson.M) {
    cursor, err := GetCollection(name).Find(context.Background(), filter)
    if err != nil {
        log.Print(err)
    }
    var ret []bson.M
    err = cursor.All(context.Background(), ret)
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
    contentChange(fmt.Sprintf("Added Collection [%s]", name))
    return GetCollection(name)
}

func GetDocument(collectionName string, documentId string) (interface{}) {
    var result interface{}
    GetCollection(collectionName).FindOne(context.Background(), bson.D{{"Id", documentId}}).Decode(&result)
    return result
}

func DeleteCollection(collectionName string) {
    err := GetCollection(collectionName).Drop(context.Background())
    if err != nil {
        log.Print(err)
    }
    contentChange(fmt.Sprintf("Delete Collection [%s]", collectionName))
}

func CreateDocument(collectionName string, doc []map[string]interface{}) {
    comb := bson.M{}
    var documentId string
    for _, element := range doc {
        comb[element["name"].(string)] = element["value"].(string)
        if element["name"] == "Id" {
            documentId = element["value"].(string)
        }
    }
    _, err := GetCollection(collectionName).InsertOne(context.Background(), comb)
    if err != nil {
        log.Print(err)
    }
    contentChange(fmt.Sprintf("Created Document [%s] in Collection [%s]", documentId, collectionName))
}

func UpdateDocument(collectionName string, documentId string, doc[]map[string]interface{}) {
    comb := bson.M{}
    for _, element := range doc {
        comb[element["name"].(string)] = element["value"].(string)
    }
    GetCollection(collectionName).FindOneAndReplace(context.Background(), bson.D{{"Id", documentId}}, comb)
}

func DeleteDocument(collectionName string, documentId string) {
    fmt.Println("Deleting Document", documentId)
    GetCollection(collectionName).FindOneAndDelete(context.Background(), bson.D{{"Id", documentId}})
    contentChange(fmt.Sprintf("Deleted Document [%s] in Collection [%s]", documentId, collectionName))
}

func DoesCMSUserExist(user string, pass string) (bool) {
    var fRes bson.M //ToDo: Check for hashed
    GetCollection("CMSUsers").FindOne(context.Background(), bson.M{"user": user}).Decode(&fRes)
    if fRes == nil {
        return false
    }
    return bcrypt.CompareHashAndPassword([]byte(fRes["pass"].(string)), []byte(pass)) == nil
}

func AddCMSUser(user string, pass string) {
    passHashed,_ := bcrypt.GenerateFromPassword([]byte(pass), 5)
    GetCollection("CMSUsers").InsertOne(context.Background(), bson.M{"user": user, "pass": string(passHashed[:])})
}
