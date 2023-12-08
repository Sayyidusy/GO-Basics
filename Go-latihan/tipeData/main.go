package main

import "fmt"

func main() {

	//pilihan menu
	fmt.Println("=== Menu ===")
	fmt.Println("1. Luas dan Keliling Persegi Panjang")
	fmt.Println("2. Luas Permukaan dan Volume Bola")
	fmt.Println("3. Keluar")

	var pilihan int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println("=== luas dan keliling persegi panjang ===")
		var panjang, lebar int

		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)

		//rumus luas persegi panjang
		luas := panjang * lebar
		//rumus keliling persegi panjang
		keliling := 2 * (panjang + lebar)

		fmt.Println("Luas persegi panjang =", luas)
		fmt.Println("Keliling persegi panjang =", keliling)
		fmt.Println("")

	case 2:
		fmt.Println("=== luas permukaan dan volume bola ===")
		// rumus luas permukaan = 4 x π x r²
		// rumus volume = 4/3 x π x r³

		var jariJari float32

		fmt.Println("masukan jari-jari: ")
		fmt.Scan(&jariJari)

		// luas permukaan bola
		luasPermukaan := 4 * 3.14 * (jariJari * jariJari)
		//volume bola
		volume := 4 / 3 * 3.14 * (jariJari * jariJari * jariJari)

		fmt.Println("luas permukaan bola =", luasPermukaan)
		fmt.Println("volume bola =", volume)

	case 3:
		fmt.Println("=== keluar ===")
		fmt.Println("Terima kasih telah menggunakan program ini")

	//jika pilihan tidak ada
	default:
		fmt.Println("Pilihan tidak tersedia")
	}

	// 1
	// Buatlah program Go yang dapat menghitung
	// luas dan keliling persegi panjang dengan panjang
	// dan lebar yang dapat dimasukkan oleh pengguna.

	// 2
	// Buatlah program Go yang dapat menghitung volume dan
	// luas permukaan bola dengan jari-jari yang dapat
	// dimasukkan oleh pengguna.

}
