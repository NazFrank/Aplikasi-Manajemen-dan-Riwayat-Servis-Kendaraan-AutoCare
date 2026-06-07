/*
	Module ini berisi fungsi-fungsi CRUD untuk mengelola riwayat servis kendaraan.
	Mencakup pencarian index servis, pengambilan plat nomor kendaraan, tampilan riwayat
	servis secara keseluruhan maupun per kendaraan, penambahan, pengubahan, dan penghapusan
	data servis. Dilengkapi perhitungan total biaya servis per kendaraan.
*/
package src

import "fmt"

// Mencari index servis berdasarkan ID. Return -1 jika tidak ditemukan
func cariIndexServis(id int) int {
	var i int
	for i = 0; i < nServis; i++ {
		if dataServis[i].ID == id {
			return i
		}
	}
	return -1
}

// Mengambil plat nomor kendaraan berdasarkan id kendaraan
func platKendaraan(idKendaraan int) string {
	var idx int
	idx = cariIndexKendaraan(idKendaraan)
	if idx == -1 {
		return "-"
	}
	return dataKendaraan[idx].PlatNomor
}

// Menampilkan semua riwayat servis
func tampilkanSemuaServis() {
	var i int
	fmt.Println("")
	fmt.Println("--- Daftar Riwayat Servis ---")
	if nServis == 0 {
		fmt.Println("Belum ada riwayat servis.")
		return
	}
	for i = 0; i < nServis; i++ {
		fmt.Printf("ID: %-3d | Plat: %-12s | Tgl: %s | Kerusakan: %-20s | Biaya: Rp %.0f | Ket: %s\n",
			dataServis[i].ID, platKendaraan(dataServis[i].IDKendaraan), dataServis[i].Tanggal,
			dataServis[i].JenisKerusakan, dataServis[i].Biaya, dataServis[i].Keterangan)
	}
}

// Menampilkan riwayat servis untuk satu kendaraan
func tampilkanServisKendaraan() {
	var idKend, i, ditemukan int
	var total float64

	tampilkanSemuaKendaraan()
	if nKendaraan == 0 {
		return
	}

	fmt.Print("Masukkan ID kendaraan: ")
	fmt.Scan(&idKend)
	if cariIndexKendaraan(idKend) == -1 {
		fmt.Println("Kendaraan tidak ditemukan.")
		return
	}

	fmt.Println("")
	fmt.Printf("--- Riwayat Servis Kendaraan %s ---\n", platKendaraan(idKend))
	ditemukan = 0
	total = 0
	for i = 0; i < nServis; i++ {
		if dataServis[i].IDKendaraan == idKend {
			fmt.Printf("ID: %-3d | Tgl: %s | Kerusakan: %-20s | Biaya: Rp %.0f | Ket: %s\n",
				dataServis[i].ID, dataServis[i].Tanggal, dataServis[i].JenisKerusakan,
				dataServis[i].Biaya, dataServis[i].Keterangan)
			total = total + dataServis[i].Biaya
			ditemukan++
		}
	}
	if ditemukan == 0 {
		fmt.Println("Belum ada riwayat servis untuk kendaraan ini.")
	} else {
		fmt.Printf("Total %d servis, Total biaya: Rp %.0f\n", ditemukan, total)
	}
}

// Menambah data riwayat servis
func tambahServis() {
	var s Servis

	if nServis >= MAX {
		fmt.Println("Data servis sudah penuh.")
		return
	}
	if nKendaraan == 0 {
		fmt.Println("Belum ada kendaraan. Silakan tambahkan kendaraan terlebih dahulu.")
		return
	}

	tampilkanSemuaKendaraan()
	fmt.Println("")
	fmt.Println("--- Tambah Riwayat Servis ---")
	fmt.Print("ID Kendaraan                  : ")
	fmt.Scan(&s.IDKendaraan)
	if cariIndexKendaraan(s.IDKendaraan) == -1 {
		fmt.Println("Kendaraan tidak ditemukan. Pembatalan.")
		return
	}

	fmt.Print("Tanggal (YYYY-MM-DD)          : ")
	fmt.Scan(&s.Tanggal)
	fmt.Print("Jenis Kerusakan (pakai _)     : ")
	fmt.Scan(&s.JenisKerusakan)
	fmt.Print("Biaya (Rp)                    : ")
	fmt.Scan(&s.Biaya)
	fmt.Print("Keterangan (pakai _)          : ")
	fmt.Scan(&s.Keterangan)

	s.ID = idServisNext
	dataServis[nServis] = s
	nServis++
	idServisNext++

	fmt.Printf("Riwayat servis berhasil ditambahkan dengan ID %d.\n", s.ID)
}

// Mengubah data riwayat servis
func ubahServis() {
	var id, idx int
	var tgl, jenis, ket string
	var biayaBaru float64

	tampilkanSemuaServis()
	if nServis == 0 {
		return
	}

	fmt.Print("Masukkan ID servis yang akan diubah: ")
	fmt.Scan(&id)
	idx = cariIndexServis(id)
	if idx == -1 {
		fmt.Println("Data servis tidak ditemukan.")
		return
	}

	fmt.Println("Ketik tanda - (minus) jika tidak ingin mengubah field string, atau 0 untuk biaya.")
	fmt.Printf("Tanggal (%s): ", dataServis[idx].Tanggal)
	fmt.Scan(&tgl)
	if tgl != "-" {
		dataServis[idx].Tanggal = tgl
	}
	fmt.Printf("Jenis Kerusakan (%s): ", dataServis[idx].JenisKerusakan)
	fmt.Scan(&jenis)
	if jenis != "-" {
		dataServis[idx].JenisKerusakan = jenis
	}
	fmt.Printf("Biaya (%.0f) [0 untuk skip]: ", dataServis[idx].Biaya)
	fmt.Scan(&biayaBaru)
	if biayaBaru > 0 {
		dataServis[idx].Biaya = biayaBaru
	}
	fmt.Printf("Keterangan (%s): ", dataServis[idx].Keterangan)
	fmt.Scan(&ket)
	if ket != "-" {
		dataServis[idx].Keterangan = ket
	}

	fmt.Println("Data servis berhasil diubah.")
}

// Menghapus data riwayat servis
func hapusServis() {
	var id, idx, i int

	tampilkanSemuaServis()
	if nServis == 0 {
		return
	}

	fmt.Print("Masukkan ID servis yang akan dihapus: ")
	fmt.Scan(&id)
	idx = cariIndexServis(id)
	if idx == -1 {
		fmt.Println("Data servis tidak ditemukan.")
		return
	}

	for i = idx; i < nServis-1; i++ {
		dataServis[i] = dataServis[i+1]
	}
	nServis--
	fmt.Println("Data servis berhasil dihapus.")
}

// Menu utama pengelolaan riwayat servis
func MenuServis() {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("--- MENU RIWAYAT SERVIS ---")
		fmt.Println("1. Tambah Riwayat Servis")
		fmt.Println("2. Ubah Riwayat Servis")
		fmt.Println("3. Hapus Riwayat Servis")
		fmt.Println("4. Tampilkan Riwayat Servis per Kendaraan")
		fmt.Println("5. Tampilkan Semua Riwayat Servis")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahServis()
		} else if pilihan == 2 {
			ubahServis()
		} else if pilihan == 3 {
			hapusServis()
		} else if pilihan == 4 {
			tampilkanServisKendaraan()
		} else if pilihan == 5 {
			tampilkanSemuaServis()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
