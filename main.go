package main

import (
	"fmt"
	"autocare/data"
	"autocare/menu"
)

func main() {
	var pilihan int

	data.InisialisasiDataContoh()

	fmt.Println("================================================")
	fmt.Println("      SELAMAT DATANG DI APLIKASI AUTOCARE       ")
	fmt.Println("   Sistem Manajemen & Riwayat Servis Kendaraan  ")
	fmt.Println("================================================")

	for pilihan != 7 {
		menu.TampilkanMenuUtama()
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			menu.MenuKendaraan()
		case 2:
			menu.MenuPemilik()
		case 3:
			menu.MenuServis()
		case 4:
			menu.MenuCariKendaraan()
		case 5:
			menu.MenuUrutkanKendaraan()
		case 6:
			menu.MenuStatistik()
		case 7:
			fmt.Println("Terima kasih telah menggunakan AutoCare. Sampai jumpa!")
		default:
			fmt.Println("[!] Pilihan tidak valid. Silakan coba lagi.")
		}

		fmt.Println()
	}
}
