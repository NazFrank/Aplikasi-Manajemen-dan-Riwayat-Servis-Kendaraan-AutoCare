/*
	Module ini berisi fungsi-fungsi CRUD untuk mengelola data kendaraan.
	Mencakup pencarian index kendaraan, pengambilan nama pemilik, tampilan daftar kendaraan,
	penambahan, pengubahan, dan penghapusan data kendaraan. Penghapusan kendaraan secara
	otomatis menghapus seluruh riwayat servis yang terkait (cascade delete).
*/

package src

import "fmt"

// Mencari index kendaraan berdasarkan ID. Return -1 jika tidak ditemukan
func cariIndexKendaraan(id int) int {
	var i int
	for i = 0; i < nKendaraan; i++ {
		if dataKendaraan[i].ID == id {
			return i
		}
	}
	return -1
}

// Mengambil nama pemilik berdasarkan id pemilik
func namaPemilik(idPemilik int) string {
	var idx int
	idx = cariIndexPemilik(idPemilik)
	if idx == -1 {
		return "-"
	}
	return dataPemilik[idx].Nama
}

// Menampilkan seluruh data kendaraan
func tampilkanSemuaKendaraan() {
	var i int
	fmt.Println("")
	fmt.Println("--- Daftar Kendaraan ---")
	if nKendaraan == 0 {
		fmt.Println("Belum ada data kendaraan.")
		return
	}
	for i = 0; i < nKendaraan; i++ {
		fmt.Printf("ID: %-3d | Plat: %-12s | %-10s %-10s | Tahun: %d | Pemilik: %s\n",
			dataKendaraan[i].ID, dataKendaraan[i].PlatNomor, dataKendaraan[i].Merk,
			dataKendaraan[i].Model, dataKendaraan[i].TahunProduksi, namaPemilik(dataKendaraan[i].IDPemilik))
	}
}

// Menambah data kendaraan baru
func tambahKendaraan() {
	var k Kendaraan
	var idx int

	if nKendaraan >= MAX {
		fmt.Println("Data kendaraan sudah penuh.")
		return
	}
	if nPemilik == 0 {
		fmt.Println("Belum ada pemilik. Silakan tambahkan pemilik terlebih dahulu.")
		return
	}

	fmt.Println("")
	fmt.Println("--- Tambah Kendaraan Baru ---")
	fmt.Print("Plat Nomor (tanpa spasi)   : ")
	fmt.Scan(&k.PlatNomor)
	fmt.Print("Merk                       : ")
	fmt.Scan(&k.Merk)
	fmt.Print("Model                      : ")
	fmt.Scan(&k.Model)
	fmt.Print("Tahun Produksi             : ")
	fmt.Scan(&k.TahunProduksi)

	tampilkanSemuaPemilik()
	fmt.Print("Masukkan ID Pemilik: ")
	fmt.Scan(&k.IDPemilik)
	idx = cariIndexPemilik(k.IDPemilik)
	if idx == -1 {
		fmt.Println("Pemilik tidak ditemukan. Pembatalan penambahan kendaraan.")
		return
	}

	k.ID = idKendaraanNext
	dataKendaraan[nKendaraan] = k
	nKendaraan++
	idKendaraanNext++

	fmt.Printf("Kendaraan berhasil ditambahkan dengan ID %d.\n", k.ID)
}

// Mengubah data kendaraan berdasarkan ID
func ubahKendaraan() {
	var id, idx, idPemBaru, idxPem, tahunBaru int
	var plat, merk, model string

	tampilkanSemuaKendaraan()
	if nKendaraan == 0 {
		return
	}

	fmt.Print("Masukkan ID kendaraan yang akan diubah: ")
	fmt.Scan(&id)
	idx = cariIndexKendaraan(id)
	if idx == -1 {
		fmt.Println("Kendaraan dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Ketik tanda - (minus) jika tidak ingin mengubah field string, atau 0 untuk field angka.")

	fmt.Printf("Plat Nomor (%s): ", dataKendaraan[idx].PlatNomor)
	fmt.Scan(&plat)
	if plat != "-" {
		dataKendaraan[idx].PlatNomor = plat
	}

	fmt.Printf("Merk (%s): ", dataKendaraan[idx].Merk)
	fmt.Scan(&merk)
	if merk != "-" {
		dataKendaraan[idx].Merk = merk
	}

	fmt.Printf("Model (%s): ", dataKendaraan[idx].Model)
	fmt.Scan(&model)
	if model != "-" {
		dataKendaraan[idx].Model = model
	}

	fmt.Printf("Tahun Produksi (%d) [0 untuk skip]: ", dataKendaraan[idx].TahunProduksi)
	fmt.Scan(&tahunBaru)
	if tahunBaru > 0 {
		dataKendaraan[idx].TahunProduksi = tahunBaru
	}

	fmt.Printf("ID Pemilik (%d) [0 untuk skip]: ", dataKendaraan[idx].IDPemilik)
	fmt.Scan(&idPemBaru)
	if idPemBaru > 0 {
		idxPem = cariIndexPemilik(idPemBaru)
		if idxPem == -1 {
			fmt.Println("Pemilik tidak ditemukan. ID pemilik tidak diubah.")
		} else {
			dataKendaraan[idx].IDPemilik = idPemBaru
		}
	}

	fmt.Println("Data kendaraan berhasil diubah.")
}

// Menghapus data kendaraan berdasarkan ID
func hapusKendaraan() {
	var id, idx, i, j int

	tampilkanSemuaKendaraan()
	if nKendaraan == 0 {
		return
	}

	fmt.Print("Masukkan ID kendaraan yang akan dihapus: ")
	fmt.Scan(&id)
	idx = cariIndexKendaraan(id)
	if idx == -1 {
		fmt.Println("Kendaraan dengan ID tersebut tidak ditemukan.")
		return
	}

	// Hapus seluruh riwayat servis kendaraan ini
	i = 0
	for i < nServis {
		if dataServis[i].IDKendaraan == id {
			for j = i; j < nServis-1; j++ {
				dataServis[j] = dataServis[j+1]
			}
			nServis--
		} else {
			i++
		}
	}

	// Geser data kendaraan
	for i = idx; i < nKendaraan-1; i++ {
		dataKendaraan[i] = dataKendaraan[i+1]
	}
	nKendaraan--
	fmt.Println("Data kendaraan dan riwayat servisnya berhasil dihapus.")
}

// Menu utama pengelolaan data kendaraan
func MenuKendaraan() {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("--- MENU DATA KENDARAAN ---")
		fmt.Println("1. Tambah Kendaraan")
		fmt.Println("2. Ubah Kendaraan")
		fmt.Println("3. Hapus Kendaraan")
		fmt.Println("4. Tampilkan Semua Kendaraan")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahKendaraan()
		} else if pilihan == 2 {
			ubahKendaraan()
		} else if pilihan == 3 {
			hapusKendaraan()
		} else if pilihan == 4 {
			tampilkanSemuaKendaraan()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
