package main

import (
	"fmt"
)

func main() {
	// ------------------------ KONTROL YAPILARI ------------------------
	yas := 20
	if yas >= 18 {
		fmt.Println("Reşit")
	} else if yas >= 13 {
		fmt.Println("Genç")
	} else {
		fmt.Println("Çocuk")
	}

	if sonuc := hesapla(-2, 2); sonuc > 0 {
		fmt.Println(sonuc)
	} else {
		fmt.Println("Sayı negatif")
	}

	// ------------------------ FOR ------------------------
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}

	// while gibi kullanma
	sayac := 0
	for sayac < 3 {
		fmt.Println("Son: ", sayac)
		sayac++
	}

	// ------------------------ SWITCH ------------------------
	switch gun := 3; gun {
	case 6, 7:
		fmt.Println(("Hafta Sonu"))
	case 1, 2, 3, 4, 5:
		fmt.Println("Hafta içi")
	default:
		fmt.Println("Geçersiz")
	}

	// ------------------------ DİZİLER VE SLICE'LAR ------------------------
	var sayilar [5]int // boyutu sabit, değişemez
	sayilar[0] = 10

	fmt.Println("Sayılar: ", sayilar)

	projeler := []string{"Portfolio Sitesi", "Deep Learning", "Go"}

	projeler = append(projeler, "Yeni Proje")

	fmt.Println("Proje sayısı: ", len(projeler))

	ilkIki := projeler[0:2]

	fmt.Println("Projeler: ", projeler)

	fmt.Println(ilkIki)

	// range ile döngü
	for index, deger := range projeler {
		fmt.Println(index, deger)
	}

	for _, deger := range projeler {
		fmt.Println(deger)
	}

}

func hesapla(a, b int) int {
	return a + b
}
