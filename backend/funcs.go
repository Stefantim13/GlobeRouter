package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func loginEndpointPost(c *gin.Context) {
  email := c.Query("email")
  password := c.Query("password")

  if !strings.Contains(email, "@") {
    log.Fatal("Invalid email address!")
  }

  row, err := config.db.Query("SELECT * FROM users WHERE email = ?", email)
  if err != nil {
    log.Fatal(err)
  }
  defer row.Close()

  for row.Next() {
    user := User{}
    err := row.Scan(&user.email, &user.password)
    if err != nil {
      log.Fatal(err)
    }

    err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.password))
    if err != nil {
      log.Fatal(err)
    }
  }
}

func signupEndpointPost(c *gin.Context) {
  email := c.Query("email")
  password := c.Query("password")
  confirmationPassword := c.Query("confirmationPassword")

  if !strings.Contains(email, "@") {
    log.Fatal("Invalid email address!")
  } else if password != confirmationPassword {
    log.Fatal("Passwords do not match!")
  }

  encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  if err != nil {
    log.Fatal(err)
  }

  config.db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, string(encryptedPassword))
}

func homeEndpoint(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func routesEndpoint(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func locationEndpoint(c *gin.Context) {
  location := c.Param("location")
  c.String(http.StatusOK, location)
}

func adminEndpoint(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func adminRequired() gin.HandlerFunc {
  return func(c *gin.Context) {
    log.Println("ok")
  }
}
