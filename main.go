package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    _ "go-api-sample/docs"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)


// @Description Restaurant information
// @Description with name, city, location, instagram link and a short description
type restaurant struct {
    Name     string  `json:"name"`
    City  string  `json:"city"`
    Location string  `json:"location"`
	Instagram string  `json:"instagram"`
    Description  string `json:"description"`
}

// Example of data
var restaurants = []restaurant{
	  {Name: "Livio", Description: "Italian restaurant", Location: "6 rue de Longchamp, 92200 Neuilly sur Seine", Instagram: "https://www.instagram.com/chezlivio/?hl=en", City: "paris"},
	  {Name: "The Ivy Chelsea garden", Description: "European food", Location: "195 197 King's Rd, London SW3 5EQ", City: "london", Instagram: "https://www.instagram.com/theivychelseagarden/?hl=en"},
}


// @title Go API sample
// @version 1.0
// @description This is a sample API

// @contact.name Kim

// @host localhost:8080
// @BasePath /
func main() {
    router := gin.Default()

    // routes of the API
    router.GET("/restaurants", getRestaurants)
    router.GET("/restaurants/:city", getRestaurantsByCity)
    router.POST("/restaurants", createRestaurants)

    // swagger route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    // run Gin server
    router.Run("localhost:8080")
}


// @Summary List all restaurants defined
// @ID get-restaurants
// @Schemes
// @Tags restaurants
// @Produce json
// @Success 200 {array} restaurant
// @Router /restaurants [get]
func getRestaurants(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, restaurants)
}


// @Summary Add a new restaurant
// @ID create-restaurant
// @Produce json
// @Param restaurant body restaurant true "Restaurant"
// @Success 200 {array} restaurant
// @Failure 400  "an error occurred while creating restaurant"
// @Router /restaurants [post]
func createRestaurants(c *gin.Context) {
    var newRestaurant restaurant

    if err := c.BindJSON(&newRestaurant); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "an error occurred while creating restaurant"})
        return
    }

    restaurants = append(restaurants, newRestaurant)
    c.IndentedJSON(http.StatusCreated, newRestaurant)
}


// @Summary List all restaurants located in given city
// @ID get-restaurants-by-city
// @Schemes
// @Tags restaurants
// @Param city path string true "City"
// @Produce json
// @Success 200 {object} restaurant
// @Failure 404 "restaurants not found"
// @Router /restaurants/{city} [get]
func getRestaurantsByCity(c *gin.Context) {
    city := c.Param("city")

    for _, a := range restaurants {
        if a.City == city {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "restaurants not found"})
}