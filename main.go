package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Masukkan batas atas untuk mencari bilangan prima: ")
	fmt.Scan(&limit)
	fmt.Printf("Bilangan Prima antara 1 hingga %d adalah:\n", limit)
	for n := 2; n <= limit; n++ {
		isPrima := true
		if n > 1 {
			for i := 2; i * i <= n; i++ { // Kondisi untuk loop
				if n % i == 0 { 
					isPrima = false
					break
				}
			}
			if isPrima { 
				fmt.Println(n)
			}
		}
	}
}
