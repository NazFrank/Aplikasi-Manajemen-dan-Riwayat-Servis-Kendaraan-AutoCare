package statistik

import (
	"autocare/data"
	"fmt"
)

// ============================================================
// VARIABEL UNTUK STATISTIK
// ============================================================

var (
	// NamaKerusakan menyimpan jenis-jenis kerusakan yang pernah tercatat
	NamaKerusakan [data.MAKS_DATA]string
	// JumlahPerKerusakan menyimpan berapa kali tiap jenis kerusakan muncul
	JumlahPerKerusakan [data.MAKS_DATA]int
	// JumlahJenisKerusakan adalah banyaknya jenis kerusakan unik
	JumlahJenisKerusakan int
)

// ============================================================
// FUNGSI BANTU
// ============================================================

// cetakGaris mencetak garis pemisah sepanjang 42 karakter
func cetakGaris() {
	fmt.Println("  ------------------------------------------")
}

// ulangKarakter mencetak karakter sebanyak n kali
func ulangKarakter(karakter string, n int) string {
	var hasil string
	var i int
	hasil = ""
	for i = 0; i < n; i++ {
		hasil = hasil + karakter
	}
	return hasil
}

// ============================================================
// FUNGSI STATISTIK
// ============================================================

// HitungServisPerBulan menghitung jumlah servis untuk setiap bulan
// pada tahun tertentu lalu menampilkan hasilnya
func HitungServisPerBulan(tahun string) {
	// Array untuk 12 bulan (indeks 0 = Januari, dst.)
	var jumlahPerBulan [12]int
	var namaBulan [12]string
	var i int

	namaBulan[0] = "Januari"
	namaBulan[1] = "Februari"
	namaBulan[2] = "Maret"
	namaBulan[3] = "April"
	namaBulan[4] = "Mei"
	namaBulan[5] = "Juni"
	namaBulan[6] = "Juli"
	namaBulan[7] = "Agustus"
	namaBulan[8] = "September"
	namaBulan[9] = "Oktober"
	namaBulan[10] = "November"
	namaBulan[11] = "Desember"

	// Inisialisasi semua bulan dengan 0
	for i = 0; i < 12; i++ {
		jumlahPerBulan[i] = 0
	}

	// Hitung servis per bulan
	// Format tanggal: "YYYY-MM-DD"
	for i = 0; i < data.JumlahServis; i++ {
		var tgl string
		tgl = data.DaftarServis[i].TanggalServis

		// Pastikan panjang string cukup untuk diproses (minimal 7 karakter)
		if len(tgl) >= 7 {
			var tahunData string
			var bulanData string
			tahunData = tgl[0:4] // ambil 4 karakter pertama (tahun)
			bulanData = tgl[5:7] // ambil 2 karakter setelah "-" (bulan)

			if tahunData == tahun {
				// Konversi bulan dari string ke indeks array
				var indeksBulan int
				switch bulanData {
				case "01":
					indeksBulan = 0
				case "02":
					indeksBulan = 1
				case "03":
					indeksBulan = 2
				case "04":
					indeksBulan = 3
				case "05":
					indeksBulan = 4
				case "06":
					indeksBulan = 5
				case "07":
					indeksBulan = 6
				case "08":
					indeksBulan = 7
				case "09":
					indeksBulan = 8
				case "10":
					indeksBulan = 9
				case "11":
					indeksBulan = 10
				case "12":
					indeksBulan = 11
				default:
					indeksBulan = -1
				}

				if indeksBulan >= 0 {
					jumlahPerBulan[indeksBulan]++
				}
			}
		}
	}

	// Tampilkan hasil
	fmt.Printf("\n  Statistik Servis Kendaraan Tahun %s\n", tahun)
	cetakGaris()
	fmt.Printf("  %-12s | %-6s | %s\n", "Bulan", "Jumlah", "Grafik")
	cetakGaris()

	var totalServis int
	totalServis = 0

	for i = 0; i < 12; i++ {
		var grafik string
		grafik = ulangKarakter("█", jumlahPerBulan[i])
		fmt.Printf("  %-12s | %-6d | %s\n", namaBulan[i], jumlahPerBulan[i], grafik)
		totalServis = totalServis + jumlahPerBulan[i]
	}

	cetakGaris()
	fmt.Printf("  Total Servis : %d kendaraan\n", totalServis)
}

// HitungKategoriKerusakan menghitung dan menampilkan statistik
// kategori kerusakan yang paling sering muncul
func HitungKategoriKerusakan() {
	var i int
	var j int
	var ditemukan bool

	// Reset data statistik kerusakan
	JumlahJenisKerusakan = 0
	for i = 0; i < data.MAKS_DATA; i++ {
		NamaKerusakan[i] = ""
		JumlahPerKerusakan[i] = 0
	}

	// Hitung frekuensi setiap jenis kerusakan
	for i = 0; i < data.JumlahServis; i++ {
		var jenis string
		jenis = data.DaftarServis[i].JenisKerusakan
		ditemukan = false

		// Cek apakah jenis kerusakan ini sudah pernah dicatat
		for j = 0; j < JumlahJenisKerusakan; j++ {
			if NamaKerusakan[j] == jenis {
				JumlahPerKerusakan[j]++
				ditemukan = true
			}
		}

		// Jika belum ada, tambahkan jenis kerusakan baru
		if !ditemukan {
			NamaKerusakan[JumlahJenisKerusakan] = jenis
			JumlahPerKerusakan[JumlahJenisKerusakan] = 1
			JumlahJenisKerusakan++
		}
	}

	// Urutkan berdasarkan frekuensi tertinggi (bubble sort sederhana)
	for i = 0; i < JumlahJenisKerusakan-1; i++ {
		for j = 0; j < JumlahJenisKerusakan-1-i; j++ {
			if JumlahPerKerusakan[j] < JumlahPerKerusakan[j+1] {
				// Tukar jumlah
				var tempJumlah int
				tempJumlah = JumlahPerKerusakan[j]
				JumlahPerKerusakan[j] = JumlahPerKerusakan[j+1]
				JumlahPerKerusakan[j+1] = tempJumlah

				// Tukar nama
				var tempNama string
				tempNama = NamaKerusakan[j]
				NamaKerusakan[j] = NamaKerusakan[j+1]
				NamaKerusakan[j+1] = tempNama
			}
		}
	}

	// Tampilkan hasil
	fmt.Println("\n  Statistik Kategori Kerusakan (Paling Sering)")
	cetakGaris()
	fmt.Printf("  %-4s %-18s %-8s %s\n", "No", "Jenis Kerusakan", "Jumlah", "Grafik")
	cetakGaris()

	for i = 0; i < JumlahJenisKerusakan; i++ {
		var grafik string
		grafik = ulangKarakter("▓", JumlahPerKerusakan[i])
		fmt.Printf("  %-4d %-18s %-8d %s\n",
			i+1, NamaKerusakan[i], JumlahPerKerusakan[i], grafik)
	}
	cetakGaris()

	if JumlahJenisKerusakan > 0 {
		fmt.Printf("  Kerusakan Terbanyak: %s (%d kali)\n",
			NamaKerusakan[0], JumlahPerKerusakan[0])
	}
}
