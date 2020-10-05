package main

import "fmt"

var p,l,t int

func volume(){
	
	var v int

	fmt.Print("Masukkan Nilai Panjang: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan Nilai Lebar: ")
	fmt.Scan(&l)
	fmt.Print("Masukkan Nilai Tinggi: ")
	fmt.Scan(&t)

	v = p * l * t

	fmt.Println("Jadi Volumenya adalah = ", v)
}

func luas(){

	var L int

	fmt.Print("Masukkan Nilai Panjang: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan Nilai Lebar: ")
	fmt.Scan(&l)
	fmt.Print("Masukkan Nilai Tinggi: ")
	fmt.Scan(&t)

	L = 2*(p*l+p*t+l*t)

	fmt.Println("Jadi Luas Permukaannya adalah : ", L)
}

// func panjang(){
// 	fmt.Print("Rumus menggunakan volume atau luas permukaan")
// 	if
// }

func main() {
	var menu int

	fmt.Println("Hay Ini Kalkulator Balok Sederhana")
	fmt.Println("1. volume")
	fmt.Println("2. Luas Permukaan")
	// fmt.Println("3. Panjang")
	// fmt.Println("4. Lebar")
	// fmt.Println("5. Diagonal Bidang atau Sisi")
	// fmt.Println("6. Diagonal Ruang")
	// fmt.Println("7. Luas Bidang Diagonal")
	fmt.Print("Pilih angka diatas untuk menggunakan : ")
	fmt.Scan(&menu)
	switch {
		case menu == 1:
			volume()
		case menu == 2:
			luas()
		default:
			fmt.Println("Menu Tidak Ada")
	}
	
}
