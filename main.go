// Recipes API
//
// description: "This is a sample recipes API. Youcan find out more about the API at"
//
// 	Schemes: http
// 	Host: localhost:8080
// 	BasePath: /
// 	Version: 1.0.0
// 	Contact: Wagner Ricardo Wagner<wagnerricardonet@gmail.com>
//  SecurityDefinitions:
//  api_key:
//    type: apiKey
//    name: Authorization
//    in: header
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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/wagnerww/go-gin-recipes-api.git/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler
var authHandler *handlers.AuthHandler

func init() {

	profile := os.Getenv("PROFILE")
	var env = "./env/.env"
	if profile != "" {
		env = env + "." + profile
	}

	if err := godotenv.Load(env); err != nil {
		log.Printf("No .env file found")
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("recipes")

	collectionUsers := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("users")

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
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)

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

func SetupServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	/* PERSONALIZADO

	router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost
													:3000"},
			AllowMethods:     []string{"GET", "OPTIONS"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		}))

	*/

	router.Use(gin.LoggerWithFormatter(func(
		param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s\n",
			param.TimeStamp.Format("2006-01-02T15:04:05"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	router.GET("/recipes", ListRecipesHandler)
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/signup", authHandler.SignUpHandler)
	router.POST("/refresh", authHandler.RefreshHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", NewRecipeHandler)

		authorized.PUT("/recipes/:id", UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", DeleteRecipeHandler)
		authorized.GET("/recipes/search", SearchRecipesHandler)
	}
	return router
}

func main() {
	SetupServer().Run()
}
