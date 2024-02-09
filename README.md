# enigma-laundry-godb-sbmt

Challenge Golang Pembuatan Aplikasi Laundry Enigma

Enigma Laundry Console App

- Deskripsi
  Aplikasi konsol Enigma Laundry memungkinkan Anda mencatat transaksi di toko laundry, melibatkan pengelolaan pelanggan, layanan, dan transaksi laundry.

- Persyaratan

1. Go (minimal versi 1.11)
2. PostgreSQL

- Langkah-langkah Menjalankan Aplikasi Clone repositori ke dalam direktori lokal Anda:
  ```bash
  git clone https://github.com/saepudinasep/Aplikasi-Laundry.git
  cd username/repositori
  ```
  Atau
  ```bash
  git clone https://git.enigmacamp.com/enigma-20/asep-saepudin/challenge-godb.git
  cd username/repositori
  ```

1. Inisialisasi Go Module:

   go mod init nama-proyek-anda

2. Install Dependensi:

   go get

3. Inisialisasi Database:

Buat database PostgreSQL dengan nama enigma_laundry.
Sesuaikan konfigurasi koneksi database dalam file main.go pada fungsi initDB.

4. Jalankan Aplikasi:

   go run main.go

- Berikut ini adalah tampilan aplikasi ketika program dijalankan:

Menu Aplikasi
Aplikasi ini menyediakan beberapa menu untuk mengelola data pelanggan, layanan, dan transaksi laundry.

Menu Utama

1. Master Customer
2. Master Layanan
3. Make Orders
4. Exit

- Contoh Penggunaan

1.  Master Customer:
    Menampilkan menu customer yaitu;

    1.  Add Customer (Tambah Pelanggan)
        Memasukkan informasi baru tentang pelanggan ke dalam database.
    2.  Update Customer
        Mengubah informasi dari suatu pelanggan yang sudah terdaftar.
    3.  Delete Customer
        Menghapus suatu pelanggan yang sudah terdaftar dari database.
    4.  View All Customer
        Menampilkan semua data pelanggan yang telah terdaftar.
    5.  View Customer By Id
        Menampilkan detail dari suatu pelanggan berdasarkan idnya.
    6.  Keluar
        Masukkan angka sesuai pilihan Anda : 0

2.  Master Layanan:
    Menampilkan menu layanan yaitu;

    1.  Add Layanan (Tambah Layanan)
        Memasukkan jenis layanan baru ke dalam database.
    2.  Update Layanan
        Merubah jenis layanan yang sudah terdaftar.
    3.  Delete Layanan
        Menghapus jenis layanan yang sudah terdaftar dari database.
    4.  View All Layanan
        Menampilkan semua jenis layanan yang tersimpan di database.
    5.  View Layanan By Id
        Menampilkan detail dari suatu jenis layanan berdasarkan idnya.
    6.  Keluar
        Masukkan angka sesuai pilihan Anda : 0

3.  Make Orders:
    Menampilkan menu orders yaitu;

    1.  Make Orders
        Melakukan order dengan memasukkan kode customer dan kode layanan yang akan dipesan.
    2.  Take Orders
        Ambil order yang belum selesai dilakukan oleh customer.

4.  Keluar:
    Keluar dari aplikasi.
    Masukkan angka sesuai pilihan Anda : 0

Catatan
Pastikan untuk mengganti konfigurasi database sesuai dengan kebutuhan Anda.
Gunakan Go Modules untuk manajemen dependensi.

Dengan dokumentasi ini, pengguna dapat dengan mudah memahami langkah-langkah untuk menjalankan aplikasi dan cara menggunakan setiap menu yang disediakan. Anda dapat menyimpannya dalam format `README.md` di repositori GitHub atau mengekspornya ke format Word atau PDF jika diperlukan.
