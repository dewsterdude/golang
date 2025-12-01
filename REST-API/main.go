package main

// http package that allows sending or handling HTTP requests
// for rest api which will have advanced routing... will use gin - by go community
// go get -u github.com/gin-gonic/gin
// installed sqllite3  go get github.com/mattn/go-sqlite3

import (
	// import the sqlite3 driver

	"example.com/REST-API/db"
	"example.com/REST-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// setup DB
	db.InitDB()

	// setup server
	server := gin.Default()       // creates a router with default middleware: logger and recovery (crash-free) middleware
	routes.RegisterRoutes(server) // server is a pointer
	// setup handler for HTTPS requests - but need the PATH
	server.Run(":8080") // localhost:8080 - port to listen on to rewquets

}
