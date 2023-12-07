package main

import "fmt"

func mainn() {
	// tipe data int
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("ini tipe data int = ", 40)

	//tipe data float atau koma
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)

	//tipe data string
	fmt.Println("Halo" + " " + "Dunia")

	//tipe data boolean
	fmt.Println("benar =", true)
	fmt.Println("salah =", false)

	//len untuk menghitung jumlah karakter
	//spasi juga dihitung
	fmt.Println(len("Halo Dunia"))

	// =========== VARIABLE ===========
	//variable
	var nama string
	var umur int
	var lulus bool
	var ipk float32

	nama = "Rizky"
	umur = 20
	lulus = true
	ipk = 3.92

	//output
	fmt.Println("Nama saya adalah", nama)
	fmt.Println("Umur saya adalah", umur)
	fmt.Println("Apakah saya sudah kuliah lulus?", lulus)
	fmt.Println("IPK saya adalah", ipk)

	//variable bisa diisi langsung
	var nama2 string = "Rizky"
	var umur2 int = 20

	fmt.Println("Nama saya adalah", nama2, " dan umur saya adalah", umur2, "tahun")

	//var tidak wajib ditulis
	namaBarang := "Laptop"
	hargaBarang := 5000000
	fmt.Println(namaBarang)
	fmt.Println(hargaBarang)

	fmt.Println("=====================================")
	//========= CONSTANT ==========
	//variable yang nilainya tidak bisa diubah
	const nim = 123456789
	const nik = 987654321

	fmt.Println(nim)
	fmt.Println(nik)

	//cara penulisan constant lebih singkat
	const (
		noRek = 000000000
		noKtp = 111111111
		noKK  = 222222222
	)

	fmt.Println(
		"noRek =", noRek, "\n",
		"noKtp =", noKtp, "\n",
		"noKK =", noKK,
	)

	//konfersi tipe data
	var a int32 = 100000
	var b int64 = int64(a)
	// int8 = -128 sampai 127. karena 100000 melebihi 127 maka akan error
	//tapi jika int32 dan int64 tidak error karena int64 lebih besar dari int32
	var c int8 = int8(a)

	fmt.Println(a, b, c)
	//jadi c = -96 karena 100000 melebihi 127 maka akan diulang dari -128
}
