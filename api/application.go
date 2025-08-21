package api

import (
	"net/http"

	db "github.com/Llala/simplecat/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createApplicationRequest struct {
	Name       string `json:"name" binding:"required"`
	SourceText string `json:"source_text" binding:"required"`
}

func (server *Server) createApplication(ctx *gin.Context) {
	var req createApplicationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.SourceUnitParams{
		Name: req.Name,
		Text: req.SourceText,
	}

	application, err := server.store.ParseTextTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, application)
}

type listApplicationRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListApplication(ctx *gin.Context) {
	var req listApplicationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListApplicationsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	applications, err := server.store.ListApplications(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, applications)
}

type DeleteApplicationRequest struct {
	ID int64 `form:"id" binding:"required"`
}

func (server *Server) DeleteApplication(ctx *gin.Context) {
	var req DeleteApplicationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteApplication(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, err)
}
