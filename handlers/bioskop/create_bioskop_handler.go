package handlers

import (
	"strings"

	"rest-api/repositories"
	. "rest-api/structs"

	"github.com/gin-gonic/gin"
)

func CreateBioskopHandler(ctx *gin.Context) {
	var bioskopRequest BioskopRequest
	err := ctx.ShouldBindJSON(&bioskopRequest)

	bioskopRequest.Nama = strings.TrimSpace(bioskopRequest.Nama)
	bioskopRequest.Lokasi = strings.TrimSpace(bioskopRequest.Lokasi)

	if bioskopRequest.Nama == "" || bioskopRequest.Lokasi == "" {
		ctx.JSON(400, ApiResponse{
			Message: "Nama dan lokasi wajib diisi",
			Data:    nil,
		})
		return
	}

	if bioskopRequest.Rating < 0 || bioskopRequest.Rating > 5 {
		ctx.JSON(400, ApiResponse{
			Message: "Rating harus antara 0 sampai 5",
			Data:    nil,
		})
		return
	}

	var bioskopModel Bioskop = Bioskop{
		Nama:   bioskopRequest.Nama,
		Lokasi: bioskopRequest.Lokasi,
		Rating: bioskopRequest.Rating,
	}
	bioskopModel, err = repositories.CreateBioskop(bioskopModel)

	if err != nil {
		ctx.JSON(500, ApiResponse{
			Message: "Gagal melakukan menambahkan data",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(201, ApiResponse{
		Message: "Berhasil menambahkan bioskop",
		Data:    bioskopModel,
	})
}
