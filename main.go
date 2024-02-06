package main

import (
	"challenge-godb/model"
	"challenge-godb/routes"
	"fmt"
)

func main() {
	fmt.Println("========== Selamat datang di Aplikasi Enigma Laundry ==========")
	fmt.Println("==================== *Silakan pilih menu* ====================")
	fmt.Println("1. Master Customer")
	fmt.Println("2. Master Layanan")
	fmt.Println("3. Make Orders")
	fmt.Println("0. Keluar")

	var menu int
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		routes.MenuMasterCustomer(model.Customer{})
	case 2:
		routes.MenuMasterLayanan(model.Layanan{})
	case 3:
		routes.TransaksiMasterOrder(model.Order{})
	case 0:
		fmt.Println("Terima kasih telah menggunakan Aplikasi Enigma Laundry. Sampai jumpa!")
	default:
		fmt.Println("Menu tidak valid")
	}
}
