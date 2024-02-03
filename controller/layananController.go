package controller

import (
	"challenge-godb/config"
	"challenge-godb/model"
	"database/sql"
	"fmt"
)

func AddLayanan(layanan model.Layanan) error {
	db := config.ConnectDb()
	defer db.Close()

	if layanan.Nama_Layanan == "" || layanan.Harga == 0 || layanan.Satuan == "" {
		return fmt.Errorf("field tidak boleh kosong")
	}
	if layanan.Harga < 0 {
		return fmt.Errorf("harga layanan tidak boleh negatif")
	}

	queryInsert := "INSERT INTO mst_layanan (id_layanan, nama_layanan, harga, satuan) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(queryInsert, layanan.Id_Layanan, layanan.Nama_Layanan, layanan.Harga, layanan.Satuan)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully Inserted Data!")
	}

	return nil
}

func UpdateLayanan(layanan model.Layanan) error {
	db := config.ConnectDb()
	defer db.Close()
	var err error

	if layanan.Nama_Layanan == "" || layanan.Harga == 0 || layanan.Satuan == "" {
		return fmt.Errorf("field tidak boleh kosong")
	}
	if layanan.Harga < 0 {
		return fmt.Errorf("harga layanan tidak boleh negatif")
	}

	updateLayanan := "UPDATE mst_layanan SET nama_layanan = $2, harga = $3, satuan = $4 WHERE id_layanan = $1"
	_, err = db.Exec(updateLayanan, layanan.Id_Layanan, layanan.Nama_Layanan, layanan.Harga, layanan.Satuan)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}

	return nil
}

func DeleteLayanan(id string) error {
	db := config.ConnectDb()
	defer db.Close()

	hasTransactions, err := HasLayananTransactions(id)
	if err != nil {
		return err
	}
	if hasTransactions {
		return fmt.Errorf("Layanan dengan ID %s memiliki transaksi terkait dan tidak dapat dihapus", id)
	}

	query := "DELETE FROM mst_layanan WHERE id_layanan = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully deleted data!")
	}

	return nil
}

func GetAllLayanan() []model.Layanan {
	db := config.ConnectDb()
	defer db.Close()

	sqlQuery := "SELECT * FROM mst_layanan"

	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	layanans := ScanLayanan(rows)

	return layanans
}

func ScanLayanan(rows *sql.Rows) []model.Layanan {
	layanans := []model.Layanan{}
	var err error

	for rows.Next() {
		layanan := model.Layanan{}
		err := rows.Scan(&layanan.Id_Layanan, &layanan.Nama_Layanan, &layanan.Harga, &layanan.Satuan)
		if err != nil {
			panic(err)
		}
		layanans = append(layanans, layanan)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return layanans
}

func GetLayananById(id string) model.Layanan {
	db := config.ConnectDb()
	defer db.Close()
	var err error

	sqlQuery := "SELECT * FROM mst_layanan WHERE id_layanan = $1"

	layanan := model.Layanan{}
	err = db.QueryRow(sqlQuery, id).Scan(&layanan.Id_Layanan, &layanan.Nama_Layanan, &layanan.Harga, &layanan.Satuan)
	if err == sql.ErrNoRows {
		return model.Layanan{}
	} else if err != nil {
		panic(err)
	}

	return layanan
}

func HasLayananTransactions(id string) (bool, error) {
	db := config.ConnectDb()
	defer db.Close()

	query := "SELECT COUNT(*) FROM trx_order_detail WHERE layanan_id = $1"
	var count int
	err := db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
