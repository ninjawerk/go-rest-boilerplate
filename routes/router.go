/**
 * Created by VoidArtanis on 10/22/2017
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/VoidArtanis/go-rest-boilerplate/controllers"
	"github.com/VoidArtanis/go-rest-boilerplate/middlewares"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)
	authController := new(controllers.AuthController)
	engine.POST("/login",  authController.HandleLogin)
	RegisterProtectedRoutes(engine)
	RegisterPublicRoutes(engine)
	RegisterUtilityRoutes(engine)
}

func InitMiddleware(engine *gin.Engine){
	engine.Use(middlewares.CORSMiddleware());
}