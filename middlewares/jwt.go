/**
 * Created by VoidArtanis on 10/22/2017
 */

package middlewares

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SigningKey = "$SolidSigningKey$"
)

func AuthHandler(authRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer: "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Your request is not authorized"})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied"})
			c.Abort()
			return
		}

		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			c.JSON(403, gin.H{"message": "Invalid authorization token"})
			c.Abort()
			return
		}

		//authorize

		tokenRolesIntf := valid.Claims.(jwt.MapClaims)["roles"].([]interface{})
		var tokenRoles []string
		for _, v := range tokenRolesIntf {
			tokenRoles = append(tokenRoles, v.(string))
		}
		for _, v := range authRoles {
			hasRole := contains(tokenRoles, v)
			if !hasRole {
				c.JSON(403, gin.H{"message": "Not authorized to perform action"})
				c.Abort()
			}
		}

		// set variables
		c.Set("userId", valid.Claims.(jwt.MapClaims)["user_id"])
		c.Set("username", valid.Claims.(jwt.MapClaims)["username"])

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func GenerateToken(key []byte, userId string, username string, roles []string) (string, error) {

	//new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := make(jwt.MapClaims)
	claims["user_id"] = userId
	claims["username"] = username
	claims["exp"]=time.Now().Add(time.Hour * 72).UnixNano() / int64(time.Millisecond)

	//Set user roles
	claims["roles"] = roles

	token.Claims = claims

	// Sign and get as a string
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func ValidateToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return token, err
}

