package functions
import (
	"fmt"
	"strings"
)

const jumlahKarier = 6
type InfoKarier struct {
	Kode      string
	Nama      string
	Deskripsi string
}

var dataKarier [jumlahKarier]InfoKarier = [jumlahKarier]InfoKarier{
	{"R", "Realistic", "Sering menggunakan benda seperti alat-alat, mesin, serta suka memelihara hewan dan tumbuhan. Terpaku dan memiliki fokus yang baik dalam mengerjakan sesuatu yang disukai."},
	{"I", "Investigative", "Suka mencari tahu struktur dan teori terhadap suatu hal yang terstruktur alamiah dan aturan-aturan dalam masyarakat. Memiliki rasa ingin tahu yang tinggi, serta menyukai ide-ide baru."},
	{"A", "Artistic", "Suka mengekspresikan diri dalam bentuk gambar, lukisan, musik, sastra, dan drama. Lebih peka terhadap rasa dan perasaan, serta lebih mementingkan kebebasan daripada apapun."},
	{"S", "Social", "Suka membantu orang, mendidik, dan bisa menjalin hubungan yang baik dengan lingkungan sekitar. Bisa mengendalikan diri, sabar, serta baik pada orang lain."},
	{"E", "Enterprising", "Berusaha mencapai dan mewujudkan tujuan bersama, suka bernegosiasi dan berdiskusi. Mempunyai kepercayaan diri dalam bergaul dan bermasyarakat, bisa memimpin, serta bisa memberikan pengaruh kepada orang lain."},
	{"C", "Conventional", "Suka mengurusi berkas-berkas, menganalisa data, serta bekerja sesuai langkah atau urutan yang telah ditentukan. Tertib dan mematuhi aturan, serta berhati-hati dalam melakukan sesuatu."},
}

// Menu utama informasi karier
func InformasiKarier() {
	var pilihan int
	for {
		fmt.Println("\n📚 Menu Informasi Karier")
		fmt.Println("-----------------------------")
		fmt.Println("1. Tampilkan Semua Info Karier")
		fmt.Println("2. Cari Karier (Sequential Search)")
		fmt.Println("3. Urutkan Karier (A-Z / Z-A)")
		fmt.Println("4. Kembali")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			TampilInfoKarier()
		} else if pilihan == 2 {
			CariKarier()
		} else if pilihan == 3 {
			MenuUrutKarier()
		} else if pilihan == 4 {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Menampilkan semua karier
func TampilInfoKarier() {
	fmt.Println("\n📖 Daftar Informasi Karier RIASEC")
	for i := 0; i < jumlahKarier; i++ {
		fmt.Printf("\n[%s] %s\n%s\n", dataKarier[i].Kode, dataKarier[i].Nama, dataKarier[i].Deskripsi)
	}
}

// Mencari karier dengan Sequential Search
func CariKarier() {
	var cari string
	var ditemukan bool = false

	fmt.Print("Masukkan nama kategori (misal: Social): ")
	fmt.Scan(&cari)

	for i := 0; i < jumlahKarier; i++ {
		if strings.EqualFold(dataKarier[i].Nama, cari) {
			fmt.Printf("\nDitemukan:\n[%s] %s\n%s\n", dataKarier[i].Kode, dataKarier[i].Nama, dataKarier[i].Deskripsi)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("❌ Data tidak ditemukan.")
	}
}

// Menu pengurutan karier
func MenuUrutKarier() {
	var opsi int
	fmt.Println("\n🔃 Pilih metode pengurutan nama karier:")
	fmt.Println("1. A-Z (Ascending)")
	fmt.Println("2. Z-A (Descending)")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		SelectionSortKarierAZ()
		fmt.Println("✅ Diurutkan dari A ke Z:")
		TampilInfoKarier()
	} else if opsi == 2 {
		InsertionSortKarierZA()
		fmt.Println("✅ Diurutkan dari Z ke A:")
		TampilInfoKarier()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// Mengurutkan dataKarier berdasarkan nama (A-Z)
func SelectionSortKarierAZ() {
	for i := 0; i < jumlahKarier-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahKarier; j++ {
			if strings.Compare(dataKarier[j].Nama, dataKarier[minIdx].Nama) < 0 {
				minIdx = j
			}
		}
		if minIdx != i {
			dataKarier[i], dataKarier[minIdx] = dataKarier[minIdx], dataKarier[i]
		}
	}
}

// Mengurutkan dataKarier berdasarkan nama (Z-A)
func InsertionSortKarierZA() {
	for i := 1; i < jumlahKarier; i++ {
		temp := dataKarier[i]
		j := i - 1
		for j >= 0 && strings.Compare(dataKarier[j].Nama, temp.Nama) < 0 {
			dataKarier[j+1] = dataKarier[j]
			j--
		}
		dataKarier[j+1] = temp
	}
}
