package routes

import (
	"bufio"
	"challenge-godb/controller"
	"challenge-godb/model"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func MenuMasterLayanan(layanan model.Layanan) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Menampilkan pilihan menu
		fmt.Println("Menu:")
		fmt.Println("1. Add Layanan")
		fmt.Println("2. Update Layanan")
		fmt.Println("3. Delete Layanan")
		fmt.Println("4. View All Layanan")
		fmt.Println("5. View Layanan by ID")
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
			fmt.Print("Masukkan ID Layanan: ")
			idLayanan, _ := reader.ReadString('\n')

			layanan = controller.GetLayananById(strings.TrimSpace(idLayanan))
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			if layanan.Id_Layanan != "" {
				fmt.Println("Layanan dengan ID tersebut sudah digunakan.")
			} else {
				fmt.Print("Masukkan Nama Layanan: ")
				namaLayanan, _ := reader.ReadString('\n')

				fmt.Print("Masukkan Harga Layanan: ")
				hargaLayananStr, _ := reader.ReadString('\n')
				hargaLayanan, _ := strconv.Atoi(strings.TrimSpace(hargaLayananStr))

				fmt.Print("Masukkan Satuan Layanan: ")
				satuanLayanan, _ := reader.ReadString('\n')

				layanan := model.Layanan{
					Id_Layanan:   strings.TrimSpace(idLayanan),
					Nama_Layanan: strings.TrimSpace(namaLayanan),
					Harga:        hargaLayanan,
					Satuan:       strings.TrimSpace(satuanLayanan),
				}
				err = controller.AddLayanan(layanan)
				if err != nil {
					fmt.Println("Gagal menyimpan layanan:", err)
					continue
				}
			}
		case 2:
			fmt.Print("Masukkan ID Layanan yang akan diupdate: ")
			idLayananUpdate, _ := reader.ReadString('\n')

			layanan = controller.GetLayananById(strings.TrimSpace(idLayananUpdate))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if layanan.Id_Layanan == "" {
				fmt.Println("Layanan dengan ID tertentu tidak ditemukan.")
			} else {
				fmt.Print("Masukkan Nama Layanan baru: ")
				namaLayananUpdate, _ := reader.ReadString('\n')

				fmt.Print("Masukkan Harga Layanan baru: ")
				hargaLayananStr, _ := reader.ReadString('\n')
				hargaLayanan, _ := strconv.Atoi(strings.TrimSpace(hargaLayananStr))

				fmt.Print("Masukkan Satuan Layanan baru: ")
				satuanLayananUpdate, _ := reader.ReadString('\n')

				layananUpdate := model.Layanan{
					Id_Layanan:   strings.TrimSpace(idLayananUpdate),
					Nama_Layanan: strings.TrimSpace(namaLayananUpdate),
					Harga:        hargaLayanan,
					Satuan:       strings.TrimSpace(satuanLayananUpdate),
				}

				err = controller.UpdateLayanan(layananUpdate)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case 3:
			fmt.Print("Masukkan ID Layanan yang akan dihapus: ")
			idLayananDelete, _ := reader.ReadString('\n')

			layanan = controller.GetLayananById(strings.TrimSpace(idLayananDelete))

			if layanan.Id_Layanan == "" {
				fmt.Println("Layanan dengan ID tertentu tidak ditemukan.")
			} else {
				err = controller.DeleteLayanan(strings.TrimSpace(idLayananDelete))
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case 4:
			layanans := controller.GetAllLayanan()

			// Membuat tabel baru
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID Layanan", "Nama Layanan", "Harga", "Satuan"})

			// Menambahkan data ke dalam tabel
			for _, layanan := range layanans {
				row := []string{
					(layanan.Id_Layanan),
					layanan.Nama_Layanan,
					fmt.Sprintf("%d", layanan.Harga),
					layanan.Satuan,
				}
				table.Append(row)
			}

			// Menampilkan tabel
			table.Render()
		case 5:
			fmt.Print("Masukan ID Layanan yang akan ditampilkan: ")
			idLayananById, _ := reader.ReadString('\n')

			layanan = controller.GetLayananById(strings.TrimSpace(idLayananById))
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			if layanan.Id_Layanan == "" {
				fmt.Println("Layanan dengan Id tertentu tidak ditemukan.")
			} else {
				fmt.Printf("ID Layanan: %s\n", layanan.Id_Layanan)
				fmt.Printf("Nama Layanan: %s\n", layanan.Nama_Layanan)
				fmt.Printf("Harga: %d\n", layanan.Harga)
				fmt.Printf("Satuan: %s\n", layanan.Satuan)
			}
		case 0:
			fmt.Println("Keluar dari menu Master Layanan")
			return
		default:
			fmt.Println("Menu tidak valid")
		}

	}
}
