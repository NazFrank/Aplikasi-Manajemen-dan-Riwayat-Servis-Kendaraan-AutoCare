# AutoCare - Aplikasi Manajemen & Riwayat Servis Kendaraan

## Struktur Folder (Modular)

```
autocare/
├── main.go                  → Program utama, titik masuk aplikasi
├── go.mod                   → File modul Go
│
├── data/
│   └── data.go              → Struct, variabel global, fungsi CRUD
│
├── algoritma/
│   ├── pencarian.go         → Sequential Search & Binary Search
│   └── pengurutan.go        → Selection Sort & Insertion Sort
│
├── statistik/
│   └── statistik.go         → Hitung servis per bulan & kategori kerusakan
│
├── helper/
│   └── helper.go            → Fungsi tampilan (display) & input
│
└── menu/
    └── menu.go              → Semua handler menu interaktif
```

## Fitur Aplikasi

| No | Fitur | Keterangan |
|----|-------|------------|
| 1  | Manajemen Kendaraan | Tambah, ubah, hapus data kendaraan |
| 2  | Manajemen Pemilik | Tambah, ubah data pemilik kendaraan |
| 3  | Riwayat Servis | Catat dan lihat riwayat servis per kendaraan |
| 4  | Pencarian | Sequential Search & Binary Search berdasarkan plat nomor |
| 5  | Pengurutan | Selection Sort (tahun) & Insertion Sort (tanggal servis) |
| 6  | Statistik | Grafik servis per bulan & kategori kerusakan terbanyak |

## Algoritma yang Digunakan

### Sequential Search
- Memeriksa setiap elemen dari awal hingga akhir
- Tidak memerlukan data terurut
- Cocok untuk data kecil atau data yang tidak terurut

### Binary Search
- Membagi area pencarian menjadi dua bagian setiap iterasi
- **Memerlukan data yang sudah terurut**
- Lebih cepat dari Sequential Search untuk data besar

### Selection Sort
- Cari elemen terkecil → tempatkan di depan → ulangi untuk sisa
- Digunakan untuk mengurutkan berdasarkan **tahun produksi**

### Insertion Sort
- Ambil satu elemen → sisipkan ke posisi tepat di bagian terurut
- Mirip cara mengurutkan kartu remi
- Digunakan untuk mengurutkan berdasarkan **tanggal servis terakhir**

