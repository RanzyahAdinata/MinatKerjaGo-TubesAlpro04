# 🧠 MinatKerjaGo

**MinatKerjaGo** adalah aplikasi terminal berbasis Go (Golang) yang dirancang untuk membantu pengguna mengidentifikasi minat dan potensi karier berdasarkan metode **RIASEC** (Realistic, Investigative, Artistic, Social, Enterprising, Conventional). Aplikasi ini menyediakan tes minat & keahlian, menyimpan hasil tes secara lokal, serta menampilkan informasi karier yang relevan.

---

## ✨ Fitur Utama

- 📋 **Tes Minat & Keahlian (RIASEC)**  
  Pengguna menjawab pernyataan untuk mengukur minat dalam 6 kategori kepribadian karier.

- 💾 **Penyimpanan Hasil Tes Otomatis**  
  Hasil tes disimpan dalam file `.txt` dengan format `hasiltes1.txt`, `hasiltes2.txt`, dst. di folder `memory/`.

- 🔍 **Riwayat Tes**  
  - Melihat daftar semua hasil tes  
  - Menampilkan detail hasil tes tertentu  
  - Menghapus satu atau semua riwayat tes  
  - Mencetak hasil tes ke dalam format PDF di folder `memory/pdf/`

- 📚 **Informasi Karier Berdasarkan RIASEC**  
  Menyediakan penjelasan singkat dan rekomendasi karier sesuai tipe kepribadian.

---

## 🗂 Struktur Folder
  minatkerjago-kelompok04/
│
├── main.go # Entry point aplikasi
├── tesminat.go # Logika tes minat & keahlian
├── riwayattes.go # Manajemen riwayat hasil tes
├── utils.go # Fungsi utilitas pendukung
├── soal/ # Folder berisi soal-soal tes RIASEC
├── memory/ # Penyimpanan hasil tes (.txt)
└── memory/pdf/ # Penyimpanan hasil tes dalam format PDF


---

## 🚀 Cara Menjalankan

1. Clone repositori:  
   `git clone https://github.com/nama-user/minatkerjago-kelompok04.git`  
   `cd minatkerjago-kelompok04`

2. Pastikan Go sudah terpasang:  
   `go version`

3. Jalankan aplikasi:  
   `go run main.go`

4. Ikuti instruksi untuk melakukan tes dan mengelola hasil.

---

## 👥 Tim Pengembang

- Ranzyah Adinata Aldo  
- Wafiq Aditya Wiyono

---

## 📄 Lisensi

Proyek ini dibuat untuk keperluan akademik. Bebas digunakan untuk pembelajaran dengan menyertakan atribusi kepada pengembang.

---

Terima kasih telah menggunakan **MinatKerjaGo**! Semoga aplikasi ini membantu kamu menemukan karier yang tepat. 🚀

