/**
 * Created by VoidArtanis on 11/2/2017
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/VoidArtanis/GoRestBoilerPlate/middlewares"
	"github.com/VoidArtanis/GoRestBoilerPlate/shared"
)

type AuthController struct{}


func (this AuthController)HandleLogin(c *gin.Context) {
	userId:="123"
	username:="Beast"
	roles:= []string{shared.RoleAdmin, shared.RoleProUser}

	// do user auth here

	//issue token
	token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId,username, roles)
	if err != nil {

	}
	c.JSON(200, token)
}
