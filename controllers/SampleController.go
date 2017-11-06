/**
 * Created by VoidArtanis on 10/22/2017
 */

package controllers

import "github.com/gin-gonic/gin"

func GetSecretText(c *gin.Context){
	c.JSON(200, "Hi this is a secret message. Auth was successful!")
}

func GetPublicText(c *gin.Context){
	c.JSON(200, "Hi this is a public message!")
}
