package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var config Config

func main() {
	config.db = initDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(config.db)
	//PopulateAlmostCompleteGraphFromCSV()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/test", test)

	r.POST("/login", loginEndpointPost)
	r.POST("/signup", signupEndpointPost)
	r.GET("/routes/:from/:to", routesEndpoint)
	r.POST("/save-route", saveRouteEndpoint)
	r.DELETE("/delete-route", deleteRouteEndpoint)
	r.GET("/:id", savedTravelsEndpoint)

	// API pentru React
	r.POST("/api/best-route", bestRouteEndpoint)
	r.GET("/api/worldcities.csv", serveWorldCitiesCSV)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
