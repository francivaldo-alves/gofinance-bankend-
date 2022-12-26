package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/francivaldo-alves/gofinance-bankend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	CategoryID  int32     `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Value       int32     `json:"value" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

// Funcação da PI para cadastar um account
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}
	var categoryId = req.CategoryID
	var accountType = req.Type
	category, err := server.store.GetCategory(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	var categoryTypeIsDiffentOdAccountType = category.Type != accountType
	if categoryTypeIsDiffentOdAccountType {
		ctx.JSON(http.StatusBadRequest, "Account type is diferente od Category type")
	} else {

		arg := db.CreateAccountParams{
			UserID:      req.UserID,
			CategoryID:  categoryId,
			Title:       req.Title,
			Type:        accountType,
			Description: req.Description,
			Value:       req.Value,
			Date:        req.Date,
		}
		account, err := server.store.CreateAccount(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusOK, account)
	}

}

// Funcação da PI para buscar uma account
type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {

	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id" binding:"required" `
	CategoryID  int32     `json:"category_id" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Title       string    `json:"title" `
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}

// Funcação da PI para buscar uma account
func (server *Server) getAccounts(ctx *gin.Context) {
	var req getAccountsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	var accounts interface{}

	var parameterHasUserIdType = req.UserID > 0 && len(req.Type) > 0

	FilterAsByUserIdAndType := req.CategoryID == 0 && len(req.Date.GoString()) == 0 &&
		len(req.Description) == 0 && len(req.Title) == 0 && parameterHasUserIdType
	if FilterAsByUserIdAndType {

		arg := db.GetAccountsByUserIdAndTypeParams{
			UserID: req.UserID,
			Type:   req.Type,
		}

		accountsByUserIdAndType, err := server.store.GetAccountsByUserIdAndType(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndType
	}
	FilterAsByUserIdAndTypeAndCategoryId := req.CategoryID != 0 && len(req.Date.GoString()) == 0 &&
		len(req.Description) == 0 && len(req.Title) == 0 && parameterHasUserIdType
	if FilterAsByUserIdAndTypeAndCategoryId {

		arg := db.GetAccountsByUserIdAndTypeAndCategoryIdParams{
			UserID:     req.UserID,
			Type:       req.Type,
			CategoryID: req.CategoryID,
		}

		accountsByUserIdAndTypeAndCategoryId, err := server.store.GetAccountsByUserIdAndTypeAndCategoryId(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndTypeAndCategoryId
	}
	FilterAsByUserIdAndTypeAndCategoryIdAndTitle := req.CategoryID != 0 && len(req.Date.GoString()) == 0 &&
		len(req.Description) == 0 && len(req.Title) > 0 && parameterHasUserIdType
	if FilterAsByUserIdAndTypeAndCategoryIdAndTitle {

		arg := db.GetAccountsByUserIdAndTypeAndCategoryIdAndTitleParams{
			UserID:     req.UserID,
			Type:       req.Type,
			CategoryID: req.CategoryID,
			Title:      req.Title,
		}

		accountsByUserIdAndTypeAndCategoryIdAndTitle, err := server.store.GetAccountsByUserIdAndTypeAndCategoryIdAndTitle(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndTypeAndCategoryIdAndTitle
	}
	FilterAsByUserIdAndTypeAndCategoryIdAndTitleAndDescription := req.CategoryID != 0 && len(req.Date.GoString()) == 0 &&
		len(req.Description) > 0 && len(req.Title) > 0 && parameterHasUserIdType
	if FilterAsByUserIdAndTypeAndCategoryIdAndTitleAndDescription {

		arg := db.GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionParams{
			UserID:      req.UserID,
			Type:        req.Type,
			CategoryID:  req.CategoryID,
			Title:       req.Title,
			Description: req.Description,
		}

		accountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription, err := server.store.GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription
	}
	FilterAsByUserIdAndData := req.CategoryID == 0 && len(req.Date.GoString()) > 0 &&
		len(req.Description) == 0 && len(req.Title) == 0 && parameterHasUserIdType
	if FilterAsByUserIdAndData {

		arg := db.GetAccountsByUserIdAndTypeIdAndDateParams{
			UserID: req.UserID,
			Type:   req.Type,
			Date:   req.Date,
		}

		accountsByUserIdAndDate, err := server.store.GetAccountsByUserIdAndTypeIdAndDate(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndDate
	}
	FilterAsByUserIdAndDescription := req.CategoryID == 0 && len(req.Date.GoString()) == 0 &&
		len(req.Description) > 0 && len(req.Title) == 0 && parameterHasUserIdType
	if FilterAsByUserIdAndDescription {

		arg := db.GetAccountsByUserIdAndTypeIdAndDescriptionParams{
			UserID:      req.UserID,
			Type:        req.Type,
			Description: req.Description,
		}

		accountsByUserIdAndDescription, err := server.store.GetAccountsByUserIdAndTypeIdAndDescription(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndDescription
	}
	FilterAsByUserIdAndTitle := req.CategoryID == 0 && len(req.Date.GoString()) == 0 &&
		len(req.Description) == 0 && len(req.Title) > 0 && parameterHasUserIdType
	if FilterAsByUserIdAndTitle {

		arg := db.GetAccountsByUserIdAndTypeIdAndTitleParams{
			UserID: req.UserID,
			Type:   req.Type,
			Title:  req.Title,
		}

		accountsByUserIdAndTitle, err := server.store.GetAccountsByUserIdAndTypeIdAndTitle(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndTitle
	}
	FilterAsParameters := req.CategoryID > 0 && len(req.Date.GoString()) > 0 &&
		len(req.Description) > 0 && len(req.Title) > 0 && parameterHasUserIdType
	if FilterAsParameters {

		arg := db.GetAccountsParams{
			UserID:      req.UserID,
			Type:        req.Type,
			Title:       req.Title,
			CategoryID:  req.CategoryID,
			Description: req.Description,
			Date:        req.Date,
		}

		accountsFilterAsParameters, err := server.store.GetAccounts(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accounts = accountsByUserIdAndTitle
	}

	ctx.JSON(http.StatusOK, accounts)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" `
	Title       string `json:"title" `
	Description string `json:"description" `
	Value       int32  `json:"value"`
}

// Funcação da PI para cadastar uma categoria
func (server *Server) updateAccount(ctx *gin.Context) {
	var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}
	arg := db.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}
	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

type deleteAccountByRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

// Funcação da PI para deletar uma categoria
func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountByRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, true)
}

type GetAccountsGraphParams struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountGraph(ctx *gin.Context) {

	var req GetAccountsGraphParams
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	countGraph, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, countGraph)
}

type GetAccountsReportsParams struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountRports(ctx *gin.Context) {

	var req GetAccountsReportsParams
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	sumReports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sumReports)
}
