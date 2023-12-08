package main

import "fmt"

func main() {
	//tipe data array
	var pakaian [3]string
	pakaian[0] = "Kemeja"
	pakaian[1] = "Celana"
	pakaian[2] = "Jaket"

	fmt.Println(pakaian)

	//array bisa diisi langsung
	var buah = [3]string{"Apel", "Jeruk", "Mangga"}
	fmt.Println(buah)

	var angka = [5]int{
		0,
		1,
		2,
		3,
		4,
	}
	fmt.Println(angka)

	//array bisa diisi tanpa menuliskan jumlah elemennya
	// ... ini adalah operator variadic yang artinya jumlah elemen array akan menyesuaikan dengan jumlah elemen yang diisi
	var hewan = [...]string{"Kucing", "anjing", "Kelinci", "Kambing"}
	fmt.Println(hewan)

	//cara mengetahui jumlah elemen array
	fmt.Println(len(hewan))
	//cara mengetahui kapasitas array
	fmt.Println(cap(hewan))
	//cara mengetahui jumlah elemen yang belum terisi
	fmt.Println(cap(hewan) - len(hewan))

}
