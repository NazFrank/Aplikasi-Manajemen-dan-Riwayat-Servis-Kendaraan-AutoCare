package data

import "fmt"

// ============================================================
// STRUKTUR DATA (STRUCT)
// ============================================================

// Pemilik menyimpan data pemilik kendaraan
type Pemilik struct {
	ID      int
	Nama    string
	NoTelp  string
	Alamat  string
}

// Kendaraan menyimpan data kendaraan beserta pemiliknya
type Kendaraan struct {
	ID              int
	PlatNomor       string
	Merek           string
	Model           string
	TahunProduksi   int
	IDPemilik       int
	TanggalServisTerakhir string // format: "YYYY-MM-DD"
}

// RiwayatServis menyimpan catatan setiap kali kendaraan diservis
type RiwayatServis struct {
	ID             int
	IDKendaraan    int
	TanggalServis  string // format: "YYYY-MM-DD"
	JenisKerusakan string
	Deskripsi      string
	BiayaServis    int
}

// ============================================================
// KONSTANTA
// ============================================================

const MAKS_DATA int = 100

// ============================================================
// VARIABEL GLOBAL (PENYIMPANAN DATA)
// ============================================================

var (
	DaftarPemilik    [MAKS_DATA]Pemilik
	JumlahPemilik    int

	DaftarKendaraan  [MAKS_DATA]Kendaraan
	JumlahKendaraan  int

	DaftarServis     [MAKS_DATA]RiwayatServis
	JumlahServis     int

	IDPemilikBerikutnya    int
	IDKendaraanBerikutnya  int
	IDServisBerikutnya     int
)

// ============================================================
// FUNGSI INISIALISASI DATA CONTOH
// ============================================================

// InisialisasiDataContoh mengisi data awal supaya aplikasi
// tidak kosong saat pertama kali dijalankan
func InisialisasiDataContoh() {
	IDPemilikBerikutnya = 1
	IDKendaraanBerikutnya = 1
	IDServisBerikutnya = 1

	// --- Data Pemilik ---
	TambahPemilik("Budi Santoso", "08123456789", "Jl. Melati No. 5, Solo")
	TambahPemilik("Siti Rahayu", "08234567890", "Jl. Anggrek No. 12, Solo")
	TambahPemilik("Agus Wijaya", "08345678901", "Jl. Mawar No. 3, Sukoharjo")

	// --- Data Kendaraan ---
	TambahKendaraan("AB 1234 CD", "Toyota", "Avanza", 2020, 1, "2024-11-10")
	TambahKendaraan("AB 5678 EF", "Honda", "Civic", 2019, 1, "2024-12-05")
	TambahKendaraan("AB 9012 GH", "Yamaha", "NMAX", 2022, 2, "2025-01-20")
	TambahKendaraan("AB 3456 IJ", "Suzuki", "Ertiga", 2018, 3, "2024-10-30")
	TambahKendaraan("AB 7890 KL", "Honda", "Beat", 2021, 2, "2025-02-14")

	// --- Data Riwayat Servis ---
	TambahServis(1, "2024-08-10", "Mesin", "Ganti oli mesin dan filter", 250000)
	TambahServis(1, "2024-11-10", "Rem", "Ganti kampas rem depan dan belakang", 350000)
	TambahServis(2, "2024-09-15", "Kelistrikan", "Ganti aki dan sekring", 600000)
	TambahServis(2, "2024-12-05", "Mesin", "Tune up mesin", 400000)
	TambahServis(3, "2025-01-20", "Mesin", "Ganti oli mesin", 150000)
	TambahServis(4, "2024-07-22", "Kaki-kaki", "Ganti shock absorber", 800000)
	TambahServis(4, "2024-10-30", "Rem", "Ganti minyak rem", 100000)
	TambahServis(5, "2025-02-14", "Mesin", "Ganti oli dan busi", 200000)

	fmt.Println("[OK] Data contoh berhasil dimuat.")
	fmt.Println()
}

// ============================================================
// FUNGSI CRUD PEMILIK
// ============================================================

// TambahPemilik menambahkan data pemilik baru ke dalam array
func TambahPemilik(nama string, noTelp string, alamat string) bool {
	if JumlahPemilik >= MAKS_DATA {
		return false
	}

	var pemilikBaru Pemilik
	pemilikBaru.ID = IDPemilikBerikutnya
	pemilikBaru.Nama = nama
	pemilikBaru.NoTelp = noTelp
	pemilikBaru.Alamat = alamat

	DaftarPemilik[JumlahPemilik] = pemilikBaru
	JumlahPemilik++
	IDPemilikBerikutnya++

	return true
}

// UbahPemilik mengubah data pemilik berdasarkan ID
func UbahPemilik(id int, nama string, noTelp string, alamat string) bool {
	var i int
	for i = 0; i < JumlahPemilik; i++ {
		if DaftarPemilik[i].ID == id {
			DaftarPemilik[i].Nama = nama
			DaftarPemilik[i].NoTelp = noTelp
			DaftarPemilik[i].Alamat = alamat
			return true
		}
	}
	return false
}

