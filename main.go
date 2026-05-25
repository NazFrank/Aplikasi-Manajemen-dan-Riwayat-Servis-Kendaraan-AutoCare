package main

import "fmt"

// ======== STRUCT ========

type Pemilik struct {
	ID      int
	Nama    string
	Telepon string
	Alamat  string
}

type Kendaraan struct {
	ID            int
	PlatNomor     string
	Merek         string
	Model         string
	TahunProduksi int
	IDPemilik     int
}

type ServisRecord struct {
	ID             int
	IDKendaraan    int
	TanggalServis  string
	JenisKerusakan string
	Deskripsi      string
	Biaya          int
}

// ======== DATA GLOBAL ========

var (
	daftarPemilik [100]Pemilik
	jumlahPemilik int

	daftarKendaraan [100]Kendaraan
	jumlahKendaraan int

	daftarServis [500]ServisRecord
	jumlahServis int

	idPemilikNext   int
	idKendaraanNext int
	idServisNext    int
)

// ======== MAIN ========

func main() {
	var (
		pilihan int
	)

	idPemilikNext = 1
	idKendaraanNext = 1
	idServisNext = 1
	jumlahPemilik = 0
	jumlahKendaraan = 0
	jumlahServis = 0

	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("         AUTOCARE - Manajemen Servis   ")
		fmt.Println("========================================")
		fmt.Println("1. Manajemen Pemilik")
		fmt.Println("2. Manajemen Kendaraan")
		fmt.Println("3. Manajemen Servis")
		fmt.Println("4. Pencarian Kendaraan")
		fmt.Println("5. Pengurutan Kendaraan")
		fmt.Println("6. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan AutoCare.")
			break
		} else if pilihan == 1 {

			// ---- MENU PEMILIK ----
			var subPilihan int
			fmt.Println()
			fmt.Println("--- Manajemen Pemilik ---")
			fmt.Println("1. Tambah Pemilik")
			fmt.Println("2. Ubah Pemilik")
			fmt.Println("3. Hapus Pemilik")
			fmt.Println("4. Tampilkan Semua Pemilik")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				// Tambah Pemilik
				var (
					nama    string
					telepon string
					alamat  string
				)
				fmt.Print("Nama    : ")
				fmt.Scan(&nama)
				fmt.Print("Telepon : ")
				fmt.Scan(&telepon)
				fmt.Print("Alamat  : ")
				fmt.Scan(&alamat)

				daftarPemilik[jumlahPemilik].ID = idPemilikNext
				daftarPemilik[jumlahPemilik].Nama = nama
				daftarPemilik[jumlahPemilik].Telepon = telepon
				daftarPemilik[jumlahPemilik].Alamat = alamat

				jumlahPemilik++
				idPemilikNext++

				fmt.Println("Pemilik berhasil ditambahkan.")

			} else if subPilihan == 2 {
				// Ubah Pemilik
				var (
					idCari  int
					nama    string
					telepon string
					alamat  string
					i       int
					ketemu  bool
				)
				fmt.Print("Masukkan ID Pemilik yang ingin diubah: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < jumlahPemilik {
					if daftarPemilik[i].ID == idCari {
						ketemu = true
						fmt.Print("Nama baru    : ")
						fmt.Scan(&nama)
						fmt.Print("Telepon baru : ")
						fmt.Scan(&telepon)
						fmt.Print("Alamat baru  : ")
						fmt.Scan(&alamat)

						daftarPemilik[i].Nama = nama
						daftarPemilik[i].Telepon = telepon
						daftarPemilik[i].Alamat = alamat

						fmt.Println("Data pemilik berhasil diubah.")
					}
					i++
				}
				if ketemu == false {
					fmt.Println("Pemilik dengan ID tersebut tidak ditemukan.")
				}

			} else if subPilihan == 3 {
				// Hapus Pemilik
				var (
					idCari int
					i      int
					j      int
					ketemu bool
				)
				fmt.Print("Masukkan ID Pemilik yang ingin dihapus: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < jumlahPemilik {
					if daftarPemilik[i].ID == idCari {
						ketemu = true
						j = i
						for j < jumlahPemilik-1 {
							daftarPemilik[j] = daftarPemilik[j+1]
							j++
						}
						jumlahPemilik--
						fmt.Println("Pemilik berhasil dihapus.")
					}
					i++
				}
				if ketemu == false {
					fmt.Println("Pemilik dengan ID tersebut tidak ditemukan.")
				}

			} else if subPilihan == 4 {
				// Tampilkan Semua Pemilik
				var i int
				if jumlahPemilik == 0 {
					fmt.Println("Belum ada data pemilik.")
				} else {
					fmt.Println()
					fmt.Println("Daftar Pemilik:")
					fmt.Println("ID  | Nama            | Telepon       | Alamat")
					i = 0
					for i < jumlahPemilik {
						fmt.Printf("%-4d| %-16s| %-14s| %s\n",
							daftarPemilik[i].ID,
							daftarPemilik[i].Nama,
							daftarPemilik[i].Telepon,
							daftarPemilik[i].Alamat)
						i++
					}
				}
			}

		} else if pilihan == 2 {

			// ---- MENU KENDARAAN ----
			var subPilihan int
			fmt.Println()
			fmt.Println("--- Manajemen Kendaraan ---")
			fmt.Println("1. Tambah Kendaraan")
			fmt.Println("2. Ubah Kendaraan")
			fmt.Println("3. Hapus Kendaraan")
			fmt.Println("4. Tampilkan Semua Kendaraan")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				// Tambah Kendaraan
				var (
					platNomor     string
					merek         string
					model         string
					tahunProduksi int
					idPemilik     int
				)
				fmt.Print("Plat Nomor    : ")
				fmt.Scan(&platNomor)
				fmt.Print("Merek         : ")
				fmt.Scan(&merek)
				fmt.Print("Model         : ")
				fmt.Scan(&model)
				fmt.Print("Tahun Produksi: ")
				fmt.Scan(&tahunProduksi)
				fmt.Print("ID Pemilik    : ")
				fmt.Scan(&idPemilik)

				daftarKendaraan[jumlahKendaraan].ID = idKendaraanNext
				daftarKendaraan[jumlahKendaraan].PlatNomor = platNomor
				daftarKendaraan[jumlahKendaraan].Merek = merek
				daftarKendaraan[jumlahKendaraan].Model = model
				daftarKendaraan[jumlahKendaraan].TahunProduksi = tahunProduksi
				daftarKendaraan[jumlahKendaraan].IDPemilik = idPemilik

				jumlahKendaraan++
				idKendaraanNext++

				fmt.Println("Kendaraan berhasil ditambahkan.")

			} else if subPilihan == 2 {
				// Ubah Kendaraan
				var (
					idCari        int
					platNomor     string
					merek         string
					model         string
					tahunProduksi int
					idPemilik     int
					i             int
					ketemu        bool
				)
				fmt.Print("Masukkan ID Kendaraan yang ingin diubah: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < jumlahKendaraan {
					if daftarKendaraan[i].ID == idCari {
						ketemu = true
						fmt.Print("Plat Nomor baru    : ")
						fmt.Scan(&platNomor)
						fmt.Print("Merek baru         : ")
						fmt.Scan(&merek)
						fmt.Print("Model baru         : ")
						fmt.Scan(&model)
						fmt.Print("Tahun Produksi baru: ")
						fmt.Scan(&tahunProduksi)
						fmt.Print("ID Pemilik baru    : ")
						fmt.Scan(&idPemilik)

						daftarKendaraan[i].PlatNomor = platNomor
						daftarKendaraan[i].Merek = merek
						daftarKendaraan[i].Model = model
						daftarKendaraan[i].TahunProduksi = tahunProduksi
						daftarKendaraan[i].IDPemilik = idPemilik

						fmt.Println("Data kendaraan berhasil diubah.")
					}
					i++
				}
				if ketemu == false {
					fmt.Println("Kendaraan dengan ID tersebut tidak ditemukan.")
				}

			} else if subPilihan == 3 {
				// Hapus Kendaraan
				var (
					idCari int
					i      int
					j      int
					ketemu bool
				)
				fmt.Print("Masukkan ID Kendaraan yang ingin dihapus: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < jumlahKendaraan {
					if daftarKendaraan[i].ID == idCari {
						ketemu = true
						j = i
						for j < jumlahKendaraan-1 {
							daftarKendaraan[j] = daftarKendaraan[j+1]
							j++
						}
						jumlahKendaraan--
						fmt.Println("Kendaraan berhasil dihapus.")
					}
					i++
				}
				if ketemu == false {
					fmt.Println("Kendaraan dengan ID tersebut tidak ditemukan.")
				}

			} else if subPilihan == 4 {
				// Tampilkan Semua Kendaraan
				var i int
				if jumlahKendaraan == 0 {
					fmt.Println("Belum ada data kendaraan.")
				} else {
					fmt.Println()
					fmt.Println("Daftar Kendaraan:")
					fmt.Println("ID  | Plat       | Merek       | Model       | Tahun | ID Pemilik")
					i = 0
					for i < jumlahKendaraan {
						fmt.Printf("%-4d| %-11s| %-12s| %-12s| %-6d| %d\n",
							daftarKendaraan[i].ID,
							daftarKendaraan[i].PlatNomor,
							daftarKendaraan[i].Merek,
							daftarKendaraan[i].Model,
							daftarKendaraan[i].TahunProduksi,
							daftarKendaraan[i].IDPemilik)
						i++
					}
				}
			}

		} else if pilihan == 3 {

			// ---- MENU SERVIS ----
			var subPilihan int
			fmt.Println()
			fmt.Println("--- Manajemen Servis ---")
			fmt.Println("1. Tambah Servis")
			fmt.Println("2. Ubah Servis")
			fmt.Println("3. Hapus Servis")
			fmt.Println("4. Tampilkan Semua Servis")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				// Tambah Servis
				var (
					idKendaraan    int
					tanggal        string
					jenisKerusakan string
					deskripsi      string
					biaya          int
				)
				fmt.Print("ID Kendaraan    : ")
				fmt.Scan(&idKendaraan)
				fmt.Print("Tanggal (YYYY-MM-DD): ")
				fmt.Scan(&tanggal)
				fmt.Print("Jenis Kerusakan : ")
				fmt.Scan(&jenisKerusakan)
				fmt.Print("Deskripsi       : ")
				fmt.Scan(&deskripsi)
				fmt.Print("Biaya (Rp)      : ")
				fmt.Scan(&biaya)

				daftarServis[jumlahServis].ID = idServisNext
				daftarServis[jumlahServis].IDKendaraan = idKendaraan
				daftarServis[jumlahServis].TanggalServis = tanggal
				daftarServis[jumlahServis].JenisKerusakan = jenisKerusakan
				daftarServis[jumlahServis].Deskripsi = deskripsi
				daftarServis[jumlahServis].Biaya = biaya

				jumlahServis++
				idServisNext++

				fmt.Println("Data servis berhasil ditambahkan.")

			} else if subPilihan == 2 {
				// Ubah Servis
				var (
					idCari         int
					idKendaraan    int
					tanggal        string
					jenisKerusakan string
					deskripsi      string
					biaya          int
					i              int
					ketemu         bool
				)
				fmt.Print("Masukkan ID Servis yang ingin diubah: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < jumlahServis {
					if daftarServis[i].ID == idCari {
						ketemu = true
						fmt.Print("ID Kendaraan baru    : ")
						fmt.Scan(&idKendaraan)
						fmt.Print("Tanggal baru (YYYY-MM-DD): ")
						fmt.Scan(&tanggal)
						fmt.Print("Jenis Kerusakan baru : ")
						fmt.Scan(&jenisKerusakan)
						fmt.Print("Deskripsi baru       : ")
						fmt.Scan(&deskripsi)
						fmt.Print("Biaya baru (Rp)      : ")
						fmt.Scan(&biaya)

						daftarServis[i].IDKendaraan = idKendaraan
						daftarServis[i].TanggalServis = tanggal
						daftarServis[i].JenisKerusakan = jenisKerusakan
						daftarServis[i].Deskripsi = deskripsi
						daftarServis[i].Biaya = biaya

						fmt.Println("Data servis berhasil diubah.")
					}
					i++
				}
				if ketemu == false {
					fmt.Println("Servis dengan ID tersebut tidak ditemukan.")
				}

			} else if subPilihan == 3 {
				// Hapus Servis
				var (
					idCari int
					i      int
					j      int
					ketemu bool
				)
				fmt.Print("Masukkan ID Servis yang ingin dihapus: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < jumlahServis {
					if daftarServis[i].ID == idCari {
						ketemu = true
						j = i
						for j < jumlahServis-1 {
							daftarServis[j] = daftarServis[j+1]
							j++
						}
						jumlahServis--
						fmt.Println("Data servis berhasil dihapus.")
					}
					i++
				}
				if ketemu == false {
					fmt.Println("Servis dengan ID tersebut tidak ditemukan.")
				}

			} else if subPilihan == 4 {
				// Tampilkan Semua Servis
				var i int
				if jumlahServis == 0 {
					fmt.Println("Belum ada data servis.")
				} else {
					fmt.Println()
					fmt.Println("Daftar Servis:")
					fmt.Println("ID  | ID Kend | Tanggal    | Jenis Kerusakan | Deskripsi       | Biaya")
					i = 0
					for i < jumlahServis {
						fmt.Printf("%-4d| %-8d| %-11s| %-16s| %-16s| %d\n",
							daftarServis[i].ID,
							daftarServis[i].IDKendaraan,
							daftarServis[i].TanggalServis,
							daftarServis[i].JenisKerusakan,
							daftarServis[i].Deskripsi,
							daftarServis[i].Biaya)
						i++
					}
				}
			}

		} else if pilihan == 4 {

			// ---- PENCARIAN KENDARAAN ----
			var (
				subPilihan int
				cariPlat   string
			)
			fmt.Println()
			fmt.Println("--- Pencarian Kendaraan ---")
			fmt.Println("1. Sequential Search (berdasarkan Plat Nomor)")
			fmt.Println("2. Binary Search (berdasarkan Plat Nomor - data harus sudah diurutkan)")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			fmt.Print("Masukkan Plat Nomor yang dicari: ")
			fmt.Scan(&cariPlat)

			if subPilihan == 1 {
				// Sequential Search
				var (
					i      int
					ketemu bool
					idx    int
				)
				ketemu = false
				i = 0
				for i < jumlahKendaraan {
					if daftarKendaraan[i].PlatNomor == cariPlat {
						ketemu = true
						idx = i
					}
					i++
				}
				if ketemu {
					fmt.Println("Kendaraan ditemukan (Sequential Search):")
					fmt.Printf("ID: %d | Plat: %s | Merek: %s | Model: %s | Tahun: %d | ID Pemilik: %d\n",
						daftarKendaraan[idx].ID,
						daftarKendaraan[idx].PlatNomor,
						daftarKendaraan[idx].Merek,
						daftarKendaraan[idx].Model,
						daftarKendaraan[idx].TahunProduksi,
						daftarKendaraan[idx].IDPemilik)
				} else {
					fmt.Println("Kendaraan dengan plat nomor tersebut tidak ditemukan.")
				}

			} else if subPilihan == 2 {
				// Binary Search (array harus terurut berdasarkan PlatNomor)
				// Urutkan dulu sementara untuk binary search (tidak mengubah data asli)
				var (
					tempArr   [100]Kendaraan
					i         int
					j         int
					tempData  Kendaraan
					kiri      int
					kanan     int
					tengah    int
					ketemu    bool
					idxKetemu int
				)

				// Salin data ke tempArr
				i = 0
				for i < jumlahKendaraan {
					tempArr[i] = daftarKendaraan[i]
					i++
				}

				// Insertion sort berdasarkan PlatNomor untuk keperluan binary search
				i = 1
				for i < jumlahKendaraan {
					tempData = tempArr[i]
					j = i - 1
					for j >= 0 && tempArr[j].PlatNomor > tempData.PlatNomor {
						tempArr[j+1] = tempArr[j]
						j--
					}
					tempArr[j+1] = tempData
					i++
				}

				// Binary Search
				kiri = 0
				kanan = jumlahKendaraan - 1
				ketemu = false
				idxKetemu = -1

				for kiri <= kanan {
					tengah = (kiri + kanan) / 2
					if tempArr[tengah].PlatNomor == cariPlat {
						ketemu = true
						idxKetemu = tengah
						kiri = kanan + 1
					} else if tempArr[tengah].PlatNomor < cariPlat {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}

				if ketemu {
					fmt.Println("Kendaraan ditemukan (Binary Search):")
					fmt.Printf("ID: %d | Plat: %s | Merek: %s | Model: %s | Tahun: %d | ID Pemilik: %d\n",
						tempArr[idxKetemu].ID,
						tempArr[idxKetemu].PlatNomor,
						tempArr[idxKetemu].Merek,
						tempArr[idxKetemu].Model,
						tempArr[idxKetemu].TahunProduksi,
						tempArr[idxKetemu].IDPemilik)
				} else {
					fmt.Println("Kendaraan dengan plat nomor tersebut tidak ditemukan.")
				}
			}

		} else if pilihan == 5 {

			// ---- PENGURUTAN KENDARAAN ----
			var (
				subPilihan int
				i          int
				j          int
				minIdx     int
				tempData   Kendaraan
				tempServis ServisRecord
			)
			fmt.Println()
			fmt.Println("--- Pengurutan Kendaraan ---")
			fmt.Println("1. Selection Sort berdasarkan Tahun Produksi")
			fmt.Println("2. Insertion Sort berdasarkan Tanggal Servis Terakhir")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				// Selection Sort berdasarkan TahunProduksi (ascending)
				i = 0
				for i < jumlahKendaraan-1 {
					minIdx = i
					j = i + 1
					for j < jumlahKendaraan {
						if daftarKendaraan[j].TahunProduksi < daftarKendaraan[minIdx].TahunProduksi {
							minIdx = j
						}
						j++
					}
					tempData = daftarKendaraan[i]
					daftarKendaraan[i] = daftarKendaraan[minIdx]
					daftarKendaraan[minIdx] = tempData
					i++
				}

				fmt.Println("Data kendaraan berhasil diurutkan berdasarkan Tahun Produksi (Selection Sort).")
				fmt.Println("ID  | Plat       | Merek       | Model       | Tahun | ID Pemilik")
				i = 0
				for i < jumlahKendaraan {
					fmt.Printf("%-4d| %-11s| %-12s| %-12s| %-6d| %d\n",
						daftarKendaraan[i].ID,
						daftarKendaraan[i].PlatNomor,
						daftarKendaraan[i].Merek,
						daftarKendaraan[i].Model,
						daftarKendaraan[i].TahunProduksi,
						daftarKendaraan[i].IDPemilik)
					i++
				}

			} else if subPilihan == 2 {
				// Insertion Sort pada daftarServis berdasarkan TanggalServis (ascending)
				// Kemudian tampilkan urutan kendaraan berdasarkan servis terakhir

				// Insertion sort daftarServis
				i = 1
				for i < jumlahServis {
					tempServis = daftarServis[i]
					j = i - 1
					for j >= 0 && daftarServis[j].TanggalServis > tempServis.TanggalServis {
						daftarServis[j+1] = daftarServis[j]
						j--
					}
					daftarServis[j+1] = tempServis
					i++
				}

				// Cari tanggal servis terakhir per kendaraan lalu tampilkan
				// Untuk tiap kendaraan, cari tanggal servis terbesar dari daftarServis
				fmt.Println("Data kendaraan berdasarkan Tanggal Servis Terakhir (Insertion Sort pada data servis):")
				fmt.Println("ID Kend | Tanggal Servis Terakhir")

				i = 0
				for i < jumlahKendaraan {
					var (
						idKend          int
						tanggalTerakhir string
						k               int
					)
					idKend = daftarKendaraan[i].ID
					tanggalTerakhir = "-"
					k = 0
					for k < jumlahServis {
						if daftarServis[k].IDKendaraan == idKend {
							tanggalTerakhir = daftarServis[k].TanggalServis
						}
						k++
					}
					fmt.Printf("%-8d| %s\n", idKend, tanggalTerakhir)
					i++
				}
			}

		} else if pilihan == 6 {

			// ---- STATISTIK ----
			var subPilihan int
			fmt.Println()
			fmt.Println("--- Statistik ---")
			fmt.Println("1. Jumlah Kendaraan Diservis per Bulan")
			fmt.Println("2. Kategori Kerusakan Paling Sering Muncul")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				// Hitung jumlah servis per bulan
				// Format tanggal: YYYY-MM-DD, bulan ada di karakter indeks 5-6
				// Kita simpan bulan unik dan hitungannya
				var (
					daftarBulan [500]string
					hitungBulan [500]int
					jumlahBulan int
					i           int
					j           int
					bulanIni    string
					ketemu      bool
				)

				jumlahBulan = 0
				i = 0
				for i < jumlahServis {
					// Ambil YYYY-MM dari tanggal (7 karakter pertama)
					bulanIni = ""
					j = 0
					for j < 7 {
						bulanIni = bulanIni + string(daftarServis[i].TanggalServis[j])
						j++
					}

					// Cek apakah bulan sudah ada
					ketemu = false
					j = 0
					for j < jumlahBulan {
						if daftarBulan[j] == bulanIni {
							hitungBulan[j]++
							ketemu = true
						}
						j++
					}
					if ketemu == false {
						daftarBulan[jumlahBulan] = bulanIni
						hitungBulan[jumlahBulan] = 1
						jumlahBulan++
					}
					i++
				}

				if jumlahBulan == 0 {
					fmt.Println("Belum ada data servis.")
				} else {
					fmt.Println("Jumlah Kendaraan Diservis per Bulan:")
					fmt.Println("Bulan       | Jumlah")
					i = 0
					for i < jumlahBulan {
						fmt.Printf("%-12s| %d\n", daftarBulan[i], hitungBulan[i])
						i++
					}
				}

			} else if subPilihan == 2 {
				// Kategori kerusakan paling sering muncul
				var (
					daftarJenis [500]string
					hitungJenis [500]int
					jumlahJenis int
					i           int
					j           int
					jenisIni    string
					ketemu      bool
					maxHitung   int
					jenisMax    string
				)

				jumlahJenis = 0
				i = 0
				for i < jumlahServis {
					jenisIni = daftarServis[i].JenisKerusakan

					ketemu = false
					j = 0
					for j < jumlahJenis {
						if daftarJenis[j] == jenisIni {
							hitungJenis[j]++
							ketemu = true
						}
						j++
					}
					if ketemu == false {
						daftarJenis[jumlahJenis] = jenisIni
						hitungJenis[jumlahJenis] = 1
						jumlahJenis++
					}
					i++
				}

				if jumlahJenis == 0 {
					fmt.Println("Belum ada data servis.")
				} else {
					fmt.Println("Kategori Kerusakan:")
					fmt.Println("Jenis Kerusakan  | Jumlah")
					maxHitung = 0
					jenisMax = ""
					i = 0
					for i < jumlahJenis {
						fmt.Printf("%-17s| %d\n", daftarJenis[i], hitungJenis[i])
						if hitungJenis[i] > maxHitung {
							maxHitung = hitungJenis[i]
							jenisMax = daftarJenis[i]
						}
						i++
					}
					fmt.Println()
					fmt.Printf("Kerusakan paling sering: %s (%d kali)\n", jenisMax, maxHitung)
				}
			}

		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
