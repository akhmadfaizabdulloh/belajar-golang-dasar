	package main

	import "fmt"

	func main() {
		const firstName string = "Akhmad Faiz"
		const lastName = "Abdulloh"
		const value = 1000

	// nilai const tidak dapat diubah
	// firstName = "Arios"

	// nilai const tidak akan di komplain oleh si golang meskipun tidak di gunakan, berbeda dengan variable

	// fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// deklarasi multiple constant
	const (
		namaDepan    = "Akhmad"
		namaBelakang = "Faiz"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
