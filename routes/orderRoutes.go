package routes

import (
	"bufio"
	"challenge-godb/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TransaksiMasterOrder(order model.Order) {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Menampilkan pilihan menu
		fmt.Println("Menu:")
		fmt.Println("1. Create Order")
		fmt.Println("2. View All Orders")
		fmt.Println("3. View Order by ID")
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
		case 2:
		case 3:
		case 0:
			fmt.Println("Keluar dari menu Order")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
