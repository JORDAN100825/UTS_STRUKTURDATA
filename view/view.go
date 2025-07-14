package view

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"uts_strukturdata/model"
	"uts_strukturdata/node"
)

func ManagementLagu(reader *bufio.Reader, role string) {
	fmt.Println("== Management Lagu ==")
	for {
		fmt.Println("\nMenu:")
		if role == "admin" {
			fmt.Println("1. Tambah Lagu")
		}
		if role == "pengelolalagu" || role == "admin" {
			fmt.Println("2. Lihat Daftar Lagu Berlisensi")
		}
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")
		choice := readLine(reader)

		switch choice {
		case "0":
			return
		case "1":
			if role == "admin" {
				InsertLagu(reader)
			} else {
				fmt.Println("Hanya admin yang boleh menambah lagu.")
			}
		case "2":
			LihatLaguBerlisensi()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func LihatLaguBerlisensi() {
	fmt.Println("=== Daftar Lagu Berlisensi ===")
	temp := model.LaguBerlisensiHead
	if temp == nil {
		fmt.Println("Belum ada lagu berlisensi.")
		return
	}
	for temp != nil {
		lagu := temp.Data
		fmt.Printf("ID: %d | Judul: %s | Penyanyi: %s | Tahun: %d | Durasi: %s\n",
			lagu.ID, lagu.Judul, lagu.Penyanyi, lagu.Tahun, lagu.Durasi)
		temp = temp.Next
	}
}



func InsertLagu(reader *bufio.Reader) {
	fmt.Print("Masukkan ID Lagu: ")
	id, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Masukkan Judul Lagu: ")
	judul := readLine(reader)
	fmt.Print("Masukkan Nama Penyanyi: ")
	penyanyi := readLine(reader)
	fmt.Print("Masukkan Tahun Rilis: ")
	tahun, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Durasi: ")
	durasi := readLine(reader)

	lagu := node.Lagu{
		ID:      id,
		Judul:   judul,
		Penyanyi: penyanyi,
		Tahun:   tahun,
		Durasi:   durasi,
	}

	fmt.Print("Apakah lagu ini berlisensi? (ya/tidak): ")
	isLicensed := strings.ToLower(readLine(reader))
		if isLicensed == "ya" {
		model.TambahLaguBerlisensi(lagu)
		fmt.Println("Lagu Berlisensi berhasil ditambahkan.")
} else {
		fmt.Println("\n=== Daftar Lagu Berlisensi untuk Referensi ===")
		LihatLaguBerlisensi() 

			fmt.Print("Masukkan ID Referensi Lagu Berlisensi: ")
			idRef, _ := strconv.Atoi(readLine(reader))

	if model.TambahLaguTidakBerlisensi(lagu.ID, idRef) {
		fmt.Println("Lagu Tidak Berlisensi berhasil ditambahkan.")
	} else {
		fmt.Println("Gagal: ID referensi tidak ditemukan.")
	}
}

}

func ViewLaguTidakBerlisensi() {
	fmt.Println("=== Daftar Lagu Tidak Berlisensi ===")
	temp := model.GetLaguTidakBerlisensi()
	for temp != nil {
		ref := model.CariLaguBerlisensi(temp.IDReferensi)
		if ref != nil {
			fmt.Printf("ID: %d | Meniru: %s oleh %s\n", temp.IDEntry, ref.Judul, ref.Penyanyi)
		} else {
			fmt.Printf("ID: %d | Referensi tidak ditemukan\n", temp.IDEntry)
		}
		temp = temp.Next
	}
}

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
