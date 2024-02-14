package utils

import (
	"fabiloco/hotel-trivoli-api/api/config"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
  Id       uint `json:"id"`
  Username string `json:"username"`
  Role     string `json:"role"`
  jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
 accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

 return accessToken.SignedString([]byte(config.Config("TOKEN_SECRET")))
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
 refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

 return refreshToken.SignedString([]byte(config.Config("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) (*UserClaims, error){
  parsedAccessToken, error := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(config.Config("TOKEN_SECRET")), nil
  })

  if error != nil {
    return nil, error 
  }

 return parsedAccessToken.Claims.(*UserClaims), nil
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
 parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
  return []byte(config.Config("TOKEN_SECRET")), nil
 })

 return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}
