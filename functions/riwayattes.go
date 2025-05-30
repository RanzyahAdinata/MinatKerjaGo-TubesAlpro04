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
		for i, label := range files {
			fmt.Printf("%d. %s\n", i+1, label)
		}

		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Lihat Detail Hasil Tes")
		fmt.Println("2. Cetak Hasil Tes ke PDF")
		fmt.Println("3. Hapus Riwayat Tes Satuan")
		fmt.Println("4. Hapus Semua Riwayat Tes")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi: ")

		fmt.Scanln(&opsi)

		// Extract nama file asli dari label
		realFiles := make([]string, len(files))
		for i, label := range files {
			realFiles[i] = strings.Split(label, " | ")[0]
		}

		switch opsi {
		case "1":
			pilihFileDanTampilkan(realFiles)
		case "2":
			pilihFileDanCetakPDF(realFiles)
		case "3":
			pilihFileDanHapus(realFiles)
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

	var hasil []string
	for _, path := range files {
		namaFile := filepath.Base(path)
		isidata, err := os.ReadFile(path)
		if err != nil {
			hasil = append(hasil, namaFile+" (gagal baca)")
			continue
		}

		var nama, tanggal string
		lines := strings.Split(string(isidata), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Nama:") {
				nama = strings.TrimSpace(strings.TrimPrefix(line, "Nama:"))
			}
			if strings.HasPrefix(line, "Tanggal Tes:") {
				tanggal = strings.TrimSpace(strings.TrimPrefix(line, "Tanggal Tes:"))
			}
		}

		label := fmt.Sprintf("%s | %s | %s", namaFile, nama, tanggal)
		hasil = append(hasil, label)
	}
	return hasil
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

	pdfFolder := filepath.Join(folder, "pdf")
	if _, err := os.Stat(pdfFolder); os.IsNotExist(err) {
		err = os.MkdirAll(pdfFolder, os.ModePerm)
		if err != nil {
			fmt.Println("❌ Gagal membuat folder pdf:", err)
			return
		}
	}

	lines := strings.Split(string(content), "\n")
	var tanggal, dominan, nama string
	skor := make(map[string]string)
	var rekomendasi []string
	section := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Nama:") {
			nama = strings.TrimSpace(strings.TrimPrefix(line, "Nama:"))
		} else if strings.HasPrefix(line, "Tanggal Tes:") {
			tanggal = strings.TrimSpace(strings.TrimPrefix(line, "Tanggal Tes:"))
		} else if strings.HasPrefix(line, "Minat Dominan:") {
			dominan = strings.TrimSpace(strings.TrimPrefix(line, "Minat Dominan:"))
			section = "dominan"
		} else if strings.HasPrefix(line, "Rekomendasi Karier:") {
			section = "rekomendasi"
		} else if strings.Contains(line, ":") && section != "rekomendasi" {
			parts := strings.SplitN(line, ":", 2)
			kunci := strings.TrimSpace(parts[0])
			nilai := strings.TrimSpace(parts[1])
			skor[kunci] = nilai
		} else if section == "rekomendasi" && strings.TrimSpace(line) != "" {
			rekomendasi = append(rekomendasi, line)
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Hasil Tes Minat & Keahlian", false)
	pdf.AddPage()

	softGray := []int{245, 245, 245}

	// Header
	pdf.SetFillColor(0, 102, 204)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 22)
	pdf.CellFormat(0, 15, "Hasil Tes Minat & Keahlian", "0", 1, "C", true, 0, "")
	pdf.Ln(5)

	// Nama dan Tanggal
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, "Nama: "+nama)
	pdf.Ln(6)
	pdf.Cell(0, 8, "Tanggal Tes: "+tanggal)
	pdf.Ln(12)

	// Skor RIASEC
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Skor RIASEC")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(softGray[0], softGray[1], softGray[2])
	for _, kategori := range []string{"Realistic", "Investigative", "Artistic", "Social", "Enterprising", "Conventional"} {
		if nilai, ok := skor[kategori]; ok {
			pdf.CellFormat(0, 8, fmt.Sprintf("%-13s : %s", kategori, nilai), "", 1, "", true, 0, "")
		}
	}
	pdf.Ln(10)

	// Dominan
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Minat Dominan")
	pdf.Ln(8)

	pdf.SetFont("Arial", "I", 12)
	pdf.MultiCell(0, 8, dominan, "", "", false)
	pdf.Ln(10)

	// Rekomendasi
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Rekomendasi Karier")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)
	for _, job := range rekomendasi {
		pdf.Cell(0, 8, "• "+job)
		pdf.Ln(6)
	}

	pdf.Ln(15)
	pdf.SetFont("Arial", "I", 10)
	pdf.SetTextColor(100, 100, 100)
	pdf.Cell(0, 10, "Dicetak otomatis oleh sistem MinatKerjaGo")

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
	var gagal int
	for _, label := range files {
		namaFile := strings.Split(label, " | ")[0]
		err := os.Remove(filepath.Join(folder, namaFile))
		if err != nil {
			fmt.Println("❌ Gagal menghapus:", namaFile, "-", err)
			gagal++
		}
	}

	if gagal == 0 {
		fmt.Println("🗑️ Semua riwayat tes berhasil dihapus.")
	} else {
		fmt.Printf("⚠️ %d file gagal dihapus.\n", gagal)
	}
}
