package menu

import (
	"autocare/algoritma"
	"autocare/data"
	"autocare/helper"
	"autocare/statistik"
	"fmt"
)

// ============================================================
// MENU UTAMA
// ============================================================

// TampilkanMenuUtama menampilkan pilihan utama aplikasi
func TampilkanMenuUtama() {
	fmt.Println("--------------------------------------")
	fmt.Println("  MENU UTAMA")
	fmt.Println("--------------------------------------")
	fmt.Println("  1. Manajemen Data Kendaraan")
	fmt.Println("  2. Manajemen Data Pemilik")
	fmt.Println("  3. Riwayat Servis Kendaraan")
	fmt.Println("  4. Pencarian Kendaraan")
	fmt.Println("  5. Pengurutan Data Kendaraan")
	fmt.Println("  6. Statistik Servis")
	fmt.Println("  7. Keluar")
	fmt.Println("--------------------------------------")
}

// ============================================================
// MENU KENDARAAN
// ============================================================

// MenuKendaraan menampilkan sub-menu manajemen kendaraan
func MenuKendaraan() {
	var pilihan int
	pilihan = 0

	for pilihan != 5 {
		helper.TampilkanHeaderJudul("MANAJEMEN DATA KENDARAAN")
		fmt.Println("  1. Lihat Semua Kendaraan")
		fmt.Println("  2. Tambah Kendaraan Baru")
		fmt.Println("  3. Ubah Data Kendaraan")
		fmt.Println("  4. Hapus Kendaraan")
		fmt.Println("  5. Kembali ke Menu Utama")
		fmt.Println()

		pilihan = helper.BacaInt("Pilihan: ")
		fmt.Println()

		switch pilihan {
		case 1:
			tampilkanSemuaKendaraan()
		case 2:
			tambahKendaraanBaru()
		case 3:
			ubahDataKendaraan()
		case 4:
			hapusKendaraan()
		case 5:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func tampilkanSemuaKendaraan() {
	helper.TampilkanHeaderJudul("DAFTAR SEMUA KENDARAAN")
	fmt.Printf("  Total kendaraan terdaftar: %d\n\n", data.JumlahKendaraan)
	helper.TampilkanTabelKendaraan(data.DaftarKendaraan, data.JumlahKendaraan)
}

func tambahKendaraanBaru() {
	helper.TampilkanHeaderJudul("TAMBAH KENDARAAN BARU")

	// Tampilkan daftar pemilik yang tersedia
	fmt.Println("  Daftar Pemilik yang Tersedia:")
	helper.TampilkanTabelPemilik()
	fmt.Println()

	var platNomor string
	var merek string
	var model string
	var tahun int
	var idPemilik int
	var tglServis string

	platNomor = helper.BacaString("  Plat Nomor          : ")
	merek = helper.BacaString("  Merek Kendaraan     : ")
	model = helper.BacaString("  Model Kendaraan     : ")
	tahun = helper.BacaInt("  Tahun Produksi      : ")
	idPemilik = helper.BacaInt("  ID Pemilik          : ")
	tglServis = helper.BacaString("  Tgl Servis Terakhir (YYYY-MM-DD): ")

	var berhasil bool
	berhasil = data.TambahKendaraan(platNomor, merek, model, tahun, idPemilik, tglServis)

	if berhasil {
		fmt.Println("\n  [OK] Kendaraan berhasil ditambahkan!")
	} else {
		fmt.Println("\n  [!] Gagal menambahkan kendaraan (data penuh).")
	}
}

func ubahDataKendaraan() {
	helper.TampilkanHeaderJudul("UBAH DATA KENDARAAN")
	tampilkanSemuaKendaraan()
	fmt.Println()

	var id int
	id = helper.BacaInt("  Masukkan ID Kendaraan yang ingin diubah: ")

	// Cek apakah ID kendaraan ada
	var i int
	var ada bool
	ada = false
	for i = 0; i < data.JumlahKendaraan; i++ {
		if data.DaftarKendaraan[i].ID == id {
			ada = true
		}
	}

	if !ada {
		fmt.Println("  [!] ID Kendaraan tidak ditemukan.")
		return
	}

	fmt.Println("\n  Masukkan data baru (kosongkan jika tidak diubah):")

	var platNomor string
	var merek string
	var model string
	var tahun int
	var idPemilik int

	platNomor = helper.BacaString("  Plat Nomor Baru    : ")
	merek = helper.BacaString("  Merek Baru         : ")
	model = helper.BacaString("  Model Baru         : ")
	tahun = helper.BacaInt("  Tahun Produksi Baru: ")
	idPemilik = helper.BacaInt("  ID Pemilik Baru    : ")

	var berhasil bool
	berhasil = data.UbahKendaraan(id, platNomor, merek, model, tahun, idPemilik)

	if berhasil {
		fmt.Println("\n  [OK] Data kendaraan berhasil diubah!")
	} else {
		fmt.Println("\n  [!] Gagal mengubah data kendaraan.")
	}
}

func hapusKendaraan() {
	helper.TampilkanHeaderJudul("HAPUS KENDARAAN")
	tampilkanSemuaKendaraan()
	fmt.Println()

	var id int
	id = helper.BacaInt("  Masukkan ID Kendaraan yang ingin dihapus: ")

	var konfirmasi string
	konfirmasi = helper.BacaString("  Yakin ingin menghapus? (ya/tidak): ")

	if konfirmasi == "ya" {
		var berhasil bool
		berhasil = data.HapusKendaraan(id)
		if berhasil {
			fmt.Println("  [OK] Kendaraan berhasil dihapus!")
		} else {
			fmt.Println("  [!] ID Kendaraan tidak ditemukan.")
		}
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

// ============================================================
// MENU PEMILIK
// ============================================================

// MenuPemilik menampilkan sub-menu manajemen data pemilik
func MenuPemilik() {
	var pilihan int
	pilihan = 0

	for pilihan != 4 {
		helper.TampilkanHeaderJudul("MANAJEMEN DATA PEMILIK")
		fmt.Println("  1. Lihat Semua Pemilik")
		fmt.Println("  2. Tambah Pemilik Baru")
		fmt.Println("  3. Ubah Data Pemilik")
		fmt.Println("  4. Kembali ke Menu Utama")
		fmt.Println()

		pilihan = helper.BacaInt("Pilihan: ")
		fmt.Println()

		switch pilihan {
		case 1:
			tampilkanSemuaPemilik()
		case 2:
			tambahPemilikBaru()
		case 3:
			ubahDataPemilik()
		case 4:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func tampilkanSemuaPemilik() {
	helper.TampilkanHeaderJudul("DAFTAR SEMUA PEMILIK")
	fmt.Printf("  Total pemilik terdaftar: %d\n\n", data.JumlahPemilik)
	helper.TampilkanTabelPemilik()
}

func tambahPemilikBaru() {
	helper.TampilkanHeaderJudul("TAMBAH PEMILIK BARU")

	var nama string
	var noTelp string
	var alamat string

	nama = helper.BacaString("  Nama Pemilik : ")
	noTelp = helper.BacaString("  No. Telepon  : ")
	alamat = helper.BacaString("  Alamat       : ")

	var berhasil bool
	berhasil = data.TambahPemilik(nama, noTelp, alamat)

	if berhasil {
		fmt.Println("\n  [OK] Pemilik berhasil ditambahkan!")
	} else {
		fmt.Println("\n  [!] Gagal menambahkan pemilik (data penuh).")
	}
}

func ubahDataPemilik() {
	helper.TampilkanHeaderJudul("UBAH DATA PEMILIK")
	tampilkanSemuaPemilik()
	fmt.Println()

	var id int
	id = helper.BacaInt("  Masukkan ID Pemilik yang ingin diubah: ")

	var nama string
	var noTelp string
	var alamat string

	nama = helper.BacaString("  Nama Baru       : ")
	noTelp = helper.BacaString("  No. Telepon Baru: ")
	alamat = helper.BacaString("  Alamat Baru     : ")

	var berhasil bool
	berhasil = data.UbahPemilik(id, nama, noTelp, alamat)

	if berhasil {
		fmt.Println("\n  [OK] Data pemilik berhasil diubah!")
	} else {
		fmt.Println("\n  [!] ID Pemilik tidak ditemukan.")
	}
}

// ============================================================
// MENU SERVIS
// ============================================================

// MenuServis menampilkan sub-menu riwayat servis
func MenuServis() {
	var pilihan int
	pilihan = 0

	for pilihan != 4 {
		helper.TampilkanHeaderJudul("RIWAYAT SERVIS KENDARAAN")
		fmt.Println("  1. Lihat Riwayat Servis Semua Kendaraan")
		fmt.Println("  2. Lihat Riwayat Servis per Kendaraan")
		fmt.Println("  3. Tambah Catatan Servis Baru")
		fmt.Println("  4. Kembali ke Menu Utama")
		fmt.Println()

		pilihan = helper.BacaInt("Pilihan: ")
		fmt.Println()

		switch pilihan {
		case 1:
			tampilkanSemuaServis()
		case 2:
			tampilkanServisPerKendaraan()
		case 3:
			tambahCatatanServis()
		case 4:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func tampilkanSemuaServis() {
	helper.TampilkanHeaderJudul("SEMUA RIWAYAT SERVIS")
	fmt.Printf("  Total catatan servis: %d\n\n", data.JumlahServis)

	fmt.Printf("  %-4s %-6s %-12s %-15s %-25s %-12s\n",
		"ID", "ID Kend", "Tanggal", "Jenis Kerusakan", "Deskripsi", "Biaya (Rp)")
	helper.TampilkanGarisTipis()

	var i int
	for i = 0; i < data.JumlahServis; i++ {
		fmt.Printf("  %-4d %-6d %-12s %-15s %-25s %-12d\n",
			data.DaftarServis[i].ID,
			data.DaftarServis[i].IDKendaraan,
			data.DaftarServis[i].TanggalServis,
			data.DaftarServis[i].JenisKerusakan,
			data.DaftarServis[i].Deskripsi,
			data.DaftarServis[i].BiayaServis)
	}
}

func tampilkanServisPerKendaraan() {
	helper.TampilkanHeaderJudul("RIWAYAT SERVIS PER KENDARAAN")

	// Tampilkan daftar kendaraan
	helper.TampilkanTabelKendaraan(data.DaftarKendaraan, data.JumlahKendaraan)
	fmt.Println()

	var idKendaraan int
	idKendaraan = helper.BacaInt("  Masukkan ID Kendaraan: ")

	// Cari dan tampilkan info kendaraan
	var i int
	for i = 0; i < data.JumlahKendaraan; i++ {
		if data.DaftarKendaraan[i].ID == idKendaraan {
			fmt.Println()
			fmt.Println("  Detail Kendaraan:")
			helper.TampilkanDetailKendaraan(data.DaftarKendaraan[i])
			fmt.Println()
			fmt.Println("  Riwayat Servis:")
			helper.TampilkanRiwayatServisKendaraan(idKendaraan)
			return
		}
	}

	fmt.Println("  [!] ID Kendaraan tidak ditemukan.")
}

func tambahCatatanServis() {
	helper.TampilkanHeaderJudul("TAMBAH CATATAN SERVIS BARU")

	// Tampilkan daftar kendaraan
	helper.TampilkanTabelKendaraan(data.DaftarKendaraan, data.JumlahKendaraan)
	fmt.Println()

	var idKendaraan int
	var tanggal string
	var jenisKerusakan string
	var deskripsi string
	var biaya int

	idKendaraan = helper.BacaInt("  ID Kendaraan          : ")
	tanggal = helper.BacaString("  Tanggal Servis (YYYY-MM-DD): ")

	fmt.Println("  Jenis Kerusakan yang tersedia:")
	fmt.Println("  [Mesin, Rem, Kelistrikan, Kaki-kaki, Body, Lainnya]")
	jenisKerusakan = helper.BacaString("  Jenis Kerusakan       : ")
	deskripsi = helper.BacaString("  Deskripsi Perbaikan   : ")
	biaya = helper.BacaInt("  Biaya Servis (Rp)     : ")

	var berhasil bool
	berhasil = data.TambahServis(idKendaraan, tanggal, jenisKerusakan, deskripsi, biaya)

	if berhasil {
		fmt.Println("\n  [OK] Catatan servis berhasil ditambahkan!")
	} else {
		fmt.Println("\n  [!] Gagal menambahkan catatan servis.")
	}
}

// ============================================================
// MENU PENCARIAN
// ============================================================

// MenuCariKendaraan menampilkan menu pencarian kendaraan
func MenuCariKendaraan() {
	var pilihan int
	pilihan = 0

	for pilihan != 3 {
		helper.TampilkanHeaderJudul("PENCARIAN KENDARAAN")
		fmt.Println("  1. Sequential Search (cari satu per satu)")
		fmt.Println("  2. Binary Search (cari dengan bagi dua)")
		fmt.Println("  3. Kembali ke Menu Utama")
		fmt.Println()
		fmt.Println("  Catatan: Binary Search memerlukan data terurut")
		fmt.Println()

		pilihan = helper.BacaInt("Pilihan: ")
		fmt.Println()

		switch pilihan {
		case 1:
			cariDenganSequential()
		case 2:
			cariDenganBinary()
		case 3:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func cariDenganSequential() {
	helper.TampilkanHeaderJudul("PENCARIAN SEQUENTIAL")
	fmt.Println("  Sequential Search memeriksa setiap kendaraan")
	fmt.Println("  satu per satu dari awal hingga akhir data.")
	fmt.Println()

	var platNomor string
	platNomor = helper.BacaString("  Masukkan Plat Nomor yang dicari: ")

	fmt.Println()
	fmt.Println("  [Proses] Menjalankan Sequential Search...")

	var indeks int
	indeks = algoritma.CariSequential(platNomor)

	if indeks != -1 {
		fmt.Printf("  [OK] Kendaraan ditemukan di posisi indeks ke-%d!\n\n", indeks)
		fmt.Println("  Detail Kendaraan:")
		helper.TampilkanDetailKendaraan(data.DaftarKendaraan[indeks])
	} else {
		fmt.Printf("  [!] Kendaraan dengan plat nomor '%s' tidak ditemukan.\n", platNomor)
	}
}

func cariDenganBinary() {
	helper.TampilkanHeaderJudul("PENCARIAN BINARY SEARCH")
	fmt.Println("  Binary Search bekerja dengan membagi area pencarian")
	fmt.Println("  menjadi dua bagian secara berulang (membutuhkan data terurut).")
	fmt.Println()
	fmt.Println("  [Persiapan] Mengurutkan data berdasarkan plat nomor...")
	algoritma.SalinDanUrutkanUntukBinarySearch()
	fmt.Println("  [OK] Data berhasil diurutkan untuk Binary Search.")
	fmt.Println()

	var platNomor string
	platNomor = helper.BacaString("  Masukkan Plat Nomor yang dicari: ")

	fmt.Println()
	fmt.Println("  [Proses] Menjalankan Binary Search...")

	var indeks int
	indeks = algoritma.CariBinary(platNomor)

	if indeks != -1 {
		fmt.Printf("  [OK] Kendaraan ditemukan di posisi indeks ke-%d (data terurut)!\n\n", indeks)
		fmt.Println("  Detail Kendaraan (dari data yang sudah diurutkan):")
		helper.TampilkanDetailKendaraan(algoritma.KendaraanTerurut[indeks])
	} else {
		fmt.Printf("  [!] Kendaraan dengan plat nomor '%s' tidak ditemukan.\n", platNomor)
	}
}

// ============================================================
// MENU PENGURUTAN
// ============================================================

// MenuUrutkanKendaraan menampilkan menu pengurutan kendaraan
func MenuUrutkanKendaraan() {
	var pilihan int
	pilihan = 0

	for pilihan != 3 {
		helper.TampilkanHeaderJudul("PENGURUTAN DATA KENDARAAN")
		fmt.Println("  1. Selection Sort - Urutkan berdasarkan Tahun Produksi")
		fmt.Println("  2. Insertion Sort - Urutkan berdasarkan Tanggal Servis Terakhir")
		fmt.Println("  3. Kembali ke Menu Utama")
		fmt.Println()

		pilihan = helper.BacaInt("Pilihan: ")
		fmt.Println()

		switch pilihan {
		case 1:
			urutkanDenganSelectionSort()
		case 2:
			urutkanDenganInsertionSort()
		case 3:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func urutkanDenganSelectionSort() {
	helper.TampilkanHeaderJudul("SELECTION SORT - BERDASARKAN TAHUN PRODUKSI")
	fmt.Println("  Cara kerja Selection Sort:")
	fmt.Println("  - Cari elemen terkecil dari semua data")
	fmt.Println("  - Tempatkan di posisi pertama")
	fmt.Println("  - Ulangi untuk sisa data yang belum terurut")
	fmt.Println()

	fmt.Println("  [Sebelum Pengurutan]")
	helper.TampilkanTabelKendaraan(data.DaftarKendaraan, data.JumlahKendaraan)

	fmt.Println()
	fmt.Println("  [Proses] Menjalankan Selection Sort...")
	algoritma.SelectionSortTahun()
	fmt.Println("  [OK] Pengurutan selesai!")

	fmt.Println()
	fmt.Println("  [Hasil Pengurutan - Ascending berdasarkan Tahun Produksi]")
	helper.TampilkanTabelKendaraan(algoritma.HasilUrutan, algoritma.JumlahHasilUrutan)
	fmt.Println()
	fmt.Println("  (Data asli tidak berubah, ini hanya tampilan terurut)")
}

func urutkanDenganInsertionSort() {
	helper.TampilkanHeaderJudul("INSERTION SORT - BERDASARKAN TANGGAL SERVIS TERAKHIR")
	fmt.Println("  Cara kerja Insertion Sort:")
	fmt.Println("  - Ambil satu elemen dari bagian yang belum terurut")
	fmt.Println("  - Sisipkan ke posisi yang tepat di bagian yang sudah terurut")
	fmt.Println("  - Mirip seperti mengurutkan kartu remi")
	fmt.Println()

	fmt.Println("  [Sebelum Pengurutan]")
	helper.TampilkanTabelKendaraan(data.DaftarKendaraan, data.JumlahKendaraan)

	fmt.Println()
	fmt.Println("  [Proses] Menjalankan Insertion Sort...")
	algoritma.InsertionSortTanggalServis()
	fmt.Println("  [OK] Pengurutan selesai!")

	fmt.Println()
	fmt.Println("  [Hasil Pengurutan - Ascending berdasarkan Tanggal Servis Terakhir]")
	helper.TampilkanTabelKendaraan(algoritma.HasilUrutan, algoritma.JumlahHasilUrutan)
	fmt.Println()
	fmt.Println("  (Data asli tidak berubah, ini hanya tampilan terurut)")
}

// ============================================================
// MENU STATISTIK
// ============================================================

// MenuStatistik menampilkan menu statistik servis
func MenuStatistik() {
	var pilihan int
	pilihan = 0

	for pilihan != 3 {
		helper.TampilkanHeaderJudul("STATISTIK SERVIS KENDARAAN")
		fmt.Println("  1. Statistik Jumlah Servis per Bulan")
		fmt.Println("  2. Statistik Kategori Kerusakan Terbanyak")
		fmt.Println("  3. Kembali ke Menu Utama")
		fmt.Println()

		pilihan = helper.BacaInt("Pilihan: ")
		fmt.Println()

		switch pilihan {
		case 1:
			tampilkanStatistikBulanan()
		case 2:
			tampilkanStatistikKerusakan()
		case 3:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func tampilkanStatistikBulanan() {
	helper.TampilkanHeaderJudul("STATISTIK SERVIS PER BULAN")

	var tahun string
	tahun = helper.BacaString("  Masukkan Tahun (contoh: 2024): ")

	statistik.HitungServisPerBulan(tahun)
}

func tampilkanStatistikKerusakan() {
	helper.TampilkanHeaderJudul("STATISTIK KATEGORI KERUSAKAN")
	statistik.HitungKategoriKerusakan()
}
