package main

import "fmt"

// EGZERSİZ 1
// notlar := []int{85, 92, 78, 65, 90} şeklinde bir slice tanımla.
// for range kullanarak notların toplamını hesapla.
// Ortalamayı hesapla (float olacak, dikkat: int/int bölme tuzağına düşme).
// if/else if/else kullanarak ortalamaya göre harf notu belirle:

// 90 ve üzeri: AA
// 80-89: BA
// 70-79: BB
// 70 altı: FF

// Bunu bir fonksiyon içine al: func notOrtalamasiHesapla(notlar []int) (float64, string) — hem ortalamayı hem harf notunu return etsin.

// EGZERSİZ 2
// 1'den 30'a kadar say. 3'e bölünenlerde "Fizz", 5'e bölünenlerde "Buzz", ikisine de bölünenlerde "FizzBuzz", diğerlerinde sayının kendisini yazdır.

func main() {
	notlar := []int{85, 92, 78, 65, 90}
	fmt.Println(notOrtalamasiHesapla(notlar))
	fizzBuzz()

}

func notOrtalamasiHesapla(notlar []int) (float64, string) {
	var toplamNot float64
	var harfNot string
	for _, not := range notlar {
		toplamNot += float64(not)
	}

	ortalamaNot := toplamNot / float64(len(notlar))

	if ortalamaNot >= 90 {
		harfNot = "AA"
	} else if ortalamaNot >= 80 {
		harfNot = "BA"
	} else if ortalamaNot >= 70 {
		harfNot = "BB"
	} else {
		harfNot = "FF"
	}
	return ortalamaNot, harfNot
}

func fizzBuzz() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
