package main

import (
	"fmt"
)

func main() {
	var uzunluk float64 = 120
	var genislik float64 = 50

	sadeceAlan := alanHesapla(uzunluk, genislik)
	sadeceCevre := cevreHesapla(uzunluk, genislik)
	alan, cevre := alanVeCevreHesapla(uzunluk, genislik)

	fmt.Println("alanHesapla = ", sadeceAlan)
	fmt.Println("cevreHesapla = ", sadeceCevre)

	fmt.Printf("Alan= %.2f - Çevre = %.2f\n", alan, cevre)

	uzunluk1, genislik1 := 120, 50
	alan1, cevre1 := alanVeCevreHesaplaGaranti(uzunluk1, genislik1)
	fmt.Printf("Alan= %.2f - Çevre = %.2f\n", alan1, cevre1)

}

func alanHesapla(uzunluk, genislik float64) float64 {
	return uzunluk * genislik
}

func cevreHesapla(uzunluk, genislik float64) float64 {
	return 2 * (uzunluk + genislik)

}

func alanVeCevreHesapla(uzunluk, genislik float64) (float64, float64) {
	alan := uzunluk * genislik
	cevre := 2 * (uzunluk + genislik)

	return alan, cevre
}

func alanVeCevreHesaplaGaranti(uzunluk, genislik int) (float64, float64) {
	uzunluk_f64 := float64(uzunluk)
	genislik_f64 := float64(genislik)
	alan := uzunluk_f64 * genislik_f64
	cevre := 2 * (uzunluk_f64 + genislik_f64)

	return alan, cevre
}
