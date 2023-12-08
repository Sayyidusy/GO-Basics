package main

import "fmt"

func main() {
	// operator perbandingan
	var x = 10
	var y = 5

	fmt.Println(x == y) // false
	fmt.Println(x != y) // true
	fmt.Println(x > y)  // true
	fmt.Println(x < y)  // false
	fmt.Println(x >= y) // true
	fmt.Println(x <= y) // false

	fmt.Println("=========== operator boolean =============")
	// operator boolean
	var ujian = 90
	var absensi = 80

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	// penulisan lebih singkat
	fmt.Println("lebih singkat =>", ujian > 80 && absensi > 80)

}
