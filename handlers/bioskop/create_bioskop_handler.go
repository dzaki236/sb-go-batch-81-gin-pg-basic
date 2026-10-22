package handlers

import (
	"database/sql"
	"log"
	"strings"

	"rest-api/structs"
	. "rest-api/structs"

	"github.com/gin-gonic/gin"
)

func CreateBioskopHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bioskopRequest BioskopRequest

		if err := ctx.ShouldBindJSON(&bioskopRequest); err != nil {
			ctx.JSON(400, structs.ApiResponse{
				Message: "Error Create Data",
				Data:    nil,
			})
			return
		}

		bioskopRequest.Nama = strings.TrimSpace(bioskopRequest.Nama)
		bioskopRequest.Lokasi = strings.TrimSpace(bioskopRequest.Lokasi)

		if bioskopRequest.Nama == "" || bioskopRequest.Lokasi == "" {
			ctx.JSON(400, structs.ApiResponse{
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

		query := `
			INSERT INTO bioskop (nama, lokasi, rating)
			VALUES ($1, $2, $3)
			RETURNING id
		`

		var bioskop Bioskop
		bioskop.Nama = bioskopRequest.Nama
		bioskop.Lokasi = bioskopRequest.Lokasi
		bioskop.Rating = bioskopRequest.Rating

		err := db.QueryRow(
			query,
			bioskop.Nama,
			bioskop.Lokasi,
			bioskop.Rating,
		).Scan(&bioskop.ID)

		if err != nil {
			log.Println("Gagal menambahkan bioskop:", err)
			ctx.JSON(500, ApiResponse{
				Message: "Gagal menambahkan bioskop",
				Data:    nil,
			})
			return
		}

		ctx.JSON(201, ApiResponse{
			Message: "Berhasil menambahkan bioskop",
			Data:    bioskop,
		})
	}
}
