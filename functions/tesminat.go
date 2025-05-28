package functions
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Soal struct {
	Pertanyaan string
	Kategori   string
}

var daftarSoal = []Soal{
	// Realistic
	{"Saya senang berada di luar ruangan untuk bekerja atau berkegiatan.", "R"},
	{"Saya tertarik menjadi montir atau teknisi kendaraan.", "R"},
	{"Saya suka menggunakan alat-alat seperti bor, palu, atau obeng.", "R"},
	{"Saya senang dengan kegiatan berbasis olahraga atau fisik.", "R"},
	{"Saya menikmati membuat atau merakit barang dengan tangan sendiri.", "R"},
	{"Saya senang belajar hal-hal teknis dan mesin.", "R"},
	{"Saya tertarik bekerja di bidang pertanian atau kehutanan.", "R"},
	{"Saya suka menggunakan kendaraan besar seperti traktor atau forklift.", "R"},
	{"Saya senang mempelajari cara kerja peralatan berat.", "R"},
	{"Saya menikmati belajar di bengkel atau lab praktik.", "R"},
	{"Saya tertarik pada pekerjaan seperti tukang kayu atau tukang las.", "R"},
	{"Saya nyaman bekerja dengan peralatan keras dan logam.", "R"},
	{"Saya lebih suka praktik langsung daripada membaca teori.", "R"},
	{"Saya merasa puas setelah menyelesaikan proyek fisik atau mekanis.", "R"},
	{"Saya suka membongkar dan merakit perangkat keras komputer.", "R"},
	{"Saya tertarik dengan pekerjaan di bidang konstruksi atau bangunan.", "R"},

	// Investigative
	{"Saya senang membaca artikel atau buku ilmiah.", "I"},
	{"Saya tertarik pada riset dan penelitian.", "I"},
	{"Saya suka menyelesaikan teka-teki logika atau soal matematika sulit.", "I"},
	{"Saya merasa puas menemukan solusi dari masalah yang kompleks.", "I"},
	{"Saya suka menganalisis data dan membuat kesimpulan.", "I"},
	{"Saya tertarik mempelajari biologi, fisika, atau kimia.", "I"},
	{"Saya ingin bekerja sebagai ilmuwan atau analis.", "I"},
	{"Saya senang mempelajari struktur atau sistem alam.", "I"},
	{"Saya penasaran dengan cara kerja sesuatu secara mendalam.", "I"},
	{"Saya suka menggunakan komputer untuk memodelkan masalah.", "I"},
	{"Saya senang meneliti penyakit dan cara penanganannya.", "I"},
	{"Saya merasa nyaman bekerja sendiri untuk memahami sesuatu.", "I"},
	{"Saya tertarik pada bidang kedokteran atau laboratorium.", "I"},
	{"Saya suka menggunakan mikroskop, alat ukur, atau perangkat lab lainnya.", "I"},
	{"Saya suka meneliti statistik untuk menjawab pertanyaan tertentu.", "I"},
	{"Saya tertarik pada kecerdasan buatan, data science, atau teknologi analitik.", "I"},

	// Artistic
	{"Saya menikmati menulis cerita, puisi, atau lagu.", "A"},
	{"Saya tertarik pada seni musik, teater, atau tari.", "A"},
	{"Saya sering punya ide unik atau orisinal.", "A"},
	{"Saya suka berekspresi melalui karya seni.", "A"},
	{"Saya tertarik belajar desain grafis atau animasi.", "A"},
	{"Saya suka menonton atau membuat film pendek.", "A"},
	{"Saya merasa nyaman menampilkan karya saya di depan orang lain.", "A"},
	{"Saya senang bereksperimen dengan gaya atau warna.", "A"},
	{"Saya ingin bekerja sebagai desainer, penulis, atau seniman.", "A"},
	{"Saya tertarik membuat konten digital yang kreatif.", "A"},
	{"Saya sering menggambar di waktu luang.", "A"},
	{"Saya merasa terinspirasi oleh musik atau visual.", "A"},
	{"Saya suka mengedit foto, video, atau audio.", "A"},
	{"Saya senang dengan kebebasan berekspresi.", "A"},
	{"Saya lebih suka aktivitas tanpa aturan kaku.", "A"},
	{"Saya suka mengunjungi galeri seni atau pertunjukan budaya.", "A"},
	{"Saya tertarik membuat produk handmade atau kerajinan.", "A"},

	// Social
	{"Saya senang mengajar atau membimbing orang lain.", "S"},
	{"Saya merasa puas saat melihat orang lain berhasil karena bantuan saya.", "S"},
	{"Saya tertarik bekerja sebagai guru, konselor, atau perawat.", "S"},
	{"Saya suka berdiskusi dan bertukar pikiran.", "S"},
	{"Saya nyaman berbicara dengan banyak orang.", "S"},
	{"Saya senang ikut dalam kegiatan sosial atau relawan.", "S"},
	{"Saya tertarik mendampingi orang yang membutuhkan.", "S"},
	{"Saya mudah berempati terhadap perasaan orang lain.", "S"},
	{"Saya suka membantu orang belajar hal baru.", "S"},
	{"Saya lebih suka kerja tim daripada kerja sendiri.", "S"},
	{"Saya suka mendengarkan cerita dan memberikan masukan.", "S"},
	{"Saya ingin berkontribusi dalam pengembangan komunitas.", "S"},
	{"Saya tertarik bidang psikologi atau sosiologi.", "S"},
	{"Saya mudah membuat orang merasa nyaman.", "S"},
	{"Saya senang bekerja di lingkungan kolaboratif.", "S"},
	{"Saya tertarik membantu anak-anak atau lansia.", "S"},
	{"Saya ingin menjadi pendidik atau aktivis sosial.", "S"},

	// Enterprising
	{"Saya suka menjadi pemimpin dalam proyek kelompok.", "E"},
	{"Saya menikmati meyakinkan orang lain terhadap ide saya.", "E"},
	{"Saya suka berkompetisi dan menciptakan target.", "E"},
	{"Saya tertarik bidang pemasaran, keuangan, atau manajemen.", "E"},
	{"Saya senang berbicara di depan umum.", "E"},
	{"Saya suka membuat rencana dan mengambil keputusan besar.", "E"},
	{"Saya nyaman dalam situasi penuh tekanan.", "E"},
	{"Saya suka membangun relasi dan jaringan baru.", "E"},
	{"Saya ingin menjadi CEO, manajer, atau pengusaha.", "E"},
	{"Saya tertarik membuat strategi penjualan atau promosi.", "E"},
	{"Saya senang memimpin orang dan mengatur peran mereka.", "E"},
	{"Saya memiliki motivasi tinggi untuk sukses.", "E"},
	{"Saya merasa percaya diri dalam menyampaikan pendapat.", "E"},
	{"Saya tertarik dunia politik, hukum, atau manajemen proyek.", "E"},
	{"Saya termotivasi oleh pencapaian dan penghargaan.", "E"},
	{"Saya suka memecahkan masalah organisasi.", "E"},
	{"Saya percaya diri menghadapi risiko dan perubahan.", "E"},

	// Conventional
	{"Saya senang bekerja dengan angka dan data.", "C"},
	{"Saya suka mengatur berkas atau dokumen.", "C"},
	{"Saya tertarik bekerja sebagai akuntan atau administrator.", "C"},
	{"Saya merasa nyaman dengan aturan yang jelas.", "C"},
	{"Saya suka menggunakan spreadsheet dan perangkat lunak kantor.", "C"},
	{"Saya teliti dalam menyusun laporan atau anggaran.", "C"},
	{"Saya senang mengetik, mengarsip, dan mengatur data.", "C"},
	{"Saya suka mengatur keuangan pribadi dengan rapi.", "C"},
	{"Saya ingin bekerja di kantor yang terorganisir.", "C"},
	{"Saya lebih suka pekerjaan rutin dan stabil.", "C"},
	{"Saya senang dengan struktur dan prosedur kerja yang jelas.", "C"},
}

