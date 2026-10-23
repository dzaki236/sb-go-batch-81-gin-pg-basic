package handlers

import (
	"rest-api/repositories"
	. "rest-api/structs"

	"github.com/gin-gonic/gin"
)

func DeleteBioskopHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := repositories.DeleteBioskop(id); err != nil {
		ctx.JSON(500, ApiResponse{
			Message: "Gagal menghapus data",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(201, ApiResponse{
		Message: "Berhasil menghapus bioskop",
		Data:    nil,
	})
}
