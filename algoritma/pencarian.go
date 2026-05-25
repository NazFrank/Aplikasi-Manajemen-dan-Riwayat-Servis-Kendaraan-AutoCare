package algoritma

import (
	"autocare/data"
)

// ============================================================
// ALGORITMA PENCARIAN
// ============================================================

// samaDengan membandingkan dua string tanpa membedakan huruf besar/kecil
// Dilakukan manual karena tidak boleh menggunakan package strings
func samaDengan(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var i int
	for i = 0; i < len(a); i++ {
		var ca byte
		var cb byte
		ca = a[i]
		cb = b[i]
		// Ubah huruf kecil ke huruf besar jika perlu
		if ca >= 'a' && ca <= 'z' {
			ca = ca - 32
		}
		if cb >= 'a' && cb <= 'z' {
			cb = cb - 32
		}
		if ca != cb {
			return false
		}
	}
	return true
}

// lebihBesar membandingkan dua string secara leksikografis (case-insensitive)
// mengembalikan true jika a > b
func lebihBesar(a string, b string) bool {
	var i int
	var panjang int
	panjang = len(a)
	if len(b) < panjang {
		panjang = len(b)
	}
	for i = 0; i < panjang; i++ {
		var ca byte
		var cb byte
		ca = a[i]
		cb = b[i]
		if ca >= 'a' && ca <= 'z' {
			ca = ca - 32
		}
		if cb >= 'a' && cb <= 'z' {
			cb = cb - 32
		}
		if ca > cb {
			return true
		}
		if ca < cb {
			return false
		}
	}
	return len(a) > len(b)
}

// CariSequential mencari kendaraan berdasarkan plat nomor
// menggunakan Sequential Search (pencarian satu per satu dari awal)
// Mengembalikan indeks posisi kendaraan, atau -1 jika tidak ditemukan
func CariSequential(platNomor string) int {
	var i int

	// Cek setiap kendaraan satu per satu dari indeks 0
	for i = 0; i < data.JumlahKendaraan; i++ {
		if samaDengan(data.DaftarKendaraan[i].PlatNomor, platNomor) {
			return i // ditemukan, kembalikan posisinya
		}
	}

	return -1 // tidak ditemukan
}

// KendaraanTerurut adalah array kendaraan yang sudah diurutkan
// digunakan sebagai persiapan Binary Search
var KendaraanTerurut [data.MAKS_DATA]data.Kendaraan
var JumlahKendaraanTerurut int

// SalinDanUrutkanUntukBinarySearch menyalin dan mengurutkan data
// kendaraan berdasarkan plat nomor (ascending) agar bisa digunakan
// oleh Binary Search (Binary Search butuh data yang sudah terurut)
func SalinDanUrutkanUntukBinarySearch() {
	var i int
	var j int
	var temp data.Kendaraan

	// Salin semua data kendaraan ke array sementara
	JumlahKendaraanTerurut = data.JumlahKendaraan
	for i = 0; i < data.JumlahKendaraan; i++ {
		KendaraanTerurut[i] = data.DaftarKendaraan[i]
	}

	// Urutkan dengan Insertion Sort berdasarkan plat nomor
	for i = 1; i < JumlahKendaraanTerurut; i++ {
		temp = KendaraanTerurut[i]
		j = i - 1
		for j >= 0 && lebihBesar(KendaraanTerurut[j].PlatNomor, temp.PlatNomor) {
			KendaraanTerurut[j+1] = KendaraanTerurut[j]
			j--
		}
		KendaraanTerurut[j+1] = temp
	}
}

// CariBinary mencari kendaraan berdasarkan plat nomor
// menggunakan Binary Search (pencarian dengan membagi dua area pencarian)
// Data HARUS sudah diurutkan sebelum memanggil fungsi ini
// Mengembalikan indeks di dalam KendaraanTerurut, atau -1 jika tidak ditemukan
func CariBinary(platNomor string) int {
	var kiri int
	var kanan int
	var tengah int

	kiri = 0
	kanan = JumlahKendaraanTerurut - 1

	for kiri <= kanan {
		// Hitung posisi tengah
		tengah = (kiri + kanan) / 2

		if samaDengan(KendaraanTerurut[tengah].PlatNomor, platNomor) {
			return tengah // ditemukan!
		} else if lebihBesar(platNomor, KendaraanTerurut[tengah].PlatNomor) {
			// Cari di bagian kanan (nilai lebih besar)
			kiri = tengah + 1
		} else {
			// Cari di bagian kiri (nilai lebih kecil)
			kanan = tengah - 1
		}
	}

	return -1 // tidak ditemukan
}
