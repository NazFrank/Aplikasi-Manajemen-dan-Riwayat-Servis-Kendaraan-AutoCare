/*
	Module ini berisi fungsi-fungsi CRUD untuk mengelola riwayat servis kendaraan.
	Mencakup pencarian index servis, pengambilan plat nomor kendaraan, tampilan riwayat
	servis secara keseluruhan maupun per kendaraan, penambahan, pengubahan, dan penghapusan
	data servis. Dilengkapi perhitungan total biaya servis per kendaraan.
*/
package src

import "fmt"

// Statistik jumlah servis per bulan (format YYYY-MM)
func statistikServisPerBulan() {
	var bulan [MAX]string
	var jumlah [MAX]int
	var n, i, j int
	var bln, tempStr string
	var tempInt int
	var ditemukan bool

	fmt.Println("")
	fmt.Println("--- Statistik Jumlah Servis Per Bulan ---")
	if nServis == 0 {
		fmt.Println("Belum ada data servis.")
		return
	}

	n = 0
	for i = 0; i < nServis; i++ {
		if len(dataServis[i].Tanggal) >= 7 {
			bln = dataServis[i].Tanggal[0:7]
			ditemukan = false
			for j = 0; j < n; j++ {
				if bulan[j] == bln {
					jumlah[j] = jumlah[j] + 1
					ditemukan = true
					break
				}
			}
			if !ditemukan {
				bulan[n] = bln
				jumlah[n] = 1
				n++
			}
		}
	}

	// urutkan berdasarkan bulan secara ascending (bubble sort sederhana)
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if bulan[j] > bulan[j+1] {
				tempStr = bulan[j]
				bulan[j] = bulan[j+1]
				bulan[j+1] = tempStr
				tempInt = jumlah[j]
				jumlah[j] = jumlah[j+1]
				jumlah[j+1] = tempInt
			}
		}
	}

	for i = 0; i < n; i++ {
		fmt.Printf("Bulan %s : %d servis\n", bulan[i], jumlah[i])
	}
}

// Statistik kategori kerusakan yang paling sering muncul
func statistikKategoriKerusakan() {
	var kategori [MAX]string
	var jumlah [MAX]int
	var n, i, j, maxJumlah int
	var k, tempStr string
	var tempInt int
	var ditemukan bool

	fmt.Println("")
	fmt.Println("--- Statistik Kategori Kerusakan ---")
	if nServis == 0 {
		fmt.Println("Belum ada data servis.")
		return
	}

	n = 0
	for i = 0; i < nServis; i++ {
		k = dataServis[i].JenisKerusakan
		ditemukan = false
		for j = 0; j < n; j++ {
			if kategori[j] == k {
				jumlah[j] = jumlah[j] + 1
				ditemukan = true
				break
			}
		}
		if !ditemukan {
			kategori[n] = k
			jumlah[n] = 1
			n++
		}
	}

	// urutkan dari jumlah terbanyak ke paling sedikit (bubble sort)
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if jumlah[j] < jumlah[j+1] {
				tempInt = jumlah[j]
				jumlah[j] = jumlah[j+1]
				jumlah[j+1] = tempInt
				tempStr = kategori[j]
				kategori[j] = kategori[j+1]
				kategori[j+1] = tempStr
			}
		}
	}

	for i = 0; i < n; i++ {
		fmt.Printf("%-25s : %d kali\n", kategori[i], jumlah[i])
	}

	maxJumlah = jumlah[0]
	fmt.Println("")
	fmt.Println("Kategori kerusakan paling sering:")
	for i = 0; i < n; i++ {
		if jumlah[i] == maxJumlah {
			fmt.Printf("  - %s (%d kali)\n", kategori[i], jumlah[i])
		}
	}
}

// Menu utama statistik
func MenuStatistik() {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("--- MENU STATISTIK ---")
		fmt.Println("1. Jumlah Kendaraan Diservis per Bulan")
		fmt.Println("2. Kategori Kerusakan Paling Sering")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			statistikServisPerBulan()
		} else if pilihan == 2 {
			statistikKategoriKerusakan()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
