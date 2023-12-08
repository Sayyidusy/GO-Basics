package main

import "fmt"

func main() {
	var n int    // banyak gerbong kereta
	fmt.Scan(&n) // input banyak gerbong kereta

	var arr = make([]int, n) // membuat slice dengan panjang n (banyak gerbong kereta) dan kapasitas n
	// membuat slice dengan panjang n (banyak gerbong kereta) dan kapasitas n sama dengan: var arr = []int{n}
	// for loop untuk mengisi slice arr dengan inputan user
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	// for loop untuk menampilkan slice arr dengan format AN,AN-1,...,A2,A1
	for i := n - 1; i >= 0; i-- {
		fmt.Print(arr[i])
		if i != 0 {
			fmt.Print(",")
		}
	}
	fmt.Println("")

}
