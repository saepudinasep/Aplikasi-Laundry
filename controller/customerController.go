package controller

import (
	"challenge-godb/model"
	"database/sql"
	"fmt"
)

func AddCustomer(customer model.Customer, tx *sql.Tx) error {
	if customer.Name == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	if len(customer.Name) < 2 || len(customer.Name) > 40 {
		return fmt.Errorf("nama harus terdiri dari 2 hingga 40 karakter")
	}

	if len(customer.No_Telp) < 10 || len(customer.No_Telp) > 12 {
		return fmt.Errorf("nomor telepon harus terdiri dari 10 hingga 12 angka")
	}

	queryInsert := "INSERT INTO mst_customers (id_customer, name, no_telp, alamat) VALUES ($1, $2, $3, $4)"
	_, err := tx.Exec(queryInsert, customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully inserted data")
	}

	return nil
}

func UpdateCustomer(customer model.Customer, tx *sql.Tx) error {
	if customer.Name == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	if len(customer.Name) < 2 || len(customer.Name) > 40 {
		return fmt.Errorf("nama harus terdiri dari 2 hingga 40 karakter")
	}

	updateCustomer := "UPDATE mst_customers SET name = $2, no_telp = $3, alamat = $4 WHERE id_customer = $1;"
	var err error
	_, err = tx.Exec(updateCustomer, customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successfully Update Data")
	}

	return nil
}

func DeleteCustomer(id string, tx *sql.Tx) error {
	exists, err := IsCustomerExists(id, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("pelanggan dengan ID %s tidak ditemukan", id)
	}

	hasTransactions, err := HasCustomerTransactions(id, tx)
	if err != nil {
		return err
	}
	if hasTransactions {
		return fmt.Errorf("pelanggan dengan ID %s memiliki transaksi terkait dan tidak dapat dihapus", id)
	}

	query := "DELETE FROM mst_customers WHERE id_customer = $1;"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func IsCustomerExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM mst_customers WHERE id_customer = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func HasCustomerTransactions(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM trx_order WHERE customer_id = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// func Validate(err error, message string, tx *sql.Tx) {
// 	if err != nil {
// 		tx.Rollback()
// 		fmt.Println(err, "Transcantion Roolback!")
// 	} else {
// 		fmt.Println("Successfully " + message + " data!")
// 	}
// }
