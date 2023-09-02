package main

import "fmt"

func main() {
	var name string

	name = "Akhmad Faiz Abdulloh"
	// variable akan error jika tidak digunakan
	fmt.Println(name)

	// ubah value variable
	name = "Akhmad Faiz"
	fmt.Println(name)

	// menyebutkan tipedata variable wajib
	// namun, opsional jika kita langsung inisialisasi datanya
	var friendName = "Budi"
	fmt.Println(friendName)

	// jika tidak di tambahkan int8 maka otomatis tereteksi sebagai int
	var age int8 = 20
	fmt.Println(age)

	// var bisa diganti dengan deklarasi awal :=
	country := "Indonesia"
	fmt.Println(country)

	// bisa membuat variable sekaligus banyak
	var (
		firstName = "Akhmad Faiz"
		lastName  = "Abdulloh"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
