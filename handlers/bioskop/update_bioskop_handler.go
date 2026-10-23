package handlers

import (
	"rest-api/repositories"
	. "rest-api/structs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateBioskopHandler(ctx *gin.Context) {
	id := ctx.Param("id")
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
	id_baru, _ := strconv.Atoi(id)
	var bioskopModel Bioskop = Bioskop{
		ID:     id_baru,
		Nama:   bioskopRequest.Nama,
		Lokasi: bioskopRequest.Lokasi,
		Rating: bioskopRequest.Rating,
	}

	bioskopModel, err = repositories.UpdateBioskop(id, bioskopModel)
	if err != nil {
		ctx.JSON(500, ApiResponse{
			Message: "Gagal melakukan perubahan data",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(201, ApiResponse{
		Message: "Berhasil melakukan perubahan bioskop",
		Data:    bioskopModel,
	})
}
