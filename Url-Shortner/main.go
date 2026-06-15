package main

import (
	"Url-Shortner/models"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
)

var db *gorm.DB
var baseURL string


// initDB initializes the database connection and migrates the URL model
func initDB(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL = os.Getenv("BASE_URL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.URL{})
	db = database
}

// GenarateShortCode generates a random short code of specified length

func GenarateShortCode(n int) string {

	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Failed to generate short code:", err)
	}
	
	return base64.URLEncoding.EncodeToString(b)[:n]
}

// CreateShortURL handles the creation of a short URL from the original URL

func CreateShortURL(c echo.Context) error {

	type Request struct {
		URL string `json:"url"`
	}
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	shortCode := GenarateShortCode(6)

	url := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
	}
	
	if err := db.Create(&url).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create short URL"})
	}

	return c.JSON(http.StatusOK, echo.Map{"short_url": fmt.Sprintf("%s/%s", baseURL, shortCode)})

}
// Redirect handles the redirection from short URL to original URL

func Redirect(c echo.Context) error {
	shortCode := c.Param("shortCode")
	var url models.URL
	if err := db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Short URL not found"})
	}

	db.Model(&url).Update("clicks", url.Clicks+1)



	url.Clicks++
	db.Save(&url)

	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

// Stats returns the original URL and click count for a given short code

func Stats(c echo.Context) error {
	shortCode := c.Param("shortCode")
	var url models.URL
	if err := db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Short URL not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"original_url": url.OriginalURL,
		"short_code":   url.ShortCode,
		"clicks":       url.Clicks,
	})
}


func main() {
	initDB()
	e:= echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.POST("/shorten", CreateShortURL)
	e.GET("/:shortCode", Redirect)
	e.GET("stats/:shortCode", Stats)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))


}

