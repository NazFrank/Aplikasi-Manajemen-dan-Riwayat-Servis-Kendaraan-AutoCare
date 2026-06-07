/*
	Module ini berisi fungsi-fungsi pengurutan data kendaraan menggunakan dua algoritma:
	Selection Sort untuk mengurutkan kendaraan berdasarkan tahun produksi secara ascending,
	dan Insertion Sort untuk mengurutkan berdasarkan tanggal servis terakhir secara descending.
	Terdapat pula fungsi pembantu untuk mengambil tanggal servis terakhir tiap kendaraan
	dan menampilkan hasil pengurutan beserta informasi servis.
*/
package src

import "fmt"

// Mengambil tanggal servis terakhir untuk satu kendaraan.
// Return "" jika belum pernah diservis.
func tanggalServisTerakhir(idKendaraan int) string {
	var i int
	var tgl string
	tgl = ""
	for i = 0; i < nServis; i++ {
		if dataServis[i].IDKendaraan == idKendaraan {
			if dataServis[i].Tanggal > tgl {
				tgl = dataServis[i].Tanggal
			}
		}
	}
	return tgl
}

// Selection Sort: urutkan dataKendaraan berdasarkan TahunProduksi (ascending)
func selectionSortTahun() {
	var i, j, idxMin int
	var temp Kendaraan
	for i = 0; i < nKendaraan-1; i++ {
		idxMin = i
		for j = i + 1; j < nKendaraan; j++ {
			if dataKendaraan[j].TahunProduksi < dataKendaraan[idxMin].TahunProduksi {
				idxMin = j
			}
		}
		if idxMin != i {
			temp = dataKendaraan[i]
			dataKendaraan[i] = dataKendaraan[idxMin]
			dataKendaraan[idxMin] = temp
		}
	}
}

// Insertion Sort: urutkan dataKendaraan berdasarkan tanggal servis terakhir (descending - paling baru di atas)
func insertionSortTanggalServis() {
	var i, j int
	var key Kendaraan
	var keyTgl, banding string
	for i = 1; i < nKendaraan; i++ {
		key = dataKendaraan[i]
		keyTgl = tanggalServisTerakhir(key.ID)
		j = i - 1
		for j >= 0 {
			banding = tanggalServisTerakhir(dataKendaraan[j].ID)
			if banding < keyTgl {
				dataKendaraan[j+1] = dataKendaraan[j]
				j--
			} else {
				break
			}
		}
		dataKendaraan[j+1] = key
	}
}

// Menampilkan kendaraan beserta tanggal servis terakhirnya
func tampilkanKendaraanDenganTanggalServis() {
	var i int
	var tgl string
	fmt.Println("--- Daftar Kendaraan (dengan Tanggal Servis Terakhir) ---")
	if nKendaraan == 0 {
		fmt.Println("Belum ada data kendaraan.")
		return
	}
	for i = 0; i < nKendaraan; i++ {
		tgl = tanggalServisTerakhir(dataKendaraan[i].ID)
		if tgl == "" {
			tgl = "Belum_pernah_servis"
		}
		fmt.Printf("ID: %-3d | Plat: %-12s | %-10s %-10s | Tahun: %d | Servis Terakhir: %s\n",
			dataKendaraan[i].ID, dataKendaraan[i].PlatNomor, dataKendaraan[i].Merk,
			dataKendaraan[i].Model, dataKendaraan[i].TahunProduksi, tgl)
	}
}

// Menu utama pengurutan kendaraan
func MenuPengurutan() {
	var pilihan int

	for {
		fmt.Println("")
		fmt.Println("--- MENU PENGURUTAN KENDARAAN ---")
		fmt.Println("1. Selection Sort - berdasarkan Tahun Produksi (lama ke baru)")
		fmt.Println("2. Insertion Sort - berdasarkan Tanggal Servis Terakhir (baru ke lama)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			if nKendaraan == 0 {
				fmt.Println("Belum ada data kendaraan.")
			} else {
				selectionSortTahun()
				fmt.Println("")
				fmt.Println("Data berhasil diurutkan berdasarkan Tahun Produksi:")
				tampilkanSemuaKendaraan()
			}
		} else if pilihan == 2 {
			if nKendaraan == 0 {
				fmt.Println("Belum ada data kendaraan.")
			} else {
				insertionSortTanggalServis()
				fmt.Println("")
				fmt.Println("Data berhasil diurutkan berdasarkan Tanggal Servis Terakhir:")
				tampilkanKendaraanDenganTanggalServis()
			}
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
