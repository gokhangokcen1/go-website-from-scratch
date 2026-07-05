package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hello world gokhan")
	fmt.Println("My favorite number is ", rand.Intn(10))

	// DEĞİŞKENLER
	var isim string
	isim = "Gökhan"

	var yas int = 24

	var maas float64 = 15000.50

	fmt.Printf("isim: %s, yas: %d, maas: %f\n", isim, yas, maas)

	var kuzen = "Furkan"

	yonetmen := "Kaan"
	yonetmen_yas := 21

	fmt.Println(kuzen, yonetmen, yonetmen_yas)

	// SABİT: CONST
	// diğerlerini kullanmak zorunluyken bunu kullanmak zorunlu değil
	const Pi = 3.14159

	// TEMEL TİPLER
	var a int = 5
	var c float64 = 2.5
	// a + c --> Hata, derlenmez
	sonuc := float64(a) + c
	fmt.Println("a + c = ", sonuc)

	// FONKSİYONLAR -> ÇALIŞTIRMA
	toplam_sonuc := toplam(7, 5)
	fmt.Println("7 + 5 = ", toplam_sonuc)

	b, k := bolVeKalan(17, 5)
	fmt.Println("17/5 -> Bolen - kalan: ", b, k)

}

// FONKSİYONLAR

func toplam(a, b int) int {
	return a + b
}

// func toplam(a int, b int) int{
// 	return a + b
// }

func bolVeKalan(a, b int) (int, int) {
	bolum := a / b
	kalan := a % b
	return bolum, kalan
}
