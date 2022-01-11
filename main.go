// Recipes API
//
// description: "This is a sample recipes API. Youcan find out more about the API at"
//
// 	Schemes: http
// 	Host: localhost:8080
// 	BasePath: /
// 	Version: 1.0.0
// 	Contact: Wagner Ricardo Wagner<wagnerricardonet@gmail.com>
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/wagnerww/go-gin-recipes-api.git/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection

var recipesHandler *handlers.RecipesHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("recipes")

	log.Println("Connected to MongoDB")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping()
	fmt.Println(status)

	recipesHandler = handlers.NewRecipesHandler(ctx,
		collection, redisClient)
}

func NewRecipeHandler(c *gin.Context) {
	recipesHandler.NewRecipeHandler(c)
}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	recipesHandler.ListRecipesHandler(c)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the recipe
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
//     '404':
//         description: Invalid recipe ID
func UpdateRecipeHandler(c *gin.Context) {
	recipesHandler.UpdateRecipeHandler(c)
}

func DeleteRecipeHandler(c *gin.Context) {
	/*	id := c.Param("id")

		index := -1
		for i := 0; i < len(recipes); i++ {
			if recipes[i].ID == id {
				index = i
			}
		}

		if index == -1 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Recipe not found",
			})
			return
		}

		recipes = append(recipes[:index], recipes[index+1:]...)
		c.JSON(http.StatusOK, gin.H{
			"message": "Receipe has been deleted",
		})*/
}

func SearchRecipesHandler(c *gin.Context) {
	/*	tag := c.Query("tag")
		listOfRecipes := make([]Recipe, 0)
		for i := 0; i < len(recipes); i++ {
			found := false
			for _, t := range recipes[i].Tags {
				if strings.EqualFold(t, tag) {
					found = true
				}
			}

			if found {
				listOfRecipes = append(listOfRecipes, recipes[i])
			}
		}

		c.JSON(http.StatusOK, listOfRecipes)
	*/
}

func main() {
	router := gin.Default()

	router.GET("/recipes", ListRecipesHandler)

	authorized := router.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.POST("/recipes", NewRecipeHandler)

		authorized.PUT("/recipes/:id", UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", DeleteRecipeHandler)
		authorized.GET("/recipes/search", SearchRecipesHandler)
	}
	router.Run()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != os.Getenv("X_API_KEY") {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}
