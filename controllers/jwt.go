package controllers

import (
    "time"
    "github.com/golang-jwt/jwt"
)

var key = "tmpkey" //ToDo: move to an env variable.

type identityClaim struct {
    *jwt.StandardClaims
    authorized bool `json:"authorized"`
}

func IsAuthorized(tokenS string) bool {
    if (tokenS == "") {
        return false;
    }
    t, _ := jwt.ParseWithClaims(tokenS, &identityClaim{}, func(token *jwt.Token) (interface{}, error) {
        return key, nil
    })
    
    claims := t.Claims.(*identityClaim)
    return claims.authorized
}

func CreateAuthorizedToken() (string, error) {
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    token.Claims = &identityClaim {
        &jwt.StandardClaims {
            ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
        },
        true,
    }
    return token.SignedString(key)
}
