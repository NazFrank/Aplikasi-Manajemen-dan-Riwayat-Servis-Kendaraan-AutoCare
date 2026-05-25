package algoritma

import "autocare/data"

// ============================================================
// ALGORITMA PENGURUTAN
// ============================================================

// HasilUrutan menyimpan salinan data kendaraan yang sudah diurutkan
// agar data asli tidak berubah
var HasilUrutan [data.MAKS_DATA]data.Kendaraan
var JumlahHasilUrutan int

// SelectionSortTahun mengurutkan kendaraan berdasarkan tahun produksi
// menggunakan algoritma Selection Sort (ascending = dari terkecil)
// Cara kerja: cari nilai terkecil, taruh di depan, ulangi untuk sisa
func SelectionSortTahun() {
	var i int
	var j int
	var indeksMin int
	var temp data.Kendaraan

	// Salin data asli agar tidak diubah
	JumlahHasilUrutan = data.JumlahKendaraan
	for i = 0; i < data.JumlahKendaraan; i++ {
		HasilUrutan[i] = data.DaftarKendaraan[i]
	}

	// Proses Selection Sort
	for i = 0; i < JumlahHasilUrutan - 1; i++ {
		// Anggap elemen pertama pada area yang belum terurut sebagai minimum
		indeksMin = i

		// Cari nilai minimum di sisa array
		for j = i + 1; j < JumlahHasilUrutan; j++ {
			if HasilUrutan[j].TahunProduksi < HasilUrutan[indeksMin].TahunProduksi {
				indeksMin = j
			}
		}

		// Tukar elemen saat ini dengan elemen minimum yang ditemukan
		if indeksMin != i {
			temp = HasilUrutan[i]
			HasilUrutan[i] = HasilUrutan[indeksMin]
			HasilUrutan[indeksMin] = temp
		}
	}
}

// InsertionSortTanggalServis mengurutkan kendaraan berdasarkan
// tanggal servis terakhir menggunakan algoritma Insertion Sort (ascending)
// Cara kerja: ambil satu elemen, sisipkan ke posisi yang tepat di bagian
// yang sudah terurut sebelumnya (mirip seperti mengurutkan kartu remi)
func InsertionSortTanggalServis() {
	var i int
	var j int
	var temp data.Kendaraan

	// Salin data asli agar tidak diubah
	JumlahHasilUrutan = data.JumlahKendaraan
	for i = 0; i < data.JumlahKendaraan; i++ {
		HasilUrutan[i] = data.DaftarKendaraan[i]
	}

	// Proses Insertion Sort
	for i = 1; i < JumlahHasilUrutan; i++ {
		// Simpan elemen yang sedang diperiksa
		temp = HasilUrutan[i]
		j = i - 1

		// Geser elemen-elemen yang lebih besar ke kanan
		// Perbandingan string tanggal "YYYY-MM-DD" bisa langsung
		// karena format ini bisa diurutkan secara leksikografis
		for j >= 0 && HasilUrutan[j].TanggalServisTerakhir > temp.TanggalServisTerakhir {
			HasilUrutan[j + 1] = HasilUrutan[j]
			j--
		}

		// Sisipkan elemen ke posisi yang tepat
		HasilUrutan[j + 1] = temp
	}
}
