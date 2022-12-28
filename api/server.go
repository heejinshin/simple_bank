// This is where we implement our HTTP API server.
package api 

import (
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store *db.Store   // we have to access to store object to have a new account
	router *gin.Engine //help each api to corect handler
} 

//New Server creates a new HTTP server and setup routing.
func NewServer(store *db.Store) * Server { // we take the db and return a server 
	server := &Server{store:store}
	router := gin.Default()

	// add routes to router 
	router.POST("/accounts", server.createAccount)  // for the create Account, make the method in server struct

	server.router = router 
	return server 
}

func errorResponse(err error) gin.H { // This function will take an error as input, and it will return a gin.H; infact just a shortcut for map[string]interface{}, so we can store whatever key-value data that we want in it. 
	return gin.H{"error": err.Error()}  // Error() makes err to better form 
}
