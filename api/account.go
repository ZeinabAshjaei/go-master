package api

import (
	"database/sql"
	db "github.com/ZeinabAshjaei/go-master/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server Server) createAccount(context *gin.Context) {
	var req CreateAccountRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(context, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				context.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server Server) getAccount(context *gin.Context) {
	var req getAccountRequest

	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	account, err := server.store.GetAccount(context, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	context.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (server Server) listAccount(context *gin.Context) {
	var req listAccountRequest

	if err := context.ShouldBindQuery(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListAccountParams{
		Limit:  req.PageSize,
		Offset: req.PageID,
	}
	accounts, err := server.store.ListAccount(context, arg)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, accounts)
}

type deleteAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server Server) deleteAccount(context *gin.Context) {
	var req deleteAccountRequest

	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := server.store.DeleteAccount(context, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, db.Account{})
}