//Set skor awalnya 0 semua
var skorKategori = map[string]int{
	"R": 0, "I": 0, "A": 0, "S": 0, "E": 0, "C": 0,
}

//Buat ambil soal acak dari semua soal  
func ambilRandomSoal(soalList []Soal, jumlah int) []Soal {
	rand.Seed(time.Now().UnixNano())
	shuffled := make([]Soal, len(soalList))
	copy(shuffled, soalList)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	if jumlah > len(shuffled) {
		return shuffled
	}
	return shuffled[:jumlah]
}

//instruksi dan random appear soalnya
func TesMinatKeahlian() {
	fmt.Println()
	fmt.Println("🧠 TES MINAT & KEAHLIAN BERDASARKAN RIASEC")
	fmt.Println("============================================")
	fmt.Println("Jawab dengan pilihan berikut:")
	fmt.Println("  A = Sangat Suka / Yakin / Iya")
	fmt.Println("  B = Cukup Suka / Mungkin")
	fmt.Println("  C = Tidak Sama Sekali")
	fmt.Println("--------------------------------------------")

	// Input nama user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama Anda: ")
	inputNama, _ := reader.ReadString('\n')
	namaUser = strings.TrimSpace(inputNama)

	mapSoal := map[string][]Soal{
		"R": {}, "I": {}, "A": {}, "S": {}, "E": {}, "C": {},
	}
	for _, soal := range daftarSoal {
		mapSoal[soal.Kategori] = append(mapSoal[soal.Kategori], soal)
	}

	var soalTes []Soal
	for _, kode := range []string{"R", "I", "A", "S", "E", "C"} {
		soalTes = append(soalTes, ambilRandomSoal(mapSoal[kode], 5)...)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(soalTes), func(i, j int) {
		soalTes[i], soalTes[j] = soalTes[j], soalTes[i]
	})

	skorKategori = map[string]int{"R": 0, "I": 0, "A": 0, "S": 0, "E": 0, "C": 0}

	for i, soal := range soalTes {
		fmt.Printf("%d. %s\n", i+1, soal.Pertanyaan)
		fmt.Print("Jawaban Anda (A/B/C): ")
		input, _ := reader.ReadString('\n')
		jawaban := strings.ToUpper(strings.TrimSpace(input))

		switch jawaban {
		case "A":
			skorKategori[soal.Kategori] += 2
		case "B":
			skorKategori[soal.Kategori] += 1
		case "C":
			skorKategori[soal.Kategori] += 0
		default:
			fmt.Println("⚠️ Jawaban tidak valid. Skor diabaikan.")
		}
		fmt.Println()
	}
	tampilkanHasil()
}

