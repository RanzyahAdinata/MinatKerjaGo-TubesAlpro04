package functions
import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"github.com/jung-kurt/gofpdf"
)
const folder = "memory"

func RiwayatHasilTes() {
	var opsi string
	for {
		files := getDaftarHasilTes()
		if len(files) == 0 {
			fmt.Println("\n📭 Tidak ada hasil tes yang tersimpan.")
			return
		}

		fmt.Println("\n📁 DAFTAR HASIL TES")
		for i, file := range files {
			fmt.Printf("%d. %s\n", i+1, file)
		}

		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Lihat Detail Hasil Tes")
		fmt.Println("2. Cetak Hasil Tes ke PDF")
		fmt.Println("3. Hapus Riwayat Tes Satuan")
		fmt.Println("4. Hapus Semua Riwayat Tes")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi: ")

		fmt.Scanln(&opsi)

		switch opsi {
		case "1":
			pilihFileDanTampilkan(files)
		case "2":
			pilihFileDanCetakPDF(files)
		case "3":
			pilihFileDanHapus(files)
		case "4":
			hapusSemua()
		case "5":
			return
		default:
			fmt.Println("❌ Opsi tidak valid.")
		}
	}
}

func getDaftarHasilTes() []string {
	files, _ := filepath.Glob(filepath.Join(folder, "hasiltes*.txt"))
	sort.Strings(files)
	for i := range files {
		files[i] = filepath.Base(files[i])
	}
	return files
}

func pilihFileDanTampilkan(files []string) {
	fmt.Print("Masukkan nomor file yang ingin dilihat: ")
	var idx int
	fmt.Scanln(&idx)

	if idx < 1 || idx > len(files) {
		fmt.Println("❌ Nomor tidak valid.")
		return
	}
	path := filepath.Join(folder, files[idx-1])
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("❌ Gagal membaca file:", err)
		return
	}
	fmt.Println("\n📄 ISI HASIL TES")
	fmt.Println("----------------------------")
	fmt.Println(string(data))
}

func pilihFileDanHapus(files []string) {
	fmt.Print("Masukkan nomor file yang ingin dihapus: ")
	var idx int
	fmt.Scanln(&idx)

	if idx < 1 || idx > len(files) {
		fmt.Println("❌ Nomor tidak valid.")
		return
	}

	path := filepath.Join(folder, files[idx-1])
	if err := os.Remove(path); err != nil {
		fmt.Println("❌ Gagal menghapus:", err)
		return
	}
	fmt.Println("🗑️ File berhasil dihapus:", files[idx-1])
}

func pilihFileDanCetakPDF(files []string) {
	fmt.Print("Masukkan nomor file yang ingin dicetak ke PDF: ")
	var idx int
	fmt.Scanln(&idx)

	if idx < 1 || idx > len(files) {
		fmt.Println("❌ Nomor tidak valid.")
		return
	}

	txtPath := filepath.Join(folder, files[idx-1])
	content, err := os.ReadFile(txtPath)
	if err != nil {
		fmt.Println("❌ Gagal membaca file:", err)
		return
	}

	// Buat folder pdf kalok belum ada
	pdfFolder := filepath.Join(folder, "pdf")
	if _, err := os.Stat(pdfFolder); os.IsNotExist(err) {
		err = os.MkdirAll(pdfFolder, os.ModePerm)
		if err != nil {
			fmt.Println("❌ Gagal membuat folder pdf:", err)
			return
		}
	}

	// Persiapan data
	lines := strings.Split(string(content), "\n")
	var tanggal, dominan string
	var skor map[string]string = make(map[string]string)
	var rekomendasi []string

	// Parsing isi file
	section := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "Tanggal Tes:") {
			tanggal = strings.TrimSpace(strings.TrimPrefix(line, "Tanggal Tes:"))
		} else if strings.HasPrefix(line, "Minat Dominan:") {
			dominan = strings.TrimSpace(strings.TrimPrefix(line, "Minat Dominan:"))
			section = "dominan"
		} else if strings.HasPrefix(line, "Rekomendasi Karier:") {
			section = "rekomendasi"
		} else if strings.Contains(line, ":") && section != "rekomendasi" {
			parts := strings.SplitN(line, ":", 2)
			skor[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		} else if section == "rekomendasi" && strings.TrimSpace(line) != "" {
			rekomendasi = append(rekomendasi, line)
		}
	}

	// Mapping emoji per kategori
	ikon := map[string]string{
		"Realistic":     "🔧",
		"Investigative": "🔬",
		"Artistic":      "🎨",
		"Social":        "🤝",
		"Enterprising":  "💼",
		"Conventional":  "📊",
	}

	// Buat PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Hasil Tes Minat & Keahlian", false)
	pdf.AddPage()

	// Header
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, "🧠 Hasil Tes Minat & Keahlian")
	pdf.Ln(12)

	// Tanggal
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, "📅 Tanggal Tes: "+tanggal)
	pdf.Ln(10)

	// Skor
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "📊 Skor RIASEC")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	for kategori, nilai := range skor {
		icon := ikon[kategori]
		pdf.Cell(0, 8, fmt.Sprintf("%s %-13s : %s", icon, kategori, nilai))
		pdf.Ln(8)
	}

	// Dominan
	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "🎯 Minat Dominan Anda")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	dominanEmoji := ikon[dominan]
	pdf.MultiCell(0, 8, fmt.Sprintf("%s %s", dominanEmoji, dominan), "", "", false)

	// Rekomendasi Karier
	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "💼 Rekomendasi Karier")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	for _, job := range rekomendasi {
		pdf.Cell(0, 8, "• "+job)
		pdf.Ln(8)
	}

	// Simpan ke memory/pdf
	output := strings.Replace(files[idx-1], ".txt", ".pdf", 1)
	outputPath := filepath.Join(pdfFolder, output)
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		fmt.Println("❌ Gagal menyimpan PDF:", err)
		return
	}
	fmt.Println("✅ File PDF berhasil disimpan di:", outputPath)
}

func hapusSemua() {
	files := getDaftarHasilTes()
	if len(files) == 0 {
		fmt.Println("📭 Tidak ada file hasil tes.")
		return
	}
	for _, file := range files {
		os.Remove(filepath.Join(folder, file))
	}
	fmt.Println("🗑️ Semua riwayat tes berhasil dihapus.")
}
