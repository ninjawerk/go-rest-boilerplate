/**
 * Created by VoidArtanis on 10/22/2017
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/VoidArtanis/GoRestBoilerPlate/middlewares"
	"github.com/VoidArtanis/GoRestBoilerPlate/controllers"
)

func RegisterProtectedRoutes(r *gin.Engine){

	authGroup := r.Group("/auth")

	authGroup.Use(middlewares.AuthHandler("admin"))
	{
		authGroup.GET("/getmessage",controllers.GetSecretText)
	}
}
