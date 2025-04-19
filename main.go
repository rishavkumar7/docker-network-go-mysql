package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dbString)
	if err != nil {
		fmt.Printf("Error while connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)`)
	if err != nil {
		fmt.Printf("Error while creating table: %v\n", err)
		return
	}

	router := gin.Default()
	router.POST("/add", func(c *gin.Context) {
		addUser(c, db)
	})
	router.Run(":8500")
}

func addUser(c *gin.Context, db *sql.DB) {
	var user struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"success": false,
		})
		return
	}
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while inserting new user",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User added successfully",
		"success": true,
	})
}
