package main

import (
	"autocare/src"
	"fmt"
)

func main() {
	var pilihan int

	src.IsiDataAwal()

	for {
		fmt.Println("")
		fmt.Println("==========================================")
		fmt.Println("      APLIKASI AUTOCARE (Bengkel)         ")
		fmt.Println("==========================================")
		fmt.Println("CATATAN: gunakan underscore (_) sebagai pengganti spasi pada input.")
		fmt.Println("Contoh: Budi_Santoso, B1234ABC, Ganti_Oli")
		fmt.Println("------------------------------------------")
		fmt.Println("1. Kelola Data Pemilik")
		fmt.Println("2. Kelola Data Kendaraan")
		fmt.Println("3. Kelola Riwayat Servis")
		fmt.Println("4. Cari Kendaraan (berdasarkan Plat Nomor)")
		fmt.Println("5. Urutkan Data Kendaraan")
		fmt.Println("6. Statistik")
		fmt.Println("0. Keluar")
		fmt.Println("------------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			src.MenuPemilik()
		} else if pilihan == 2 {
			src.MenuKendaraan()
		} else if pilihan == 3 {
			src.MenuServis()
		} else if pilihan == 4 {
			src.MenuPencarian()
		} else if pilihan == 5 {
			src.MenuPengurutan()
		} else if pilihan == 6 {
			src.MenuStatistik()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan AutoCare. Sampai jumpa!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia. Silakan coba lagi.")
		}
	}
}
