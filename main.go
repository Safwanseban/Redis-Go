package main

import (
	"github.com/Safwanseban/Redis-Go/configs"
	"github.com/Safwanseban/Redis-Go/routes"
	"github.com/gin-gonic/gin"
)

var env = configs.Loadenv()
var R = gin.Default()

func init() {
	configs.GetRedis()
}
func main() {

	routes.Routes(R)

	R.Run(env.Port)

}
