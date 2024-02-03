package controller

import (
	"challenge-godb/config"
	"challenge-godb/model"
	"database/sql"
	"fmt"
)

func AddCustomer(customer model.Customer) error {
	db := config.ConnectDb()
	defer db.Close()

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
	_, err := db.Exec(queryInsert, customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully Inserted Data!")
	}

	return nil
}

func UpdateCustomer(customer model.Customer) error {
	db := config.ConnectDb()
	defer db.Close()
	var err error

	if customer.Name == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	if len(customer.Name) < 2 || len(customer.Name) > 40 {
		return fmt.Errorf("nama harus terdiri dari 2 hingga 40 karakter")
	}

	updateCustomer := "UPDATE mst_customers SET name = $2, no_telp = $3, alamat = $4 WHERE id_customer = $1;"
	_, err = db.Exec(updateCustomer, customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
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

	hasTransactions, err := HasCustomerTransactions(id)
	if err != nil {
		return err
	}
	if hasTransactions {
		return fmt.Errorf("pelanggan dengan ID %s memiliki transaksi terkait dan tidak dapat dihapus", id)
	}

	query := "DELETE FROM mst_customers WHERE id_customer = $1;"
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully deleted data!")
	}

	return nil
}

func IsCustomerExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM mst_customers WHERE id_customer = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		panic(err)
	}

	return customers
}

func GetCustomerById(id string) model.Customer {
	db := config.ConnectDb()
	defer db.Close()
	var err error

	sqlQuery := "SELECT * FROM mst_customers WHERE id_customer = $1"

	customer := model.Customer{}
	err = db.QueryRow(sqlQuery, id).Scan(&customer.Id_Customer, &customer.Name, &customer.No_Telp, &customer.Alamat)
	if err == sql.ErrNoRows {
		return model.Customer{}
	} else if err != nil {
		panic(err)
	}

	return customer
}

func HasCustomerTransactions(id string) (bool, error) {
	db := config.ConnectDb()
	defer db.Close()

	query := "SELECT COUNT(*) FROM trx_order WHERE customer_id = $1;"
	var count int
	err := db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
