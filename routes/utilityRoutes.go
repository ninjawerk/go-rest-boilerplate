/**
 * Created by VoidArtanis on 10/22/2017
 */

package routes

import "github.com/gin-gonic/gin"

func RegisterUtilityRoutes(r *gin.Engine){
	registerRing(r)
}

func registerRing(r *gin.Engine){
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}