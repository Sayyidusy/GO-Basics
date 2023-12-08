package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
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

			//teks warna hijau
			fmt.Println("\033[32m Luas persegi panjang =", luas, "\033[0m")
			fmt.Print("\033[32m Keliling persegi panjang =", keliling, "\033[0m")
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

			//teks warna hijau
			fmt.Println("\033[32m Luas permukaan bola =", luasPermukaan, "\033[0m")
			fmt.Println("\033[32m Volume bola =", volume, "\033[0m")

		case 3:
			fmt.Println("Terima kasih")

		//jika pilihan tidak ada
		default:
			fmt.Println("Pilihan tidak tersedia")
		}

		fmt.Println("apakah anda ingin keluar? (Y/N)")
		var keluar string
		fmt.Scan(&keluar)

		keluar = strings.ToUpper(keluar)
		//kondisi jika ingin keluar
		//jika user memasukan huruf besar atau kecil maka akan tetap dijalankan
		if keluar == "Y" {
			break
		} else if keluar == "N" {
			continue
		} else {
			fmt.Println("Pilihan tidak tersedia")
		}
	}

	// soal latihan tipe data
	// 1
	// Buatlah program Go yang dapat menghitung
	// luas dan keliling persegi panjang dengan panjang
	// dan lebar yang dapat dimasukkan oleh pengguna.

	// 2
	// Buatlah program Go yang dapat menghitung volume dan
	// luas permukaan bola dengan jari-jari yang dapat
	// dimasukkan oleh pengguna.

}
