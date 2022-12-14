package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter()
	initaliseHandlers(router)

}
func CORS() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	return headers, methods, origins
}
func initaliseHandlers(router *mux.Router) {
	var headers, methods, origins handlers.CORSOption = CORS()
	router.HandleFunc("/form", controllers.Create).Methods("POST")
	router.HandleFunc("/get", controllers.Getall)
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(router))

}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "my-secret-pw",
			DB:         "aws",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}
}
