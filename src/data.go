/*
	Module ini dimana tempat deklarasi variable publik, tipe data struct.
	Serta terdapat fungsi untuk mengisi data saat program baru saja dijalankan. 
*/
package src

// Kapasitas maksimum data (ukuran tetap)
const MAX int = 200

// Variabel global untuk menyimpan data di memory
var (
	dataPemilik [MAX]Pemilik
	dataKendaraan [MAX]Kendaraan
	dataServis [MAX]Servis

	nPemilik, nKendaraan, nServis int
	idPemilikNext, idKendaraanNext, idServisNext int
)

// Tipe data Struct untuk Pemilik kendaraan
type Pemilik struct {
	ID int
	Nama string
	Alamat string
	NoHP string
}

// Tipe data Struct data untuk Kendaraan
type Kendaraan struct {
	ID int
	PlatNomor string
	Merk string
	Model string
	TahunProduksi int
	IDPemilik int
}

// Tipe data Struct untuk Riwayat Servis
type Servis struct {
	ID int
	IDKendaraan int
	Tanggal string
	JenisKerusakan string
	Biaya float64
	Keterangan string
}

// Memasukan data awal kedalam variable array
func IsiDataAwal() {
	dataPemilik[0] = Pemilik{ID: 1, Nama: "Budi_Santoso", Alamat: "Jl._Merdeka_No.10", NoHP: "081234567890"}
	dataPemilik[1] = Pemilik{ID: 2, Nama: "Siti_Aminah", Alamat: "Jl._Sudirman_No.25", NoHP: "082345678901"}
	dataPemilik[2] = Pemilik{ID: 3, Nama: "Ahmad_Rizki", Alamat: "Jl._Pahlawan_No.7", NoHP: "083456789012"}
	nPemilik = 3
	idPemilikNext = 4

	dataKendaraan[0] = Kendaraan{ID: 1, PlatNomor: "B1234ABC", Merk: "Toyota", Model: "Avanza", TahunProduksi: 2018, IDPemilik: 1}
	dataKendaraan[1] = Kendaraan{ID: 2, PlatNomor: "D5678XYZ", Merk: "Honda", Model: "Brio", TahunProduksi: 2020, IDPemilik: 2}
	dataKendaraan[2] = Kendaraan{ID: 3, PlatNomor: "B9012DEF", Merk: "Suzuki", Model: "Ertiga", TahunProduksi: 2015, IDPemilik: 1}
	dataKendaraan[3] = Kendaraan{ID: 4, PlatNomor: "F3456GHI", Merk: "Daihatsu", Model: "Xenia", TahunProduksi: 2019, IDPemilik: 3}
	dataKendaraan[4] = Kendaraan{ID: 5, PlatNomor: "B7890JKL", Merk: "Toyota", Model: "Innova", TahunProduksi: 2022, IDPemilik: 2}
	nKendaraan = 5
	idKendaraanNext = 6

	dataServis[0] = Servis{ID: 1, IDKendaraan: 1, Tanggal: "2025-01-15", JenisKerusakan: "Ganti_Oli", Biaya: 250000, Keterangan: "Servis_rutin"}
	dataServis[1] = Servis{ID: 2, IDKendaraan: 1, Tanggal: "2025-03-20", JenisKerusakan: "Ganti_Kampas_Rem", Biaya: 450000, Keterangan: "Kampas_rem_habis"}
	dataServis[2] = Servis{ID: 3, IDKendaraan: 2, Tanggal: "2025-02-10", JenisKerusakan: "Tune_Up", Biaya: 350000, Keterangan: "Tune_up_mesin"}
	dataServis[3] = Servis{ID: 4, IDKendaraan: 3, Tanggal: "2025-01-25", JenisKerusakan: "Ganti_Oli", Biaya: 280000, Keterangan: "Servis_rutin"}
	dataServis[4] = Servis{ID: 5, IDKendaraan: 4, Tanggal: "2025-04-05", JenisKerusakan: "Ganti_Ban", Biaya: 1200000, Keterangan: "Ban_depan_aus"}
	dataServis[5] = Servis{ID: 6, IDKendaraan: 5, Tanggal: "2025-02-18", JenisKerusakan: "Ganti_Oli", Biaya: 300000, Keterangan: "Servis_rutin"}
	dataServis[6] = Servis{ID: 7, IDKendaraan: 2, Tanggal: "2025-04-22", JenisKerusakan: "Ganti_Aki", Biaya: 850000, Keterangan: "Aki_soak"}
	dataServis[7] = Servis{ID: 8, IDKendaraan: 3, Tanggal: "2025-05-10", JenisKerusakan: "Ganti_Oli", Biaya: 260000, Keterangan: "Servis_rutin"}
	nServis = 8
	idServisNext = 9
}