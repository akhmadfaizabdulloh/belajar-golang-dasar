package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	// karna int8 memiliki limit maksimum sampai 127
	// maka dia akan mengembalikan nilai dari yang maksimum, balik lagi ke yang paling minimum
	fmt.Println(nilai8)

	var name = "Faiz"
	var f = name[0]
	var fString string = string(f)

	fmt.Println(name)
	fmt.Println(fString)

	// variabel f >>> tipe datanya uint8 alias (byte)
}
