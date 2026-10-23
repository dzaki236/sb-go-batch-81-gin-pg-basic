package repositories

import (
	"database/sql"
	"errors"
	"rest-api/config"
	. "rest-api/structs"
	"strconv"
)

func CreateBioskop(bioskop Bioskop) (Bioskop, error) {
	query := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := config.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID)
	if err != nil {
		return Bioskop{}, err
	}

	return bioskop, nil
}

func GetBioskop() ([]Bioskop, error) {
	query := `
		SELECT id, nama, lokasi, rating
		FROM bioskop
		ORDER BY id ASC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bioskopData := []Bioskop{}
	for rows.Next() {
		var bioskop Bioskop
		if err := rows.Scan(
			&bioskop.ID,
			&bioskop.Nama,
			&bioskop.Lokasi,
			&bioskop.Rating,
		); err != nil {
			return nil, err
		}
		bioskopData = append(bioskopData, bioskop)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bioskopData, nil
}

func GetDetailBioskop(id string) (Bioskop, error) {
	id_baru, _ := strconv.Atoi(id)

	query := `
		SELECT id, nama, lokasi, rating
		FROM bioskop
		Where id = $1 LImit 1
	`
	var bioskop Bioskop

	err := config.DB.QueryRow(query, id_baru).Scan(
		&bioskop.ID,
		&bioskop.Nama,
		&bioskop.Lokasi,
		&bioskop.Rating)

	if err != nil || err == sql.ErrNoRows {
		return bioskop, errors.New("Data tidak tersedia")
	}
	return bioskop, nil
}

func UpdateBioskop(id string, bioskop Bioskop) (Bioskop, error) {
	id_baru, _ := strconv.Atoi(id)
	query := `
			UPDATE bioskop
			SET nama = $1, lokasi = $2, rating = $3
			WHERE id = $4
		`

	result, err := config.DB.Exec(
		query,
		bioskop.Nama,
		bioskop.Lokasi,
		bioskop.Rating,
		id_baru,
	)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return Bioskop{}, err
	}
	return bioskop, nil
}

func DeleteBioskop(id string) error {
	id_baru, _ := strconv.Atoi(id)
	query := `DELETE FROM bioskop WHERE id = $1`
	result, err := config.DB.Exec(query, id_baru)
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("Data tersebut sudah tidak tersedia")
	}
	return nil

}
