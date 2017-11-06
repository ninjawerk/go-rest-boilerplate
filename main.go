/**
 * Created by VoidArtanis on 10/22/2017
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/VoidArtanis/GoRestBoilerPlate/routes"
	"github.com/VoidArtanis/GoRestBoilerPlate/shared"
)

var DB = make(map[string]string)

func main() {

	//Db Connect and Close
	shared.Init()
	defer shared.CloseDb()

	//Init Gin
	r := gin.Default()
	routes.InitRouter(r)

	//Run Server
	r.Run(":8080")
}