package routes

import (
	"bufio"
	"challenge-godb/model"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		case 2:

		case 3:

		case 4:

		case 5:

		case 0:

		default:

		}

	}
}