//Hasil tesnya disini
var hasilDominan string
var namaUser string
func tampilkanHasil() {
	fmt.Println("📊 HASIL TES MINAT & KEAHLIAN")
	fmt.Println("--------------------------------")

	tipe := map[string]string{
		"R": "Realistic", "I": "Investigative", "A": "Artistic",
		"S": "Social", "E": "Enterprising", "C": "Conventional",
	}

	var maxSkor int
	var dominan string
	for kode, nilai := range skorKategori {
		fmt.Printf("%-14s: %d\n", tipe[kode], nilai)
		if nilai > maxSkor {
			maxSkor = nilai
			dominan = kode
		}
	}

	hasilDominan = dominan
	fmt.Println("\n🎯 Minat Dominan Anda:", tipe[dominan])
	fmt.Println("💼 Rekomendasi Karier:")
	fmt.Println(saranKarierString(dominan))

	//tanya simpan hasil disini
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nApakah Anda ingin menyimpan hasil tes ini? (Y/N): ")
	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.ToUpper(strings.TrimSpace(pilihan))

	if pilihan == "Y" {
		simpanHasilTes()
		fmt.Println("✅ Hasil tes berhasil disimpan.")
	} else {
		fmt.Println("ℹ️ Hasil tes tidak disimpan.")
	}
}

//Simpan hasilnya disini
func simpanHasilTes() {
	folder := "memory"
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println("❌ Gagal membuat folder:", err)
		return
	}

	base := "hasiltes"
	ext := ".txt"
	idx := 1
	var filename string
	for {
		filename = filepath.Join(folder, fmt.Sprintf("%s%d%s", base, idx, ext))
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			break
		}
		idx++
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("❌ Gagal membuat file:", err)
		return
	}
	defer file.Close()

	now := time.Now().Format("2006-01-02 15:04:05")

	namaKategori := map[string]string{
		"R": "Realistic", "I": "Investigative", "A": "Artistic",
		"S": "Social", "E": "Enterprising", "C": "Conventional",
	}

	fmt.Fprintf(file, "Tanggal Tes : %s\n", now)
	fmt.Fprintf(file, "Nama Peserta: %s\n\n", namaUser)

	fmt.Fprintln(file, "Skor Kategori RIASEC:")
	for _, kode := range []string{"R", "I", "A", "S", "E", "C"} {
		fmt.Fprintf(file, "%s : %d\n", namaKategori[kode], skorKategori[kode])
	}
	dominanLengkap := namaKategori[hasilDominan]

	file.WriteString("\nMinat Dominan: " + dominanLengkap + "\n\n")
	file.WriteString("Rekomendasi Karier:\n")
	for _, saran := range strings.Split(saranKarierString(hasilDominan), "\n") {
		file.WriteString(saran + "\n")
	}
}

//Saran karirnya disini
func saranKarierString(kode string) string {
	switch kode {
	case "R":
		return "- Teknisi, Mekanik, Insinyur, Konstruksi\n"
	case "I":
		return "- Ilmuwan, Data Analyst, Peneliti\n"
	case "A":
		return "- Desainer, Seniman, Penulis, Konten Kreator\n"
	case "S":
		return "- Guru, Konselor, Psikolog, Aktivis Sosial\n"
	case "E":
		return "- Manajer, Wirausahawan, Sales, Politisi\n"
	case "C":
		return "- Akuntan, Admin, Analis Sistem, Auditor\n"
	default:
		return "- Tidak diketahui\n"
	}
}