// HapusPemilik menghapus data pemilik berdasarkan ID
// dengan cara menggeser elemen array ke kiri
func HapusPemilik(id int) bool {
	var i int
	var posisi int
	posisi = -1

	for i = 0; i < JumlahPemilik; i++ {
		if DaftarPemilik[i].ID == id {
			posisi = i
		}
	}

	if posisi == -1 {
		return false
	}

	for i = posisi; i < JumlahPemilik - 1; i++ {
		DaftarPemilik[i] = DaftarPemilik[i + 1]
	}
	JumlahPemilik--
	return true
}

// CariPemilikByID mencari data pemilik berdasarkan ID
func CariPemilikByID(id int) (Pemilik, bool) {
	var i int
	var pemilikKosong Pemilik
	for i = 0; i < JumlahPemilik; i++ {
		if DaftarPemilik[i].ID == id {
			return DaftarPemilik[i], true
		}
	}
	return pemilikKosong, false
}

// ============================================================
// FUNGSI CRUD KENDARAAN
// ============================================================

// TambahKendaraan menambahkan data kendaraan baru
func TambahKendaraan(platNomor string, merek string, model string,
	tahun int, idPemilik int, tglServis string) bool {
	if JumlahKendaraan >= MAKS_DATA {
		return false
	}

	var kendaraanBaru Kendaraan
	kendaraanBaru.ID = IDKendaraanBerikutnya
	kendaraanBaru.PlatNomor = platNomor
	kendaraanBaru.Merek = merek
	kendaraanBaru.Model = model
	kendaraanBaru.TahunProduksi = tahun
	kendaraanBaru.IDPemilik = idPemilik
	kendaraanBaru.TanggalServisTerakhir = tglServis

	DaftarKendaraan[JumlahKendaraan] = kendaraanBaru
	JumlahKendaraan++
	IDKendaraanBerikutnya++

	return true
}

// UbahKendaraan mengubah data kendaraan berdasarkan ID
func UbahKendaraan(id int, platNomor string, merek string, model string,
	tahun int, idPemilik int) bool {
	var i int
	for i = 0; i < JumlahKendaraan; i++ {
		if DaftarKendaraan[i].ID == id {
			DaftarKendaraan[i].PlatNomor = platNomor
			DaftarKendaraan[i].Merek = merek
			DaftarKendaraan[i].Model = model
			DaftarKendaraan[i].TahunProduksi = tahun
			DaftarKendaraan[i].IDPemilik = idPemilik
			return true
		}
	}
	return false
}

// HapusKendaraan menghapus data kendaraan berdasarkan ID
func HapusKendaraan(id int) bool {
	var i int
	var posisi int
	posisi = -1

	for i = 0; i < JumlahKendaraan; i++ {
		if DaftarKendaraan[i].ID == id {
			posisi = i
		}
	}

	if posisi == -1 {
		return false
	}

	for i = posisi; i < JumlahKendaraan - 1; i++ {
		DaftarKendaraan[i] = DaftarKendaraan[i + 1]
	}
	JumlahKendaraan--
	return true
}

// ============================================================
// FUNGSI CRUD RIWAYAT SERVIS
// ============================================================

// TambahServis menambahkan catatan servis baru
func TambahServis(idKendaraan int, tanggal string,
	jenisKerusakan string, deskripsi string, biaya int) bool {
	if JumlahServis >= MAKS_DATA {
		return false
	}

	var servisBaru RiwayatServis
	servisBaru.ID = IDServisBerikutnya
	servisBaru.IDKendaraan = idKendaraan
	servisBaru.TanggalServis = tanggal
	servisBaru.JenisKerusakan = jenisKerusakan
	servisBaru.Deskripsi = deskripsi
	servisBaru.BiayaServis = biaya

	DaftarServis[JumlahServis] = servisBaru
	JumlahServis++
	IDServisBerikutnya++

	// Perbarui tanggal servis terakhir pada kendaraan terkait
	var i int
	for i = 0; i < JumlahKendaraan; i++ {
		if DaftarKendaraan[i].ID == idKendaraan {
			DaftarKendaraan[i].TanggalServisTerakhir = tanggal
		}
	}

	return true
}

// HapusServis menghapus data servis berdasarkan ID
func HapusServis(id int) bool {
	var i int
	var posisi int
	posisi = -1

	for i = 0; i < JumlahServis; i++ {
		if DaftarServis[i].ID == id {
			posisi = i
		}
	}

	if posisi == -1 {
		return false
	}

	for i = posisi; i < JumlahServis - 1; i++ {
		DaftarServis[i] = DaftarServis[i + 1]
	}
	JumlahServis--
	return true
}
