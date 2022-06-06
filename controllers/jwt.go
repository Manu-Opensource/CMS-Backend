package controllers

import (
    "time"
    "github.com/golang-jwt/jwt"
)

var key = []byte(Getenv("SECRET_KEY"))

type identityClaim struct {
    Authorized bool
    *jwt.StandardClaims
}

func IsAuthorized(tokenS string) bool {
    if (tokenS == "") {
        return false;
    }
    t, _ := jwt.ParseWithClaims(tokenS, &identityClaim{}, func(token *jwt.Token) (interface{}, error) {
        return key, nil
    })
    
    claims := t.Claims.(*identityClaim)
    return claims.Authorized
}

func CreateAuthorizedToken() (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    token.Claims = &identityClaim {
        true,
        &jwt.StandardClaims {
            ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
        },
    }
    return token.SignedString(key)
}

func AssignAuthToken(user string, pass string) (string, error) {
    isAuthorized := DoesCMSUserExist(user, pass)
    if isAuthorized {
        return CreateAuthorizedToken()
    } else {
        return "Invalid", nil
    }
}
