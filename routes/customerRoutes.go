package routes

import (
	"bufio"
	"challenge-godb/controller"
	"challenge-godb/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuMasterCustomer(customer model.Customer) {

	reader := bufio.NewReader(os.Stdin)

	for {
		// Menampilkan pilihan menu
		fmt.Println("Menu:")
		fmt.Println("1. Add Customer")
		fmt.Println("2. Update Customer")
		fmt.Println("3. Delete Customer")
		fmt.Println("4. View All Customers")
		fmt.Println("5. View Customer by ID")
		fmt.Println("0. Keluar")
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
			fmt.Print("Masukkan ID Customer: ")
			idCustomer, _ := reader.ReadString('\n')

			customer = controller.GetCustomerById(strings.TrimSpace(idCustomer))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if customer.Id_Customer != "" {
				fmt.Println("Customer dengan ID tersebut sudah digunakan.")
			} else {
				fmt.Print("Masukkan Nama Customer: ")
				namaCustomer, _ := reader.ReadString('\n')

				fmt.Print("Masukkan Nomor Telepon Customer: ")
				teleponCustomer, _ := reader.ReadString('\n')

				fmt.Print("Masukkan Alamat Customer: ")
				alamatCustomer, _ := reader.ReadString('\n')

				customer := model.Customer{
					Id_Customer: strings.TrimSpace(idCustomer),
					Name:        strings.TrimSpace(namaCustomer),
					No_Telp:     strings.TrimSpace(teleponCustomer),
					Alamat:      strings.TrimSpace(alamatCustomer),
				}
				err = controller.AddCustomer(customer)
				if err != nil {
					fmt.Println("Gagal menyimpan customer:", err)
					continue
				}
			}

		case 2:
			fmt.Print("Masukkan ID Customer yang akan diupdate: ")
			idCustomerUpdateStr, _ := reader.ReadString('\n')

			customer = controller.GetCustomerById(strings.TrimSpace(idCustomerUpdateStr))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if customer.Id_Customer == "" {
				fmt.Println("Customer dengan ID tertentu tidak ditemukan.")
			} else {
				fmt.Print("Masukkan Nama Customer baru: ")
				namaCustomerUpdate, _ := reader.ReadString('\n')

				fmt.Print("Masukkan Nomor Telepon Customer baru: ")
				teleponCustomerUpdate, _ := reader.ReadString('\n')

				fmt.Print("Masukkan Alamat Customer baru: ")
				alamatCustomerUpdate, _ := reader.ReadString('\n')

				customerUpdate := model.Customer{
					Id_Customer: strings.TrimSpace(idCustomerUpdateStr),
					Name:        strings.TrimSpace(namaCustomerUpdate),
					No_Telp:     strings.TrimSpace(teleponCustomerUpdate),
					Alamat:      strings.TrimSpace(alamatCustomerUpdate),
				}

				err = controller.UpdateCustomer(customerUpdate)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case 3:
			fmt.Print("Masukkan ID Customer yang akan dihapus: ")
			idCustomerDelete, _ := reader.ReadString('\n')

			customer = controller.GetCustomerById(strings.TrimSpace(idCustomerDelete))

			if customer.Id_Customer == "" {
				fmt.Println("Customer dengan ID tertentu tidak ditemukan.")
			} else {
				err = controller.DeleteCustomer(strings.TrimSpace(idCustomerDelete))
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case 4:
			customers := controller.GetAllCustomer()
			for _, customer := range customers {
				fmt.Println(customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
			}
		case 5:
			fmt.Print("Masukkan ID Customer yang akan ditampilkan: ")
			idCustomerByID, _ := reader.ReadString('\n')

			customer := controller.GetCustomerById(strings.TrimSpace(idCustomerByID))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if customer.Id_Customer == "" {
				fmt.Println("customer dengan Id tertentu tidak ditemukan.")
			} else {
				fmt.Printf("ID Customer: %s\n", customer.Id_Customer)
				fmt.Printf("Nama: %s\n", customer.Name)
				fmt.Printf("No Telp: %s\n", customer.No_Telp)
				fmt.Printf("Alamat: %s\n", customer.Alamat)
			}
		case 0:
			fmt.Println("Keluar dari menu Master Customer")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
