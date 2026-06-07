/*
	Module ini berisi fungsi-fungsi pencarian kendaraan berdasarkan plat nomor.
	Menyediakan dua algoritma pencarian: Sequential Search (linear, tanpa prasyarat urutan)
	dan Binary Search (membutuhkan data terurut, lebih efisien untuk dataset besar).
	Dilengkapi fungsi pengurutan data sebelum Binary Search dan tampilan detail kendaraan
	beserta riwayat servisnya.
*/
package src

import "fmt"

// Sequential Search berdasarkan plat nomor.
// Mengembalikan index data kendaraan, atau -1 jika tidak ditemukan.
func sequentialSearchPlat(plat string) int {
	var i int
	for i = 0; i < nKendaraan; i++ {
		if dataKendaraan[i].PlatNomor == plat {
			return i
		}
	}
	return -1
}

// Mengurutkan dataKendaraan berdasarkan PlatNomor (selection sort)
// agar dapat digunakan untuk Binary Search.
func urutkanKendaraanByPlat() {
	var i, j, idxMin int
	var temp Kendaraan
	for i = 0; i < nKendaraan-1; i++ {
		idxMin = i
		for j = i + 1; j < nKendaraan; j++ {
			if dataKendaraan[j].PlatNomor < dataKendaraan[idxMin].PlatNomor {
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

// Binary Search berdasarkan plat nomor.
// PRASYARAT: data sudah diurutkan berdasarkan PlatNomor.
// Mengembalikan index data kendaraan, atau -1 jika tidak ditemukan.
func binarySearchPlat(plat string) int {
	var low, high, mid int
	low = 0
	high = nKendaraan - 1
	for low <= high {
		mid = (low + high) / 2
		if dataKendaraan[mid].PlatNomor == plat {
			return mid
		} else if dataKendaraan[mid].PlatNomor < plat {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Menampilkan detail kendaraan termasuk riwayat servisnya
func tampilkanDetailKendaraan(idx int) {
	var i, banyakServis int
	var k Kendaraan
	k = dataKendaraan[idx]

	fmt.Println("")
	fmt.Println("--- Detail Kendaraan Ditemukan ---")
	fmt.Printf("ID: %d | Plat: %s | %s %s | Tahun: %d | Pemilik: %s\n",
		k.ID, k.PlatNomor, k.Merk, k.Model, k.TahunProduksi, namaPemilik(k.IDPemilik))

	fmt.Println("Riwayat Servis:")
	banyakServis = 0
	for i = 0; i < nServis; i++ {
		if dataServis[i].IDKendaraan == k.ID {
			fmt.Printf("  - Tgl: %s | Kerusakan: %s | Biaya: Rp %.0f\n",
				dataServis[i].Tanggal, dataServis[i].JenisKerusakan, dataServis[i].Biaya)
			banyakServis++
		}
	}
	if banyakServis == 0 {
		fmt.Println("  (Belum ada riwayat servis)")
	}
}

// Menu utama pencarian kendaraan
func MenuPencarian() {
	var pilihan, idx int
	var plat string

	for {
		fmt.Println("")
		fmt.Println("--- MENU PENCARIAN KENDARAAN ---")
		fmt.Println("1. Sequential Search (berdasarkan Plat Nomor)")
		fmt.Println("2. Binary Search (berdasarkan Plat Nomor)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			if nKendaraan == 0 {
				fmt.Println("Belum ada data kendaraan.")
			} else {
				fmt.Print("Masukkan plat nomor yang dicari: ")
				fmt.Scan(&plat)
				idx = sequentialSearchPlat(plat)
				if idx == -1 {
					fmt.Println("Kendaraan tidak ditemukan.")
				} else {
					tampilkanDetailKendaraan(idx)
				}
			}
		} else if pilihan == 2 {
			if nKendaraan == 0 {
				fmt.Println("Belum ada data kendaraan.")
			} else {
				fmt.Println("Data akan diurutkan dulu berdasarkan plat nomor untuk Binary Search...")
				urutkanKendaraanByPlat()
				fmt.Println("Data setelah diurutkan:")
				tampilkanSemuaKendaraan()
				fmt.Print("Masukkan plat nomor yang dicari: ")
				fmt.Scan(&plat)
				idx = binarySearchPlat(plat)
				if idx == -1 {
					fmt.Println("Kendaraan tidak ditemukan.")
				} else {
					tampilkanDetailKendaraan(idx)
				}
			}
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
