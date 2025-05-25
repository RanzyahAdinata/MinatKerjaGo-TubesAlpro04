package main
import (
	"fmt"
	"os"
	"MinatKerjaGo/functions"
)

func main() {
	for {
		clearScreen()
		printHeader()
		printMenu()

		var pilihan int
		fmt.Print("🔍 Pilih menu (1-4): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			functions.TesMinatKeahlian()
		case 2:
			functions.InformasiKarier()
		case 3:
			functions.RiwayatHasilTes()
		case 4:
			functions.KonfirmasiKeluar()
			os.Exit(0)
		default:
			fmt.Println("⚠️ Pilihan tidak valid, silakan coba lagi.")
		}
		fmt.Println()
		pause()
	}
}

func printHeader() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      🎯 APLIKASI PEMILIHAN KARIER 🎯                 ║")
	fmt.Println("║         	BERDASARKAN PROFIL MINAT & KEAHLIAN PENGGUNA           ║")
	fmt.Println("║            	  ✨ by Ranzyah Adinata & Wafiq Aditya ✨              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                             👋 Selamat Datang!                        ")
	fmt.Println("         Silakan pilih menu di bawah untuk memulai penggunaan 🧭       ")
	fmt.Println("────────────────────────────────────────────────────────────────────────")
}

func printMenu() {
	fmt.Println("📋 1. Tes Minat & Keahlian")
	fmt.Println("📚 2. Informasi Karier")
	fmt.Println("📑 3. Riwayat Hasil Tes")
	fmt.Println("❌ 4. Exit")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") 
}

func pause() {
	fmt.Print("⏎ Tekan ENTER untuk kembali ke menu utama...")
	fmt.Scanln()
	fmt.Scanln()
}
