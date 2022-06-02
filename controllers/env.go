package controllers

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()    

    if err != nil {
        log.Fatal("Error whilst attempting to load .env file")
    }
}

func Getenv(k string) (string) {
    return os.Getenv(k)
}
