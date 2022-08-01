package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Data structure
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

func main() {
    router := gin.Default()
    router.GET("/restaurants", getRestaurants)
    router.GET("/restaurants/:city", getRestaurantsByCity)
    router.POST("/restaurants", postRestaurants)

    router.Run("localhost:8080")
}

// getRestaurants responds with the list of all restaurants as JSON
func getRestaurants(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, restaurants)
}

// postRestaurants adds a restaurant from JSON received in the request body
func postRestaurants(c *gin.Context) {
    var newRestaurant restaurant

    if err := c.BindJSON(&newRestaurant); err != nil {
        return
    }

    restaurants = append(restaurants, newRestaurant)
    c.IndentedJSON(http.StatusCreated, newRestaurant)
}

// getRestaurantsByCity locates the restaurants whose city value matches the parameter city
// parameter sent by the client, then returns that restaurants as a response
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