package main

import (
  "github.com/gin-gonic/gin"
)

var config Config

func main() {
  config.db = initDB()
  defer config.db.Close()

  r := gin.New()
  r.Use(gin.Logger())

  r.GET("/", homeEndpoint)
  r.POST("/login", loginEndpointPost)
  r.POST("/signup", signupEndpointPost)
  r.GET("/routes", routesEndpoint)
  r.GET("/routes/:location", locationEndpoint)

  admin := r.Group("/")
  admin.Use(adminRequired())
  {
    admin.GET("/admin", adminEndpoint)
  }
  r.Run()
}
