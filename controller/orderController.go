package controller

import (
	"challenge-godb/config"
	"challenge-godb/model"
	"database/sql"
	"time"
)

func CreateOrder(tx *sql.Tx, order *model.Order) error {
	tanggalMasuk := time.Now()

	query := "INSERT INTO trx_order (customer_id, tanggal_masuk, tanggal_keluar) VALUES ($1, $2, $3) RETURNING id_order"
	err := tx.QueryRow(query, order.Customer_Id, tanggalMasuk, nil).Scan(&order.Id_Order)

	if err != nil {
		panic(err)
	}

	order.Tanggal_Masuk = tanggalMasuk
	// order.Tanggal_Keluar = tanggalKeluar
	return nil
}

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func GetOrderByCustomerId(id string) model.Order {
	db := config.ConnectDb()
	defer db.Close()
	var err error

	sqlQuery := "SELECT id_order, customer_id, tanggal_masuk FROM trx_order WHERE customer_id = $1 AND tanggal_keluar IS NULL"

	order := model.Order{}
	err = db.QueryRow(sqlQuery, id).Scan(&order.Id_Order, &order.Customer_Id, &order.Tanggal_Masuk)

	if err == sql.ErrNoRows {
		return model.Order{}
	} else if err != nil {
		panic(err)
	}

	return order
}

func GetOrderByOrderId(id string) model.Order {
	db := config.ConnectDb()
	defer db.Close()
	var err error

	sqlQuery := "SELECT * FROM trx_order WHERE id_order = $1"

	order := model.Order{}
	err = db.QueryRow(sqlQuery, id).Scan(&order.Id_Order, &order.Customer_Id, &order.Tanggal_Masuk, &order.Tanggal_Keluar, &order.Penerima)
	if err == sql.ErrNoRows {
		return model.Order{}
	} else if err != nil {
		panic(err)
	}

	return order
}

func ScanOrder(rows *sql.Rows) []model.Order {
	orders := []model.Order{}
	var err error

	for rows.Next() {
		order := model.Order{}
		err := rows.Scan(&order.Id_Order, &order.Customer_Id, &order.Tanggal_Masuk)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orders
}

func UpdateOrder(tx *sql.Tx, order *model.Order) error {
	updateOrder := "UPDATE trx_order SET tanggal_keluar = $2, penerima = $3 WHERE id_order = $1"

	_, err := tx.Exec(updateOrder, order.Id_Order, order.Tanggal_Keluar, order.Penerima)
	if err != nil {
		panic(err)
	}

	return nil
}

func FormatDatabaseDate(tanggal time.Time) string {
	return tanggal.Format("2006-01-02")
}
