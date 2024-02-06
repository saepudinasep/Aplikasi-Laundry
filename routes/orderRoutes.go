package routes

import (
	"bufio"
	"challenge-godb/config"
	"challenge-godb/controller"
	"challenge-godb/model"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func TransaksiMasterOrder(order model.Order) {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Menampilkan pilihan menu
		fmt.Println("Menu:")
		fmt.Println("1. Create Orders")
		fmt.Println("2. Take Orders")
		fmt.Println("3. View Orders")
		fmt.Println("4. View Orders by Customers Id")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")

		// Membaca input pilihan menu dari pengguna
		menuStr, _ := reader.ReadString('\n')
		menu, err := strconv.Atoi(strings.TrimSpace(menuStr))
		if err != nil {
			fmt.Println("Input tidak valid")
			continue
		}

		switch menu {
		case 1:
			addOrders(reader)
		case 2:
			db := config.ConnectDb()
			// var customerId string
			fmt.Print("Masukkan ID Customer: ")
			customerId, _ := reader.ReadString('\n')
			// Start Order
			orders := controller.GetOrderByCustomerId(strings.TrimSpace(customerId))
			for _, order := range orders {
				tanggal := controller.FormatDatabaseDate(order.Tanggal_Masuk)
				fmt.Println(strings.Repeat("=", 50))
				fmt.Println("No: ", order.Id_Order)
				fmt.Println("Tanggal Masuk: ", tanggal)

				// Start Customer
				fmt.Println(strings.Repeat("=", 50))
				customer := controller.GetCustomerById(strings.TrimSpace(order.Customer_Id))
				fmt.Println("Nama Customer: ", customer.Name)
				fmt.Println("No HP: ", customer.No_Telp)
				// End Customer

				// Start Order Detail
				orderDetails := controller.GetOrderDetailByOrderId(order.Id_Order)
				for _, orderDetail := range orderDetails {

					// Start Layanan
					// layanan = controller.GetLayananById(strings.TrimSpace(orderDetail.Layanan_Id))
					// fmt.Println(strings.Repeat("=", 50))
					// fmt.Println()
					// End Layanan

					fmt.Println(orderDetail.Id_Order_Detail, orderDetail.Order_Id, orderDetail.Layanan_Id, orderDetail.Quantity)

					orderIDInt, err := strconv.Atoi(order.Id_Order)
					if err != nil {
						panic(err)
					}

					totalHarga, err := controller.GetTotalHarga(db, orderIDInt)
					if err != nil {
						panic(err)
					}

					fmt.Printf("Total Harga: %d\n", totalHarga)

					fmt.Print("Masukkan Penerima: ")
					namaPenerimaUpdate, _ := reader.ReadString('\n')

					orderUpdate := model.Order{
						Id_Order:       order.Id_Order,
						Tanggal_Keluar: time.Now().Local(),
						Penerima:       strings.TrimSpace(namaPenerimaUpdate),
					}

					err = controller.UpdateOrder(orderUpdate)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Terimakasih Sudah Transaksi!")
					}
				} // End Order Detail
			} // End Order

			// fmt.Println("Customer  : ", customerId)
			// fmt.Print("Masukkan Nama Penerima: ")
			// fmt.Scan(&penerima)
			// fmt.Println("Penerima  : ", penerima)
		case 3:
		case 4:
		case 0:
			fmt.Println("Keluar dari menu Order")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}

func addOrders(reader *bufio.Reader) {
	db := config.ConnectDb()
	defer db.Close()
	// Mulai transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Input ID customer dari pengguna
	var customerID string
	fmt.Print("Masukkan ID Customer: ")
	fmt.Scan(&customerID)

	customer := controller.GetCustomerById(customerID)

	if customer.Id_Customer == "" {
		fmt.Println("Id Customer belum terdaftar")
	} else {
		fmt.Printf("ID Customer: %s\n", customer.Id_Customer)
		fmt.Printf("Nama: %s\n", customer.Name)
		fmt.Printf("No Telp: %s\n", customer.No_Telp)
		fmt.Printf("Alamat: %s\n", customer.Alamat)
		var Check string
		fmt.Print("Apakah Benar itu Data Anda? (y/t): ")
		fmt.Scan(&Check)
		reader.ReadString('\n')
		if Check == "y" {

			// Input data ke tabel trx_order
			order := model.Order{Customer_Id: customerID}
			err = controller.CreateOrder(tx, &order)
			if err != nil {
				tx.Rollback()
				panic(err)
			}

			// Loop untuk menambahkan layanan tambahan
			for {
				// Input ID layanan dan jumlah dari pengguna
				var layananID string
				var quantity int
				fmt.Print("Masukkan Layanan ID: ")
				fmt.Scan(&layananID)

				layanan := controller.GetLayananById(layananID)

				if layanan.Id_Layanan == "" {
					fmt.Println("Id dengan layanan tersebut belum tersedia.")
				} else {

					fmt.Printf("ID Layanan: %s\n", layanan.Id_Layanan)
					fmt.Printf("Nama Layanan: %s\n", layanan.Nama_Layanan)
					fmt.Printf("Harga: %d\n", layanan.Harga)
					fmt.Printf("Satuan: %s\n", layanan.Satuan)

					fmt.Print("Masukkan Jumlah: ")
					fmt.Scan(&quantity)

					// Input data ke tabel trx_order_detail
					orderDetail := model.OrderDetail{Order_Id: order.Id_Order, Layanan_Id: layananID, Quantity: quantity}
					_, err = controller.CreateOrderDetail(tx, &orderDetail)
					if err != nil {
						tx.Rollback()
						panic(err)
					}

					// Tanyakan apakah ada layanan tambahan
					var tambahanLayanan string
					fmt.Print("Apakah ada layanan tambahan? (y/t): ")
					fmt.Scan(&tambahanLayanan)
					reader.ReadString('\n')

					// Jika tidak ada layanan tambahan, keluar dari loop
					if tambahanLayanan != "y" {
						break
					}
				}

			}

			// Commit transaksi
			err = tx.Commit()
			if err != nil {
				panic(err)
			}

			fmt.Println("Transaction successfully!")

			// Konversi order.ID dari string ke int
			orderIDInt, err := strconv.Atoi(order.Id_Order)
			if err != nil {
				panic(err)
			}

			// Tampilkan total harga
			totalHarga, err := controller.GetTotalHarga(db, orderIDInt)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Total Harga: %d\n", totalHarga)
			fmt.Println("Id Order Anda: ", order.Id_Order)
		}
	}

}
