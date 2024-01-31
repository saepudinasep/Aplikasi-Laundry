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
)

func MenuMasterCustomer(customer model.Customer) {
	db := config.ConnectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

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

			err = controller.AddCustomer(customer, tx)
			if err != nil {
				fmt.Println("Gagal menyimpan customer:", err)
				continue
			}

			fmt.Println("Customer berhasil disimpan.")
			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		case 2:
			// fmt.Println("Anda memilih 2")
			fmt.Print("Masukkan ID Customer yang akan diupdate: ")
			idCustomerUpdateStr, _ := reader.ReadString('\n')
			// Mengonversi ID Customer yang diambil dari string ke int
			// idCustomerUpdate, err := strings.TrimSpace(idCustomerUpdateStr)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// 	return
			// }

			exists, err := controller.IsCustomerExists(strings.TrimSpace(idCustomerUpdateStr), tx)
			if err != nil {
				return
			}
			if !exists {
				fmt.Println("Data tidak ditemukan untuk ID Customer:", idCustomerUpdateStr)
				// return
			} else {
				// fmt.Println("Data tidak ditemukan")
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

				err = controller.UpdateCustomer(customerUpdate, tx)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Customer berhasil diupdate")
				}

				err = tx.Commit()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
		case 3:
			fmt.Print("Masukkan ID Customer yang akan dihapus: ")
			idCustomerDelete, _ := reader.ReadString('\n')

			err = controller.DeleteCustomer(strings.TrimSpace(idCustomerDelete), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Customer berhasil dihapus")
			}

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		case 4:
			fmt.Println("Anda memilih 4")
		case 5:
			fmt.Println("Anda memilih 5")
		case 0:
			fmt.Println("Anda memilih 0")
			return
		default:
			fmt.Println("Pilihan tidak tepat")
		}
	}
}
