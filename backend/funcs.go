package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func loginEndpointPost(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	loginError := ""

	if !strings.Contains(email, "@") {
		loginError = "Invalid email address!"
	}

	row, err := config.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	for row.Next() {
		user := User{}
		err := row.Scan(&user.email, &user.password)
		if err != nil {
			log.Fatal(err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.password))
		if err != nil {
			log.Fatal(err)
			loginError = "Invalid password!"
		}
	}
	if loginError != "" {
		c.JSON(http.StatusOK, gin.H{
			"error": loginError,
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func signupEndpointPost(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	confirmationPassword := c.Query("confirmationPassword")
	signUpError := ""

	if !strings.Contains(email, "@") {
		signUpError = "Invalid email address!"
	} else if password != confirmationPassword {
		signUpError = "Passwords do not match!"
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, string(encryptedPassword))
	if err != nil {
		log.Fatal(err)
	}
	if signUpError != "" {
		c.JSON(http.StatusOK, gin.H{
			"error": signUpError,
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func routesEndpoint(c *gin.Context) {
	from, to := c.Query("from"), c.Query("to")
	routes := getRoutes()
	travels := getTravels(routes, from, to)
	c.JSON(http.StatusOK, gin.H{
		"travels": travels,
	})
}

func saveTravelEndpoint(c *gin.Context) {
	uid, cities, totalCost, duration := c.Query("id"), c.Query("cities"), c.Query("totalCost"), c.Query("duration")
	result, err := config.db.Exec("INSERT INTO travels(cities, totalCost, duration) VALUES(?, ?, ?)", cities, totalCost, duration)
	if err != nil {
		log.Fatal(err)
	}

	tid, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.db.Exec("INSERT INTO savedTravels(tid, uid) VALUES(?, ?)", tid, uid)
	if err != nil {
		log.Fatal(err)
	}
}

func savedTravelsEndpoint(c *gin.Context) {
	id := c.Query("id")
	row, err := config.db.Query("SELECT * FROM travels t, savedTravels ts WHERE t.id == st.tid AND st.uid == ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var travels []Travel
	for row.Next() {
		travel := Travel{}
		err := row.Scan(&travel.cities, &travel.totalCost, &travel.duration)
		if err != nil {
			log.Fatal(err)
		}
		travels = append(travels, travel)
	}
	c.JSON(http.StatusOK, gin.H{
		"travels": travels,
	})
}
