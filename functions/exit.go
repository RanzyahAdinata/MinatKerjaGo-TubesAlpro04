package functions

import "fmt"

func KonfirmasiKeluar() {
	fmt.Println("❓ Apakah Anda yakin ingin keluar? (Y/N)")
	var konfirmasi string
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" || konfirmasi == "Y" {
		fmt.Println("👋 Terima kasih telah menggunakan aplikasi ini.")
	} else {
		fmt.Println("✅ Kembali ke menu utama.")
	}
}
