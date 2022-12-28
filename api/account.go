package api

import (
	"net/http"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"udemy/restfulapi/simplebank/api"
)

type createAccountRequest struct {
	// remove this balance field.We only allow clients to specify the owner’s name and the currency of the account.
	Owner    string `json:"owner" binding: "required"`
	Currency string `json:"currency" binding:"required, oneof=USD EUR"`
}

func (server *Server) createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {  
		// it the error is not nil, it means the client has provided invalid data.
		//first, give a json response to client. first is status code, second is JSON object that we want to send to client.
			// need to convert this err into a key-value object so that Gin can serialize it to JSON before returning to the client.
		c.JSON(http.StatusBadRequest, errorResponse(err)) // errorResponse() will used in various so implement in the server.go
		return 
	}
	//if the input data has no error. So we just go ahead to insert a new account into the database
	arg := db.CreateAccountParams{
		Owner : req.Owner, 
		Currency: req.Currency,
		Balance: 0,
	}
	account, err := server.store.CreateAccount(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, account)
}