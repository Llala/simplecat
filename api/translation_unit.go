package api

import (
	"database/sql"
	"net/http"

	db "github.com/Llala/simplecat/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type updateTranslationUnitRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Text string `json:"text" binding:"required"`
}

func (server *Server) updateTranslationUnit(ctx *gin.Context) {
	var req updateTranslationUnitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTranslationUnitParams{
		ID: req.ID,
		Text: sql.NullString{
			String: req.Text,
			Valid:  true,
		},
	}

	translationUnit, err := server.store.UpdateTranslationUnit(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, translationUnit)
}

type listTranslationRequest struct {
	ApplicationID int64 `form:"application_id" binding:"required"`
}

func (server *Server) GetTranslation(ctx *gin.Context) {
	var req listTranslationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TranslationUnitFormParams{
		ApplicationID: req.ApplicationID,
	}

	application, err := server.store.FormTextTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, application)
}
