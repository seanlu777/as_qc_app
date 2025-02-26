package main

import (
	"as_qc_app/db"
	"as_qc_app/internal/api"
	"as_qc_app/parsers"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var wg sync.WaitGroup

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err != nil {
		log.Fatalf("Could not get local IP: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Initialize database connection
	err = db.Initialize(dsn) // Call your initialization function here
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening database", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Successfully connected to the database!")
	startHTTPServer()
}

// HTTP Server
func startHTTPServer() {
	defer wg.Done()
	// Set Gin to release mode to minimize logging
	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin router
	router := gin.Default()

	// Configure trusted proxies

	if err := router.SetTrustedProxies([]string{"0.0.0.0/0"}); err != nil {
		panic("Failed to set trusted proxies: " + err.Error())
	}

	// Seal tracking
	router.GET("/api/health", apiHealth)
	router.POST("/api/pushRecord", pushRecord)
	router.GET("/api/getLatestRecordList", getLatestRecordList)
	router.GET("/api/getHistoryData", getHistoryData)
	// Start the server on the port 8080
	apiPort := os.Getenv("API_PORT")
	// Log the local IP and port before starting the server
	fmt.Printf("HTTP server listening on port %s\n", apiPort)
	if err := router.Run("0.0.0.0:" + apiPort); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}

func apiHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func pushRecord(c *gin.Context) {
	var req api.PushRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	a2tbData, err := parsers.ParseA2TB(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err := db.SaveA2TB(a2tbData); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func getLatestRecordList(c *gin.Context) {
	var req api.GetLatestRecordListRequest
	req.StartAt = c.Query("startAt")
	req.EndAt = c.Query("endAt")
	req.Station = c.Query("station")

	latestDataList, err := db.GetLatestRecordList(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": latestDataList})
}

func getHistoryData(c *gin.Context) {
	var req api.GetHistoryDataRequest
	startAt, err := time.Parse(time.RFC3339, c.Query("startAt"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.StartAt = startAt
	endAt, err := time.Parse(time.RFC3339, c.Query("endAt"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.EndAt = endAt
	req.TagId = c.Query("tagId")
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	historyDataList, err := db.GetHistoryData(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": historyDataList})
}
