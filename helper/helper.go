package helper

import (
	"autocare/data"
	"fmt"
)

// ============================================================
// FUNGSI TAMPILAN (DISPLAY)
// ============================================================

// TampilkanGaris menampilkan garis pemisah
func TampilkanGaris() {
	fmt.Println("============================================================")
}

// TampilkanGarisTipis menampilkan garis tipis
func TampilkanGarisTipis() {
	fmt.Println("------------------------------------------------------------")
}

// TampilkanHeaderJudul menampilkan judul bagian
func TampilkanHeaderJudul(judul string) {
	fmt.Println()
	TampilkanGaris()
	fmt.Printf("  %s\n", judul)
	TampilkanGaris()
}

// TampilkanDetailKendaraan menampilkan satu kendaraan dengan lengkap
func TampilkanDetailKendaraan(k data.Kendaraan) {
	var pemilik data.Pemilik
	var ada bool
	pemilik, ada = data.CariPemilikByID(k.IDPemilik)

	fmt.Printf("  ID Kendaraan   : %d\n", k.ID)
	fmt.Printf("  Plat Nomor     : %s\n", k.PlatNomor)
	fmt.Printf("  Merek & Model  : %s %s\n", k.Merek, k.Model)
	fmt.Printf("  Tahun Produksi : %d\n", k.TahunProduksi)
	fmt.Printf("  Servis Terakhir: %s\n", k.TanggalServisTerakhir)
	if ada {
		fmt.Printf("  Pemilik        : %s (%s)\n", pemilik.Nama, pemilik.NoTelp)
	} else {
		fmt.Printf("  Pemilik        : (Tidak Ditemukan)\n")
	}
}

// TampilkanTabelKendaraan menampilkan daftar kendaraan dalam format tabel
func TampilkanTabelKendaraan(daftar [data.MAKS_DATA]data.Kendaraan, jumlah int) {
	if jumlah == 0 {
		fmt.Println("  (Belum ada data kendaraan)")
		return
	}

	fmt.Printf("  %-4s %-14s %-10s %-12s %-6s %-12s\n",
		"ID", "Plat Nomor", "Merek", "Model", "Tahun", "Servis Terakhir")
	TampilkanGarisTipis()

	var i int
	for i = 0; i < jumlah; i++ {
		fmt.Printf("  %-4d %-14s %-10s %-12s %-6d %-12s\n",
			daftar[i].ID,
			daftar[i].PlatNomor,
			daftar[i].Merek,
			daftar[i].Model,
			daftar[i].TahunProduksi,
			daftar[i].TanggalServisTerakhir)
	}
}

// TampilkanTabelPemilik menampilkan semua data pemilik
func TampilkanTabelPemilik() {
	if data.JumlahPemilik == 0 {
		fmt.Println("  (Belum ada data pemilik)")
		return
	}

	fmt.Printf("  %-4s %-20s %-15s %-20s\n", "ID", "Nama", "No. Telp", "Alamat")
	TampilkanGarisTipis()

	var i int
	for i = 0; i < data.JumlahPemilik; i++ {
		fmt.Printf("  %-4d %-20s %-15s %-20s\n",
			data.DaftarPemilik[i].ID,
			data.DaftarPemilik[i].Nama,
			data.DaftarPemilik[i].NoTelp,
			data.DaftarPemilik[i].Alamat)
	}
}

// TampilkanRiwayatServisKendaraan menampilkan semua servis untuk satu kendaraan
func TampilkanRiwayatServisKendaraan(idKendaraan int) {
	var i int
	var ada bool
	ada = false

	fmt.Printf("  %-4s %-12s %-15s %-25s %-12s\n",
		"ID", "Tanggal", "Jenis Kerusakan", "Deskripsi", "Biaya (Rp)")
	TampilkanGarisTipis()

	for i = 0; i < data.JumlahServis; i++ {
		if data.DaftarServis[i].IDKendaraan == idKendaraan {
			fmt.Printf("  %-4d %-12s %-15s %-25s %-12d\n",
				data.DaftarServis[i].ID,
				data.DaftarServis[i].TanggalServis,
				data.DaftarServis[i].JenisKerusakan,
				data.DaftarServis[i].Deskripsi,
				data.DaftarServis[i].BiayaServis)
			ada = true
		}
	}

	if !ada {
		fmt.Println("  (Belum ada riwayat servis untuk kendaraan ini)")
	}
}

// ============================================================
// FUNGSI INPUT
// ============================================================

// BacaString membaca input string dari pengguna menggunakan fmt.Scan
// Catatan: tidak bisa membaca input yang mengandung spasi
func BacaString(prompt string) string {
	var nilai string
	fmt.Print(prompt)
	fmt.Scan(&nilai)
	return nilai
}

// BacaInt membaca input integer dari pengguna
func BacaInt(prompt string) int {
	var nilai int
	fmt.Print(prompt)
	fmt.Scan(&nilai)
	return nilai
}
