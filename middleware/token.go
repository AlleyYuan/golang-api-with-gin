package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myproject/model"
	"myproject/serializer"
	"net/http"
	"time"
	//"errors"
)

//CustomClaims to define some custom claims
type CustomClaims struct{
	ID uint `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//Auth middleware
func Auth() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK,serializer.Response{
				Status: 40005,
				Data:   nil,
				Msg:    "token invalid",
				Error:  "",
				Token:  "",
			})
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(200,serializer.Response{
				Status: 40004,
				Msg: "parse token error",
			})
			c.Abort()
			return
		}
		c.Set("claims",claims)
	}
}

//JWT secret
type JWT struct{
	secret []byte
}

//
const(
	secretkey string = "tokensecret"
)

//GenerateToken to generare token
func GenerateToken(user model.User) (string,bool){
	claims := CustomClaims{
		user.ID,
		user.Username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer: "alley",
		},
	}

	token, err := CreateToken(claims)
	if err != nil{
		return "",false
	}
	return token,true

}

//CreateToken to create token
func CreateToken(claims CustomClaims) (string ,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(secretkey))
}

//ParseToken to parse token
func ParseToken(tokenString string)(*CustomClaims,error) {
	token,err := jwt.ParseWithClaims(tokenString,&CustomClaims{},func(token *jwt.Token)(interface{},error){
		return []byte(secretkey),nil
	})
	if err != nil{
		return nil,err
	}
	if claims,ok :=token.Claims.(*CustomClaims); ok && token.Valid{
		return claims,nil
	}
	return nil,err
}