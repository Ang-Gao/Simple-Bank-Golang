# Simple-Bank-Golang
```
package api

import (
	"net/http"

	"log"
	db "myproject/db/sqlc"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    "Angx",
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		log.Fatal("Cannot create account:", err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}
```
