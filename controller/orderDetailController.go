package controller

import (
	"challenge-godb/config"
	"challenge-godb/model"
	"database/sql"
)

func CreateOrderDetail(tx *sql.Tx, orderDetail *model.OrderDetail) (string, error) {

	query := "INSERT INTO trx_order_detail (order_id, layanan_id, quantity) VALUES ($1, $2, $3) RETURNING id_order_detail"
	err := tx.QueryRow(query, orderDetail.Order_Id, orderDetail.Layanan_Id, orderDetail.Quantity).Scan(&orderDetail.Id_Order_Detail)
	if err != nil {
		panic(err)
	}
	return orderDetail.Id_Order_Detail, nil
}

func GetTotalHarga(db *sql.DB, orderID int) (int, error) {
	var totalHarga int

	rows, err := db.Query("SELECT d.quantity * l.harga FROM trx_order_detail d JOIN mst_layanan l ON d.layanan_id = l.id_layanan WHERE d.order_id = $1", orderID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var hargaLayanan int
		err := rows.Scan(&hargaLayanan)
		if err != nil {
			return 0, err
		}
		totalHarga += hargaLayanan
	}

	return totalHarga, nil
}

func GetOrderDetailByOrderId(id string) []model.OrderDetail {
	if id == "" {
		// Atau sesuaikan dengan penanganan kesalahan yang sesuai.
		panic("ID pesanan tidak valid")
	}
	db := config.ConnectDb()
	defer db.Close()

	sqlQuery := "SELECT * FROM trx_order_detail WHERE order_id = $1"

	rows, err := db.Query(sqlQuery, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	orderDetails := ScanOrderDetail(rows)

	return orderDetails
}

func ScanOrderDetail(rows *sql.Rows) []model.OrderDetail {
	orderDetails := []model.OrderDetail{}
	var err error

	for rows.Next() {
		orderDetail := model.OrderDetail{}
		err = rows.Scan(&orderDetail.Id_Order_Detail, &orderDetail.Order_Id, &orderDetail.Layanan_Id, &orderDetail.Quantity)
		if err != nil {
			panic(err)
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orderDetails
}
