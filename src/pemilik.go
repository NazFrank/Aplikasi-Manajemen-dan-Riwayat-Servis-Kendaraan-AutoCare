/*
	Module ini berisi fungsi-fungsi CRUD (Create, Read, Update, Delete) untuk mengelola
	data pemilik kendaraan. Mencakup pencarian index pemilik, tampilan daftar pemilik,
	penambahan, pengubahan, dan penghapusan data pemilik, serta menu interaktif pengelolaan
	data pemilik. Penghapusan pemilik dilindungi oleh validasi kepemilikan kendaraan aktif.
*/
package src

import "fmt"

// Mencari index pemilik berdasarkan ID. Return -1 jika tidak ditemukan
func cariIndexPemilik(id int) int {
	var i int
	for i = 0; i < nPemilik; i++ {
		if dataPemilik[i].ID == id {
			return i
		}
	}
	return -1
}

// Menampilkan seluruh data pemilik
func tampilkanSemuaPemilik() {
	var i int
	fmt.Println("")
	fmt.Println("--- Daftar Pemilik Kendaraan ---")
	if nPemilik == 0 {
		fmt.Println("Belum ada data pemilik.")
		return
	}
	for i = 0; i < nPemilik; i++ {
		fmt.Printf("ID: %-3d | Nama: %-20s | Alamat: %-25s | No HP: %s\n",
			dataPemilik[i].ID, dataPemilik[i].Nama, dataPemilik[i].Alamat, dataPemilik[i].NoHP)
	}
}

// Menambahkan data pemilik baru
func tambahPemilik() {
	var p Pemilik

	if nPemilik >= MAX {
		fmt.Println("Data pemilik sudah penuh.")
		return
	}

	fmt.Println("")
	fmt.Println("--- Tambah Pemilik Baru ---")
	fmt.Print("Nama (pakai _ utk spasi)   : ")
	fmt.Scan(&p.Nama)
	fmt.Print("Alamat (pakai _ utk spasi) : ")
	fmt.Scan(&p.Alamat)
	fmt.Print("No HP                      : ")
	fmt.Scan(&p.NoHP)

	p.ID = idPemilikNext
	dataPemilik[nPemilik] = p
	nPemilik++
	idPemilikNext++

	fmt.Printf("Pemilik berhasil ditambahkan dengan ID %d.\n", p.ID)
}

// Mengubah data pemilik berdasarkan ID
func ubahPemilik() {
	var id, idx int
	var nama, alamat, noHP string

	tampilkanSemuaPemilik()
	if nPemilik == 0 {
		return
	}

	fmt.Print("Masukkan ID pemilik yang akan diubah: ")
	fmt.Scan(&id)
	idx = cariIndexPemilik(id)
	if idx == -1 {
		fmt.Println("Pemilik dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Ketik tanda - (minus) jika tidak ingin mengubah field tersebut.")
	fmt.Printf("Nama (%s): ", dataPemilik[idx].Nama)
	fmt.Scan(&nama)
	if nama != "-" {
		dataPemilik[idx].Nama = nama
	}
	fmt.Printf("Alamat (%s): ", dataPemilik[idx].Alamat)
	fmt.Scan(&alamat)
	if alamat != "-" {
		dataPemilik[idx].Alamat = alamat
	}
	fmt.Printf("No HP (%s): ", dataPemilik[idx].NoHP)
	fmt.Scan(&noHP)
	if noHP != "-" {
		dataPemilik[idx].NoHP = noHP
	}

	fmt.Println("Data pemilik berhasil diubah.")
}

// Menghapus data pemilik berdasarkan ID
func hapusPemilik() {
	var id, idx, i int

	tampilkanSemuaPemilik()
	if nPemilik == 0 {
		return
	}

	fmt.Print("Masukkan ID pemilik yang akan dihapus: ")
	fmt.Scan(&id)
	idx = cariIndexPemilik(id)
	if idx == -1 {
		fmt.Println("Pemilik dengan ID tersebut tidak ditemukan.")
		return
	}

	// Cek apakah pemilik masih memiliki kendaraan
	for i = 0; i < nKendaraan; i++ {
		if dataKendaraan[i].IDPemilik == id {
			fmt.Println("Pemilik tidak dapat dihapus karena masih memiliki kendaraan terdaftar.")
			return
		}
	}

	// Geser data setelahnya ke kiri
	for i = idx; i < nPemilik-1; i++ {
		dataPemilik[i] = dataPemilik[i+1]
	}
	nPemilik--
	fmt.Println("Data pemilik berhasil dihapus.")
}

// Menu utama pengelolaan data pemilik
func MenuPemilik() {
	var pilihan int
	for {
		fmt.Println("")
		fmt.Println("--- MENU DATA PEMILIK ---")
		fmt.Println("1. Tambah Pemilik")
		fmt.Println("2. Ubah Pemilik")
		fmt.Println("3. Hapus Pemilik")
		fmt.Println("4. Tampilkan Semua Pemilik")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahPemilik()
		} else if pilihan == 2 {
			ubahPemilik()
		} else if pilihan == 3 {
			hapusPemilik()
		} else if pilihan == 4 {
			tampilkanSemuaPemilik()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
