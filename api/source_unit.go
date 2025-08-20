package api

import (
	"net/http"

	db "github.com/Llala/simplecat/db/sqlc"
	"github.com/gin-gonic/gin"
)

type listSourceUnitRequest struct {
	ApplicationID int32 `form:"application_id" binding:"required"`
	PageID        int32 `form:"page_id" binding:"required,min=1"`
	PageSize      int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListTextUnits(ctx *gin.Context) {
	var req listSourceUnitRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSourceUnitJoinParams{
		ApplicationID: req.ApplicationID,
		Limit:         req.PageSize,
		Offset:        (req.PageID - 1) * req.PageSize,
	}

	textUnits, err := server.store.ListSourceUnitJoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, textUnits)
}
