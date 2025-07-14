package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"uts_strukturdata/view"
)

var currentRole string

func login(reader *bufio.Reader) bool {
	fmt.Print("Login sebagai (admin/pengelolalagu): ")
	inputRole, _ := reader.ReadString('\n')
	role := strings.ToLower(strings.TrimSpace(inputRole))
	if role == "admin" || role == "pengelolalagu" {
		currentRole = role
		return true
	}
	fmt.Println("Role tidak valid.")
	return false
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	if !login(reader) {
		return
	}

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Management Lagu")
		fmt.Println("2. Tampilkan Lagu Tidak Berlisensi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		pilihan, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Input harus berupa angka.")
			continue
		}
		switch pilihan {
		case 1:
			view.ManagementLagu(reader, currentRole)
		case 2:
			view.ViewLaguTidakBerlisensi()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
