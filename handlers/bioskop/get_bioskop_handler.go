package handlers

import (
	"rest-api/repositories"
	. "rest-api/structs"

	"github.com/gin-gonic/gin"
)

func GetBioskopHandler(ctx *gin.Context) {
	bioskopData, err := repositories.GetBioskop()
	if err != nil {
		ctx.JSON(500, ApiResponse{
			Message: "Gagal mengambil data bioskop",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(200, ApiResponse{
		Message: "Berhasil mendapatkan data",
		Data:    bioskopData,
	})

}

func GetDetailBioskopHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	bioskopData, err := repositories.GetDetailBioskop(id)
	if err != nil {
		ctx.JSON(500, ApiResponse{
			Message: "Gagal mengambil data bioskop",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(200, ApiResponse{
		Message: "Berhasil mendapatkan data",
		Data:    bioskopData,
	})

}
